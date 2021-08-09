export enum ProgrammingLanguage {
  C,
  Cpp,
  Java,
  Python3,
}

export function getAllProgrammingLanguages(): ProgrammingLanguage[] {
  return [
    ProgrammingLanguage.C,
    ProgrammingLanguage.Cpp,
    ProgrammingLanguage.Java,
    ProgrammingLanguage.Python3,
  ];
}

const languageNames = ['C (GCC 8.2)', 'C++ 17', 'Java 13', 'Python 3.9'];

export function getProgrammingLanguageFromFilename(
  fileName: string
): ProgrammingLanguage {
  if (fileName.endsWith('c')) {
    return ProgrammingLanguage.C;
  }
  if (fileName.endsWith('cpp')) {
    return ProgrammingLanguage.Cpp;
  }
  if (fileName.endsWith('java')) {
    return ProgrammingLanguage.Java;
  }
  if (fileName.endsWith('py')) {
    return ProgrammingLanguage.Python3;
  }
  return ProgrammingLanguage.Cpp;
}

const languageDefaultFilename = ['main.c', 'main.cpp', 'Main.java', 'main.py'];

export function getProgrammingLanguageDefaultFilename(
  language: ProgrammingLanguage
): string {
  return languageDefaultFilename[language];
}

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
