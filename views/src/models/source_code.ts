export enum ProgrammingLanguage {
  C,
  Cpp,
  Java,
  Python3,
}

const languageNames = ['C (GCC 8.2)', 'C++ 17', 'Java 13', 'Python 3.9'];

export function getProgrammingLanguageName(
  language: ProgrammingLanguage
): string {
  return languageNames[language];
}

const languageModes = ['text/x-csrc', 'text/x-c++src', 'text/x-java', 'python'];

export function getProgrammingLanguageMode(
  language: ProgrammingLanguage
): string {
  return languageModes[language];
}

export class SourceCode {
  constructor(
    public id: string,
    public name: string,
    public language: ProgrammingLanguage,
    public content: string,
    public input: string
  ) {}
}
