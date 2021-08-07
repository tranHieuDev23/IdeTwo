import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { CodePageComponent } from './code-page.component';
import { CodePageRoutingModule } from './code-page-routing.module';

@NgModule({
  declarations: [CodePageComponent],
  imports: [CommonModule, CodePageRoutingModule],
  exports: [CodePageComponent],
})
export class CodePageModule {}
