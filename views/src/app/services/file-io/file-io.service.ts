import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root',
})
export class FileIoService {
  public async readFile(file: File): Promise<string> {
    return new Promise<string>((resolve) => {
      const reader = new FileReader();
      reader.onload = (event) => {
        resolve(event.target.result.toString());
      };
      reader.readAsText(file);
    });
  }
}
