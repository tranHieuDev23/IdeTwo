package source_code

// A programming language supported by the IdeTwo platform.
type ProgrammingLanguage int

const (
	C ProgrammingLanguage = iota
	Cpp
	Java
	Python3
)

// The status of the source code after execution.
type ExecutionStatus int

const (
	NotExecuted ExecutionStatus = iota
	CompileError
	CompileTimeout
	RuntimeError
	RuntimeTimeout
	Successful
)

// A source code on the IdeTwo platform.
type SourceCode struct {
	// The unique id of this source code.
	//
	// The user can use this id to access the source code on the website.
	Id string `json:"id"`
	// The programming language of this source code.
	Language ProgrammingLanguage `json:"language" valid:"range(0|4)"`
	// The content of the source code.
	//
	// Limited to up to 8 kB in size.
	Content string `json:"content" valid:"length(0|8192)"`
	// The lastest execution status of the source code.
	Status ExecutionStatus `json:"status" valid:"range(0:6)"`
	// The lastest input data for the source code.
	//
	// Limited to up to 8 kB in size.
	Input string `json:"input" valid:"length(0|8192),optional"`
	// The lastest output data of the source code.
	//
	// Limited to up to 8 kB in size.
	Output string `json:"output" valid:"length(0|8192),optional"`
}
