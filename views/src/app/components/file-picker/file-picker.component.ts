import {
  Component,
  ElementRef,
  EventEmitter,
  Input,
  Output,
  ViewChild,
} from '@angular/core';

@Component({
  selector: 'app-file-picker',
  templateUrl: './file-picker.component.html',
  styleUrls: ['./file-picker.component.scss'],
})
export class FilePickerComponent {
  @ViewChild('input') public input: ElementRef<HTMLInputElement>;

  @Input() public multiple = false;
  @Input() public accept = '';

  @Output() public fileSelected = new EventEmitter<File[]>();

  public openFileDialog(): void {
    this.input?.nativeElement.click();
  }

  public onFileSelected(event: Event): void {
    const { files } = event.target as HTMLInputElement;
    const results = [];
    for (let i = 0; i < files.length; i++) {
      results.push(files.item(i));
    }
    this.fileSelected.emit(results);
  }
}
