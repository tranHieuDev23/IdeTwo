package execution

// The status of the execution.
type ExecutionStatus int

const (
	NotExecuted ExecutionStatus = iota
	CompileError
	CompileTimeout
	RuntimeError
	RuntimeTimeout
	Successful
)

type Execution struct {
	// The unique id of this Execution.
	Id string `json:"id"`
	// The id of the Source Code this Execution belongs to.
	OfSourceCodeId string `json:"ofSourceCodeId"`
	// The time this Execution was initiated.
	Timestamp int64 `json:"timestamp"`
	// The status of the execution.
	Status ExecutionStatus `json:"status"`
	// The exit code of the compiled program.
	ExitCode int `json:"exitCode"`
	// The amount of time the execution took in millisecond.
	RunTime int64 `json:"runTime"`
	// If the program compiled unsuccessfully, this field's value is equal to
	// the compiler's error log.
	//
	// Otherwise, it's equal to the program's stdout output, trimmed to the
	// first 8 kB.
	Output string `json:"output"`
}
