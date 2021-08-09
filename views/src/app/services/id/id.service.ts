import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root',
})
export class IdService {
  private readonly alphabet =
    'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
  private readonly alphabetLength = this.alphabet.length;

  public makeId(length: number = 12): string {
    let result = '';
    for (let i = 0; i < length; i++) {
      result += this.alphabet.charAt(
        Math.floor(Math.random() * this.alphabetLength)
      );
    }
    return result;
  }
}
