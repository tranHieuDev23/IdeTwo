import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FilePickerComponent } from './file-picker.component';

@NgModule({
  declarations: [FilePickerComponent],
  imports: [CommonModule],
  exports: [FilePickerComponent],
})
export class FilePickerModule {}
