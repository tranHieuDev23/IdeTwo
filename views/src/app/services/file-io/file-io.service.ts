import { Injectable } from '@angular/core';
import { FileSaverService } from 'ngx-filesaver';

@Injectable({
  providedIn: 'root',
})
export class FileIoService {
  constructor(private readonly fileSaver: FileSaverService) {}

  public async readFile(file: File): Promise<string> {
    return new Promise<string>((resolve) => {
      const reader = new FileReader();
      reader.onload = (event) => {
        resolve(event.target.result.toString());
      };
      reader.readAsText(file);
    });
  }

  public saveFile(content: string, filename: string): void {
    this.fileSaver.saveText(content, filename);
  }
}
