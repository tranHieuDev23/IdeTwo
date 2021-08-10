import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { StatusCodes } from 'http-status-codes';
import { Observable } from 'rxjs';
import { Execution, ExecutionStatus } from 'src/models/execution';

const ONE_SECOND = 1000;

@Injectable({
  providedIn: 'root',
})
export class ExecutionService {
  private currentInterval: any = null;

  constructor(private readonly http: HttpClient) {}

  public async getExecution(id: string): Promise<Execution> {
    try {
      return await this.http
        .get<Execution>(`/api/executions/${id}`)
        .toPromise();
    } catch (e) {
      if (e instanceof HttpErrorResponse) {
        switch (e.status) {
          case StatusCodes.NOT_FOUND:
            throw 'Cannot find any execution with the provided id';
          case StatusCodes.INTERNAL_SERVER_ERROR:
            throw 'Internal server error';
        }
      }
      throw '';
    }
  }

  public observeExecution(id: string): Observable<Execution> {
    if (this.currentInterval) {
      clearInterval(this.currentInterval);
    }

    return new Observable((subscriber) => {
      this.currentInterval = setInterval(async () => {
        try {
          const execution = await this.getExecution(id);
          subscriber.next(execution);
          if (execution.status != ExecutionStatus.NotExecuted) {
            clearInterval(this.currentInterval);
            this.currentInterval = null;
            subscriber.complete();
          }
        } catch (e) {
          subscriber.error(e);
        }
      }, ONE_SECOND);
    });
  }
}
