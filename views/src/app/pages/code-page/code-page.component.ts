import { Component, OnInit } from '@angular/core';
import { Location } from '@angular/common';
import { ActivatedRoute } from '@angular/router';
import { SourceCodeService } from 'src/app/services/source-code/source-code.service';
import { Execution, getExecutionStatusString } from 'src/models/execution';
import {
  getProgrammingLanguageMode,
  getProgrammingLanguageName,
  ProgrammingLanguage,
  SourceCode,
} from 'src/models/source_code';

function newEmptySourceCode(): SourceCode {
  return new SourceCode(null, '', ProgrammingLanguage.Cpp, '', '');
}

function newEmptyExecution(): Execution {
  return new Execution(null, null, null, null, null, null, '');
}

@Component({
  selector: 'app-code-page',
  templateUrl: './code-page.component.html',
  styleUrls: ['./code-page.component.scss'],
})
export class CodePageComponent implements OnInit {
  public source = newEmptySourceCode();
  public execution = newEmptyExecution();

  public get languageMode(): string {
    if (!this.source || !this.source.language) {
      return null;
    }
    return getProgrammingLanguageMode(this.source.language);
  }

  public get languageName(): string {
    if (!this.source || !this.source.language) {
      return '';
    }
    return getProgrammingLanguageName(this.source.language);
  }

  public get statusString(): string {
    if (!this.execution || this.execution.status) {
      return '';
    }
    return getExecutionStatusString(this.execution.status);
  }

  constructor(
    private readonly route: ActivatedRoute,
    private readonly location: Location,
    private readonly sourceCodeService: SourceCodeService
  ) {}

  ngOnInit(): void {
    this.route.paramMap.subscribe((params) => {
      const id = params.get('id');
      if (!id) {
        this.newFile();
        return;
      }
      this.openFile(id).then();
    });
  }

  public newFile(): void {
    this.location.replaceState(`/`);
    this.source = newEmptySourceCode();
    this.execution = newEmptyExecution();
  }

  public async saveFile(): Promise<void> {
    try {
      if (!this.source.id) {
        this.source = await this.sourceCodeService.createSourceCode(
          this.source.name,
          this.source.language,
          this.source.content,
          this.source.input
        );
        this.location.replaceState(`/${this.source.id}`);
      } else {
        this.source = await this.sourceCodeService.updateSourceCode(
          this.source
        );
      }
    } catch {}
  }

  public async openFile(id: string): Promise<void> {
    try {
      this.source = await this.sourceCodeService.getSourceCode(id);
      this.execution = newEmptyExecution();
    } catch {
      this.newFile();
    }
  }

  public async saveToDevice(): Promise<void> {}

  public async changeName(name: string): Promise<void> {
    const id = this.source.id;
    const oldName = this.source.name;
    this.source.name = name;
    if (!id) {
      return;
    }
    try {
      this.source = await this.sourceCodeService.updateSourceCodeName(id, name);
    } catch {
      this.source.name = oldName;
    }
  }

  public async changeLanguageMode(
    language: ProgrammingLanguage
  ): Promise<void> {}

  public async run(): Promise<void> {}
}
