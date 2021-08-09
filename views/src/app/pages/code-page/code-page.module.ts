import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { CodePageComponent } from './code-page.component';
import { CodePageRoutingModule } from './code-page-routing.module';
import { CodemirrorModule } from '@ctrl/ngx-codemirror';
import { FormsModule } from '@angular/forms';
import { EditableTextModule } from 'src/app/components/editable-text/editable-text.module';
import { FilePickerModule } from 'src/app/components/file-picker/file-picker.module';
import { TextBlockModule } from 'src/app/components/text-block/text-block.module';

@NgModule({
  declarations: [CodePageComponent],
  imports: [
    CommonModule,
    CodePageRoutingModule,
    CodemirrorModule,
    FormsModule,
    EditableTextModule,
    FilePickerModule,
    TextBlockModule,
  ],
  exports: [CodePageComponent],
})
export class CodePageModule {}
