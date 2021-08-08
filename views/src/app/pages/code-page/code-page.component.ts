import { Component, OnInit } from '@angular/core';
import { EditorConfiguration } from 'codemirror';

@Component({
  selector: 'app-code-page',
  templateUrl: './code-page.component.html',
  styleUrls: ['./code-page.component.scss'],
})
export class CodePageComponent implements OnInit {
  public content = '';
  public options: EditorConfiguration = {
    mode: 'text/x-c++src',
    lineNumbers: true,
    theme: 'ayu-mirage',
  };

  constructor() {}

  ngOnInit(): void {}
}
