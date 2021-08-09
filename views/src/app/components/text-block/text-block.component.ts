import { Component, Input } from '@angular/core';

export enum TextBlockType {
  Basic,
  Success,
  Error,
  Info,
  Warning,
}

@Component({
  selector: 'app-text-block',
  templateUrl: './text-block.component.html',
  styleUrls: ['./text-block.component.scss'],
})
export class TextBlockComponent {
  @Input() public text = '';
  @Input() public type = TextBlockType.Basic;

  constructor() {}
}
