package execute_worker

import "github.com/tranHieuDev23/IdeTwo/models/source_code"

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
	Execute(source source_code.SourceCode) string
}
