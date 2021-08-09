import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { ProgrammingLanguage, SourceCode } from 'src/models/source_code';
import { Execution } from 'src/models/execution';
import { StatusCodes } from 'http-status-codes';

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
    try {
      return await this.http
        .post<SourceCode>(`/api/source_codes/`, {
          name,
          language,
          content,
          input,
        })
        .toPromise();
    } catch (e) {
      if (e instanceof HttpErrorResponse) {
        switch (e.status) {
          case StatusCodes.BAD_REQUEST:
            throw 'Check if your file name, source code or input meet the size requirement';
          case StatusCodes.INTERNAL_SERVER_ERROR:
            throw 'Internal server error';
        }
      }
      throw '';
    }
  }

  public async getSourceCode(id: string): Promise<SourceCode> {
    try {
      return await this.http
        .get<SourceCode>(`/api/source_codes/${id}`)
        .toPromise();
    } catch (e) {
      if (e instanceof HttpErrorResponse) {
        switch (e.status) {
          case StatusCodes.NOT_FOUND:
            throw 'Cannot find any source code with the provided id';
          case StatusCodes.INTERNAL_SERVER_ERROR:
            throw 'Internal server error';
        }
      }
      throw '';
    }
  }

  public async updateSourceCode(source: SourceCode): Promise<SourceCode> {
    let { id, name, language, content, input } = source;
    name = name.trim();
    try {
      return await this.http
        .patch<SourceCode>(`/api/source_codes/${id}`, {
          name,
          language,
          content,
          input,
        })
        .toPromise();
    } catch (e) {
      if (e instanceof HttpErrorResponse) {
        switch (e.status) {
          case StatusCodes.NOT_FOUND:
            throw 'Cannot find any source code with the provided id';
          case StatusCodes.BAD_REQUEST:
            throw 'Check if your source code or input meet the size requirement';
          case StatusCodes.INTERNAL_SERVER_ERROR:
            throw 'Internal server error';
        }
      }
      throw '';
    }
  }

  public async updateSourceCodeName(
    id: string,
    name: string
  ): Promise<SourceCode> {
    name = name.trim();
    try {
      return await this.http
        .patch<SourceCode>(`/api/source_codes/${id}/name`, {
          name,
        })
        .toPromise();
    } catch (e) {
      if (e instanceof HttpErrorResponse) {
        switch (e.status) {
          case StatusCodes.NOT_FOUND:
            throw 'Cannot find any source code with the provided id';
          case StatusCodes.BAD_REQUEST:
            throw 'Name cannot be longer than 128 characters';
          case StatusCodes.INTERNAL_SERVER_ERROR:
            throw 'Internal server error';
        }
      }
      throw '';
    }
  }

  public async updateSourceCodeLanguage(
    id: string,
    language: ProgrammingLanguage
  ): Promise<SourceCode> {
    try {
      return await this.http
        .patch<SourceCode>(`/api/source_codes/${id}/language`, {
          language,
        })
        .toPromise();
    } catch (e) {
      if (e instanceof HttpErrorResponse) {
        switch (e.status) {
          case StatusCodes.NOT_FOUND:
            throw 'Cannot find any source code with the provided id';
          case StatusCodes.BAD_REQUEST:
            throw 'Invalid programming language';
          case StatusCodes.INTERNAL_SERVER_ERROR:
            throw 'Internal server error';
        }
      }
      throw '';
    }
  }

  public async executeSourceCode(id: string): Promise<Execution> {
    try {
      return await this.http
        .post<Execution>(`/api/source_codes/${id}/execute`, {})
        .toPromise();
    } catch (e) {
      if (e instanceof HttpErrorResponse) {
        switch (e.status) {
          case StatusCodes.NOT_FOUND:
            throw 'Cannot find any source code with the provided id';
          case StatusCodes.INTERNAL_SERVER_ERROR:
            throw 'Internal server error';
        }
      }
      throw '';
    }
  }
}
