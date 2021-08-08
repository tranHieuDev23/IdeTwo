import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { ProgrammingLanguage, SourceCode } from 'src/models/source_code';
import { Execution } from 'src/models/execution';

@Injectable({
  providedIn: 'root',
})
export class SourceCodeService {
  constructor(private readonly http: HttpClient) {}

  public async createSourceCode(
    name: string,
    language: ProgrammingLanguage,
    content: string,
    input: string
  ): Promise<SourceCode> {
    name = name.trim();
    return await this.http
      .post<SourceCode>(`/api/source_codes/`, {
        name,
        language,
        content,
        input,
      })
      .toPromise();
  }

  public async getSourceCode(id: string): Promise<SourceCode> {
    return await this.http
      .get<SourceCode>(`/api/source_codes/${id}`)
      .toPromise();
  }

  public async updateSourceCode(source: SourceCode): Promise<SourceCode> {
    let { id, name, language, content, input } = source;
    name = name.trim();
    return await this.http
      .patch<SourceCode>(`/api/source_codes/${id}`, {
        name,
        language,
        content,
        input,
      })
      .toPromise();
  }

  public async updateSourceCodeName(
    id: string,
    name: string
  ): Promise<SourceCode> {
    name = name.trim();
    return await this.http
      .patch<SourceCode>(`/api/source_codes/${id}/name`, {
        name,
      })
      .toPromise();
  }

  public async updateSourceCodeLanguage(
    id: string,
    language: ProgrammingLanguage
  ): Promise<SourceCode> {
    return await this.http
      .patch<SourceCode>(`/api/source_codes/${id}/language`, {
        language,
      })
      .toPromise();
  }

  public async executeSourceCode(id: string): Promise<Execution> {
    return await this.http
      .post<Execution>(`/api/source_codes/${id}/execute`, {})
      .toPromise();
  }
}
