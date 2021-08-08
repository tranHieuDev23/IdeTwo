import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Execution } from 'src/models/execution';

@Injectable({
  providedIn: 'root',
})
export class ExecutionService {
  constructor(private readonly http: HttpClient) {}

  public async getExecution(id: string): Promise<Execution> {
    return await this.http.get<Execution>(`/api/executions/${id}`).toPromise();
  }
}
