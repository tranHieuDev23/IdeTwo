import { Component, ElementRef, OnInit, ViewChild } from '@angular/core';
import { Location } from '@angular/common';
import { ActivatedRoute } from '@angular/router';
import { SourceCodeService } from 'src/app/services/source-code/source-code.service';
import { Execution, getExecutionStatusString } from 'src/models/execution';
import {
  getAllProgrammingLanguages,
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
  @ViewChild('unsavedModal')
  public unsavedModalElement: ElementRef<HTMLElement>;

  public source = newEmptySourceCode();
  public execution = newEmptyExecution();
  public unsaved = false;

  public onUnsavedModalConfirmed: () => void = null;

  public get languageMode(): string {
    if (!this.source) {
      return null;
    }
    return getProgrammingLanguageMode(this.source.language);
  }

  public get languageName(): string {
    if (!this.source) {
      return '';
    }
    return getProgrammingLanguageName(this.source.language);
  }

  public get statusString(): string {
    if (!this.execution) {
      return '';
    }
    return getExecutionStatusString(this.execution.status);
  }

  public languageModeOptions: { value: ProgrammingLanguage; label: string }[] =
    getAllProgrammingLanguages().map((item) => {
      return {
        value: item,
        label: getProgrammingLanguageName(item),
      };
    });

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

  public onEditorUpdated(): void {
    this.unsaved = true;
  }

  public newFile(checkUnsaved: boolean = true): void {
    if (checkUnsaved && this.unsaved) {
      this.showUnsavedConfirmModal();
      this.onUnsavedModalConfirmed = () => this.newFile(false);
      return;
    }
    this.location.replaceState(`/`);
    this.source = newEmptySourceCode();
    this.execution = newEmptyExecution();
    this.unsaved = false;
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
      this.unsaved = false;
    } catch {}
  }

  public async openFile(
    id: string,
    checkUnsaved: boolean = true
  ): Promise<void> {
    if (checkUnsaved && this.unsaved) {
      this.showUnsavedConfirmModal();
      this.onUnsavedModalConfirmed = () => this.openFile(id, false).then();
      return;
    }
    try {
      this.source = await this.sourceCodeService.getSourceCode(id);
      this.execution = newEmptyExecution();
      this.unsaved = false;
    } catch {
      this.newFile();
    }
  }

  public async saveToDevice(): Promise<void> {}

  public async changeName(name: string): Promise<void> {
    if (name === this.source.name) {
      return;
    }
    const id = this.source.id;
    const oldName = this.source.name;
    this.source.name = name;
    if (!id) {
      return;
    }
    try {
      await this.sourceCodeService.updateSourceCodeName(id, name);
    } catch {
      this.source.name = oldName;
    }
  }

  public async changeLanguageMode(
    language: ProgrammingLanguage
  ): Promise<void> {
    if (language === this.source.language) {
      return;
    }
    const id = this.source.id;
    const oldLanguage = this.source.language;
    this.source.language = language;
    if (!id) {
      return;
    }
    try {
      await this.sourceCodeService.updateSourceCodeLanguage(id, language);
    } catch {
      this.source.language = oldLanguage;
    }
  }

  public async run(): Promise<void> {}

  private showUnsavedConfirmModal(): void {
    this.unsavedModalElement.nativeElement.click();
  }
}
