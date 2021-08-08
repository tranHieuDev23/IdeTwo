package job_executor

import (
	"github.com/tranHieuDev23/IdeTwo/models/execution"
	"github.com/tranHieuDev23/IdeTwo/models/source_code"
)

// The output of a JobExecutor instance.
type JobExecutorOutput struct {
	// The new status of the Execution
	Status execution.ExecutionStatus
	// The status code of the compiled program.
	ExitCode int
	// The amount of time the execution took in millisecond.
	RunTime int64
	// If the program compiled unsuccessfully, this field's value is equal to
	// the compiler's error log.
	//
	// Otherwise, it's equal to the program's stdout output, trimmed to the
	// first 8 kB.
	Output string
}

// Abstract interface of the logic to handle code execution for different
// programming language.
type JobExecutor interface {
	// Execute the program described by the provided source code, compiling the
	// program beforehand if necessary, and return either:
	//
	// - The output of the program in stdout, or
	//
	// - The error log of the compiler, if the source code compiles
	// unsuccessfully.
	Execute(source source_code.SourceCode) JobExecutorOutput
}
