package source_code

// A programming language supported by the IdeTwo platform.
type ProgrammingLanguage int

const (
	C ProgrammingLanguage = iota
	Cpp
	Java
	Python3
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
	// Limited to up to 5 kB in size.
	Content string `json:"content" valid:"length(0|5000)"`
	// The last input data for the source code.
	//
	// Limited to up to 5 kB in size.
	Input string `json:"input" valid:"length(0|5000),optional"`
	// The last output data of the source code.
	//
	// Limited to up to 5 kB in size.
	Output string `json:"output" valid:"length(0|5000),optional"`
}
