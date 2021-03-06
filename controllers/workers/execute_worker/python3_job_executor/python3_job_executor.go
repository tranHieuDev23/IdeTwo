package python3_job_executor

import (
	"bytes"
	"context"
	"fmt"
	"io/fs"
	"io/ioutil"
	"sync"

	"github.com/araddon/dateparse"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"

	"github.com/tranHieuDev23/IdeTwo/controllers/workers/execute_worker/job_executor"
	"github.com/tranHieuDev23/IdeTwo/models/execution"
	"github.com/tranHieuDev23/IdeTwo/models/source_code"
	"github.com/tranHieuDev23/IdeTwo/utils/configs"
	"github.com/tranHieuDev23/IdeTwo/utils/tempdir"
)

// Logic to handle code execution for Python source codes.
type CppJobExecutor struct {
	cli client.Client
}

var instance *CppJobExecutor
var once sync.Once
var conf = configs.GetInstance()

func (executor CppJobExecutor) Execute(source source_code.SourceCode) job_executor.JobExecutorOutput {
	dir := tempdir.New(conf.IdeTwoExecutionsDir)
	defer dir.Close()

	executor.writeSourceFile(dir, source)
	return *executor.runExecutable(dir, source)
}

// Write the source file to a temporary directory.
func (executor CppJobExecutor) writeSourceFile(dir tempdir.TempDir, source source_code.SourceCode) {
	sourceFilePath := fmt.Sprintf("%s/main.py", dir.GetPath())
	err := ioutil.WriteFile(sourceFilePath, []byte(source.Content), fs.FileMode(0444))
	if err != nil {
		panic(err)
	}
}

var resourcesConf = container.Resources{
	// 1 GB of RAM
	Memory: 1073741824,
	// 1 CPU core
	CPUQuota: 100000,
}

const timeoutStatusCode = 124

// Run the executable built from compileSourceFile().
//
// Return the program's output in stdout.
func (executor CppJobExecutor) runExecutable(dir tempdir.TempDir, source source_code.SourceCode) *job_executor.JobExecutorOutput {
	cli := executor.cli
	ctx := context.Background()
	pathBind := fmt.Sprintf("%s:/workdir", dir.GetPath())
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image:        "python:3.9-buster",
		WorkingDir:   "/workdir",
		Cmd:          []string{"timeout", "--foreground", "30s", "python3", "main.py", "|", "head", "-c", "8k"},
		AttachStdin:  true,
		AttachStdout: true,
		OpenStdin:    true,
		StdinOnce:    true,
	}, &container.HostConfig{
		Binds:     []string{pathBind},
		Resources: resourcesConf,
	}, nil, nil, "")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := cli.ContainerRemove(ctx, resp.ID, types.ContainerRemoveOptions{}); err != nil {
			panic(err)
		}
	}()

	attachResp, err := cli.ContainerAttach(ctx, resp.ID, types.ContainerAttachOptions{
		Stream: true,
		Stdin:  true,
		Stdout: true,
	})
	if err != nil {
		panic(err)
	}
	defer attachResp.Close()

	attachResp.Conn.Write([]byte(source.Input))
	attachResp.Conn.Write([]byte("\n"))

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	okChan, errChan := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case <-okChan:
		inspectResp, err := cli.ContainerInspect(ctx, resp.ID)
		if err != nil {
			panic(err)
		}

		var status execution.ExecutionStatus
		switch inspectResp.State.ExitCode {
		case 0:
			status = execution.Successful
		case timeoutStatusCode:
			status = execution.RuntimeTimeout
		default:
			status = execution.RuntimeError
		}

		exitCode := inspectResp.State.ExitCode

		startTime, err := dateparse.ParseAny(inspectResp.State.StartedAt)
		if err != nil {
			panic(err)
		}
		finishTime, err := dateparse.ParseAny(inspectResp.State.FinishedAt)
		if err != nil {
			panic(err)
		}
		runTime := finishTime.Sub(startTime).Milliseconds()

		stdoutBuffer := new(bytes.Buffer)
		stderrBuffer := new(bytes.Buffer)
		stdcopy.StdCopy(stdoutBuffer, stderrBuffer, attachResp.Reader)
		stdout := stdoutBuffer.String()

		return &job_executor.JobExecutorOutput{
			Status:   status,
			ExitCode: exitCode,
			RunTime:  runTime,
			Output:   stdout,
		}

	case err := <-errChan:
		panic(err)
	}
}

func GetInstance() CppJobExecutor {
	once.Do(func() {
		cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
		if err != nil {
			panic(err)
		}
		instance = &CppJobExecutor{
			cli: *cli,
		}
		instance.prepareImage()
	})
	return *instance
}

// Prepare the necessary Docker images, to save time when handling jobs.
func (executor *CppJobExecutor) prepareImage() {
	cli := executor.cli
	ctx := context.Background()
	_, err := cli.ImagePull(ctx, "docker.io/library/python:3.9-buster", types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
}
