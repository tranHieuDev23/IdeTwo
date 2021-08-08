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
	// The name of the source code.
	//
	// Limited to up to 128 character.
	Name string `json:"name" valid:"length(0|128)"`
	// The programming language of this source code.
	Language ProgrammingLanguage `json:"language" valid:"range(0|4)"`
	// The content of the source code.
	//
	// Limited to up to 8 kB in size.
	Content string `json:"content" valid:"length(0|8192)"`
	// The lastest input data for the source code.
	//
	// Limited to up to 8 kB in size.
	Input string `json:"input" valid:"length(0|8192),optional"`
}
