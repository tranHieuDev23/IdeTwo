import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { CodePageComponent } from './code-page.component';
import { CodePageRoutingModule } from './code-page-routing.module';
import { CodemirrorModule } from '@ctrl/ngx-codemirror';
import { FormsModule } from '@angular/forms';

@NgModule({
  declarations: [CodePageComponent],
  imports: [CommonModule, CodePageRoutingModule, CodemirrorModule, FormsModule],
  exports: [CodePageComponent],
})
export class CodePageModule {}
