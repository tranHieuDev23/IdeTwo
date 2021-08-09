import {
  Component,
  ElementRef,
  HostListener,
  OnInit,
  ViewChild,
} from '@angular/core';
import { Location } from '@angular/common';
import { ActivatedRoute } from '@angular/router';
import { SourceCodeService } from 'src/app/services/source-code/source-code.service';
import { Execution, getExecutionStatusString } from 'src/models/execution';
import {
  getAllProgrammingLanguages,
  getProgrammingLanguageFromFilename,
  getProgrammingLanguageMode,
  getProgrammingLanguageName,
  ProgrammingLanguage,
  SourceCode,
} from 'src/models/source_code';
import { NotificationService } from 'src/app/services/notification/notification.service';
import { FileIoService } from 'src/app/services/file-io/file-io.service';

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
  @ViewChild('unsavedModalToggler')
  public unsavedModalToggler: ElementRef<HTMLElement>;

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
    private readonly fileIo: FileIoService,
    private readonly sourceCodeService: SourceCodeService,
    private readonly notificationService: NotificationService
  ) {}

  ngOnInit(): void {
    this.route.paramMap.subscribe((params) => {
      const id = params.get('id');
      if (!id) {
        this.newFile();
        return;
      }
      this.openFileById(id).then();
    });
  }

  public onEditorUpdated(): void {
    this.unsaved = true;
  }

  public newFile(checkUnsaved: boolean = true): void {
    if (checkUnsaved && this.shouldShowUnsavedConfirmModal()) {
      this.showUnsavedConfirmModal();
      this.onUnsavedModalConfirmed = () => this.newFile(false);
      return;
    }
    this.location.replaceState(`/`);
    this.source = newEmptySourceCode();
    this.execution = newEmptyExecution();
    this.unsaved = true;
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
        this.notificationService.successNotification(
          'New file saved successfully'
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

  @HostListener('window:keydown.control.s', ['$event'])
  public async onSaveFileShortcut(event: KeyboardEvent): Promise<void> {
    if (event) {
      event.preventDefault();
    }
    await this.saveFile();
  }

  public async importLocalFile(
    files: File[],
    checkUnsaved: boolean = true
  ): Promise<void> {
    if (files.length === 0) {
      return;
    }
    if (checkUnsaved && this.shouldShowUnsavedConfirmModal()) {
      this.showUnsavedConfirmModal();
      this.onUnsavedModalConfirmed = () => this.importLocalFile(files, false);
      return;
    }
    const file = files[0];
    const content = await this.fileIo.readFile(file);
    const language = getProgrammingLanguageFromFilename(file.name);
    this.location.replaceState(`/`);
    this.source = new SourceCode(null, '', language, content, '');
    this.execution = newEmptyExecution();
    this.unsaved = true;
  }

  public async openFileById(
    id: string,
    checkUnsaved: boolean = true
  ): Promise<void> {
    if (checkUnsaved && this.shouldShowUnsavedConfirmModal()) {
      this.showUnsavedConfirmModal();
      this.onUnsavedModalConfirmed = () => this.openFileById(id, false).then();
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

  private shouldShowUnsavedConfirmModal(): boolean {
    if (!this.source.id) {
      return this.source.content.length > 0;
    }
    return this.unsaved;
  }

  private showUnsavedConfirmModal(): void {
    this.unsavedModalToggler.nativeElement.click();
  }
}
