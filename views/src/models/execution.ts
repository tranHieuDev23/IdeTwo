export enum ExecutionStatus {
  NotExecuted,
  CompileError,
  CompileTimeout,
  RuntimeError,
  RuntimeTimeout,
  Successful,
}

const executionStatusString = [
  'Not Executed',
  'Compile Error',
  'Compile Timeout',
  'Runtime Error',
  'Runtime Timeout',
  'Successful',
];

export function getExecutionStatusString(status: ExecutionStatus): string {
  return executionStatusString[status];
}

export class Execution {
  constructor(
    public id: string,
    public ofSourceCodeId: string,
    public timestamp: number,
    public status: ExecutionStatus,
    public exitCode: number,
    public runTime: number,
    public output: string
  ) {}
}
