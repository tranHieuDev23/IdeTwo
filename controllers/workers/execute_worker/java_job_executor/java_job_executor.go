package java_job_executor

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

// Logic to handle code execution for Java source codes.
type CJobExecutor struct {
	cli client.Client
}

var instance *CJobExecutor
var once sync.Once
var conf = configs.GetInstance()

func (executor CJobExecutor) Execute(source source_code.SourceCode) job_executor.JobExecutorOutput {
	dir := tempdir.New(conf.IdeTwoExecutionsDir)
	defer dir.Close()

	executor.writeSourceFile(dir, source)

	if err := executor.compileSourceFile(dir, source); err != nil {
		return *err
	}

	return *executor.runExecutable(dir, source)
}

// Write the source file to a temporary directory.
func (executor CJobExecutor) writeSourceFile(dir tempdir.TempDir, source source_code.SourceCode) {
	sourceFilePath := fmt.Sprintf("%s/Main.java", dir.GetPath())
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

// Run the compiler within a Debian container with javac.
func (executor CJobExecutor) compileSourceFile(dir tempdir.TempDir, source source_code.SourceCode) *job_executor.JobExecutorOutput {
	cli := executor.cli
	ctx := context.Background()
	pathBind := fmt.Sprintf("%s:/workdir", dir.GetPath())
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image:        "openjdk:13-buster",
		WorkingDir:   "/workdir",
		Cmd:          []string{"timeout", "30s", "javac", "Main.java"},
		AttachStdout: true,
		AttachStderr: true,
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
		Stdout: true,
		Stderr: true,
	})
	if err != nil {
		panic(err)
	}
	defer attachResp.Close()

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	okChan, errChan := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case data := <-okChan:
		if data.StatusCode == 0 {
			return nil
		}

		if data.StatusCode == timeoutStatusCode {
			return &job_executor.JobExecutorOutput{
				Status:  execution.CompileTimeout,
				RunTime: 0,
				Output:  "",
			}
		}
		stdoutBuffer := new(bytes.Buffer)
		stderrBuffer := new(bytes.Buffer)
		stdcopy.StdCopy(stdoutBuffer, stderrBuffer, attachResp.Reader)
		compilerLog := stderrBuffer.String()
		return &job_executor.JobExecutorOutput{
			Status:  execution.CompileError,
			RunTime: 0,
			Output:  compilerLog,
		}

	case err := <-errChan:
		panic(err)
	}
}

// Run the executable built from compileSourceFile().
//
// Return the program's output in stdout.
func (executor CJobExecutor) runExecutable(dir tempdir.TempDir, source source_code.SourceCode) *job_executor.JobExecutorOutput {
	cli := executor.cli
	ctx := context.Background()
	pathBind := fmt.Sprintf("%s:/workdir", dir.GetPath())
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image:        "openjdk:13-buster",
		WorkingDir:   "/workdir",
		Cmd:          []string{"timeout", "--foreground", "30s", "java", "Main", "|", "head", "-c", "8k"},
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

func GetInstance() CJobExecutor {
	once.Do(func() {
		cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
		if err != nil {
			panic(err)
		}
		instance = &CJobExecutor{
			cli: *cli,
		}
		instance.prepareImage()
	})
	return *instance
}

// Prepare the necessary Docker images, to save time when handling jobs.
func (executor *CJobExecutor) prepareImage() {
	cli := executor.cli
	ctx := context.Background()
	_, err := cli.ImagePull(ctx, "docker.io/library/openjdk:13-buster", types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
}
