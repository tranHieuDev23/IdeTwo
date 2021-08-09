import {
  Component,
  ElementRef,
  EventEmitter,
  Input,
  Output,
  ViewChild,
} from '@angular/core';

@Component({
  selector: 'app-editable-text',
  templateUrl: './editable-text.component.html',
  styleUrls: ['./editable-text.component.scss'],
})
export class EditableTextComponent {
  @ViewChild('textInput') public textInput: ElementRef<HTMLInputElement>;

  @Input() public text = '';
  @Input() public prompt = '';
  @Input() public maxLength: number = undefined;

  public isEditing = false;

  @Output() public textChanged = new EventEmitter<string>();

  constructor() {}

  public startEditing(): void {
    this.isEditing = true;
    setTimeout(() => {
      // this will make the execution after the above boolean has changed
      this.textInput.nativeElement.focus();
    }, 0);
  }

  public cancelEditing(): void {
    this.isEditing = false;
  }

  public submit(): void {
    this.text = this.text.trim();
    this.textChanged.emit(this.text);
    this.cancelEditing();
  }
}
