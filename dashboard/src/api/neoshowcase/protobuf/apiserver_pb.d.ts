import * as jspb from 'google-protobuf'

import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';
import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';
import * as neoshowcase_protobuf_null_pb from '../../neoshowcase/protobuf/null_pb';


export class Application extends jspb.Message {
  getId(): string;
  setId(value: string): Application;

  getRepositoryUrl(): string;
  setRepositoryUrl(value: string): Application;

  getBranchName(): string;
  setBranchName(value: string): Application;

  getBuildType(): BuildType;
  setBuildType(value: BuildType): Application;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Application.AsObject;
  static toObject(includeInstance: boolean, msg: Application): Application.AsObject;
  static serializeBinaryToWriter(message: Application, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Application;
  static deserializeBinaryFromReader(message: Application, reader: jspb.BinaryReader): Application;
}

export namespace Application {
  export type AsObject = {
    id: string,
    repositoryUrl: string,
    branchName: string,
    buildType: BuildType,
  }
}

export class ApplicationEnvironmentVariable extends jspb.Message {
  getKey(): string;
  setKey(value: string): ApplicationEnvironmentVariable;

  getValue(): string;
  setValue(value: string): ApplicationEnvironmentVariable;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ApplicationEnvironmentVariable.AsObject;
  static toObject(includeInstance: boolean, msg: ApplicationEnvironmentVariable): ApplicationEnvironmentVariable.AsObject;
  static serializeBinaryToWriter(message: ApplicationEnvironmentVariable, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ApplicationEnvironmentVariable;
  static deserializeBinaryFromReader(message: ApplicationEnvironmentVariable, reader: jspb.BinaryReader): ApplicationEnvironmentVariable;
}

export namespace ApplicationEnvironmentVariable {
  export type AsObject = {
    key: string,
    value: string,
  }
}

export class ApplicationEnvironmentVariables extends jspb.Message {
  getVariablesList(): Array<ApplicationEnvironmentVariables>;
  setVariablesList(value: Array<ApplicationEnvironmentVariables>): ApplicationEnvironmentVariables;
  clearVariablesList(): ApplicationEnvironmentVariables;
  addVariables(value?: ApplicationEnvironmentVariables, index?: number): ApplicationEnvironmentVariables;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ApplicationEnvironmentVariables.AsObject;
  static toObject(includeInstance: boolean, msg: ApplicationEnvironmentVariables): ApplicationEnvironmentVariables.AsObject;
  static serializeBinaryToWriter(message: ApplicationEnvironmentVariables, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ApplicationEnvironmentVariables;
  static deserializeBinaryFromReader(message: ApplicationEnvironmentVariables, reader: jspb.BinaryReader): ApplicationEnvironmentVariables;
}

export namespace ApplicationEnvironmentVariables {
  export type AsObject = {
    variablesList: Array<ApplicationEnvironmentVariables.AsObject>,
  }
}

export class MariaDbKey extends jspb.Message {
  getHost(): string;
  setHost(value: string): MariaDbKey;

  getDatabase(): string;
  setDatabase(value: string): MariaDbKey;

  getUser(): string;
  setUser(value: string): MariaDbKey;

  getPassword(): string;
  setPassword(value: string): MariaDbKey;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): MariaDbKey.AsObject;
  static toObject(includeInstance: boolean, msg: MariaDbKey): MariaDbKey.AsObject;
  static serializeBinaryToWriter(message: MariaDbKey, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): MariaDbKey;
  static deserializeBinaryFromReader(message: MariaDbKey, reader: jspb.BinaryReader): MariaDbKey;
}

export namespace MariaDbKey {
  export type AsObject = {
    host: string,
    database: string,
    user: string,
    password: string,
  }
}

export class MongoKey extends jspb.Message {
  getHost(): string;
  setHost(value: string): MongoKey;

  getDatabase(): string;
  setDatabase(value: string): MongoKey;

  getUser(): string;
  setUser(value: string): MongoKey;

  getPassword(): string;
  setPassword(value: string): MongoKey;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): MongoKey.AsObject;
  static toObject(includeInstance: boolean, msg: MongoKey): MongoKey.AsObject;
  static serializeBinaryToWriter(message: MongoKey, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): MongoKey;
  static deserializeBinaryFromReader(message: MongoKey, reader: jspb.BinaryReader): MongoKey;
}

export namespace MongoKey {
  export type AsObject = {
    host: string,
    database: string,
    user: string,
    password: string,
  }
}

export class ApplicationKeys extends jspb.Message {
  getMariadbkey(): MariaDbKey | undefined;
  setMariadbkey(value?: MariaDbKey): ApplicationKeys;
  hasMariadbkey(): boolean;
  clearMariadbkey(): ApplicationKeys;

  getMongokey(): MongoKey | undefined;
  setMongokey(value?: MongoKey): ApplicationKeys;
  hasMongokey(): boolean;
  clearMongokey(): ApplicationKeys;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ApplicationKeys.AsObject;
  static toObject(includeInstance: boolean, msg: ApplicationKeys): ApplicationKeys.AsObject;
  static serializeBinaryToWriter(message: ApplicationKeys, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ApplicationKeys;
  static deserializeBinaryFromReader(message: ApplicationKeys, reader: jspb.BinaryReader): ApplicationKeys;
}

export namespace ApplicationKeys {
  export type AsObject = {
    mariadbkey?: MariaDbKey.AsObject,
    mongokey?: MongoKey.AsObject,
  }
}

export class ApplicationBuildArtifact extends jspb.Message {
  getUrl(): string;
  setUrl(value: string): ApplicationBuildArtifact;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ApplicationBuildArtifact.AsObject;
  static toObject(includeInstance: boolean, msg: ApplicationBuildArtifact): ApplicationBuildArtifact.AsObject;
  static serializeBinaryToWriter(message: ApplicationBuildArtifact, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ApplicationBuildArtifact;
  static deserializeBinaryFromReader(message: ApplicationBuildArtifact, reader: jspb.BinaryReader): ApplicationBuildArtifact;
}

export namespace ApplicationBuildArtifact {
  export type AsObject = {
    url: string,
  }
}

export class ApplicationOutput extends jspb.Message {
  getOutput(): string;
  setOutput(value: string): ApplicationOutput;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ApplicationOutput.AsObject;
  static toObject(includeInstance: boolean, msg: ApplicationOutput): ApplicationOutput.AsObject;
  static serializeBinaryToWriter(message: ApplicationOutput, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ApplicationOutput;
  static deserializeBinaryFromReader(message: ApplicationOutput, reader: jspb.BinaryReader): ApplicationOutput;
}

export namespace ApplicationOutput {
  export type AsObject = {
    output: string,
  }
}

export class Build extends jspb.Message {
  getId(): string;
  setId(value: string): Build;

  getStatus(): Build.BuildStatus;
  setStatus(value: Build.BuildStatus): Build;

  getStartedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setStartedAt(value?: google_protobuf_timestamp_pb.Timestamp): Build;
  hasStartedAt(): boolean;
  clearStartedAt(): Build;

  getFinishedAt(): neoshowcase_protobuf_null_pb.NullTimestamp | undefined;
  setFinishedAt(value?: neoshowcase_protobuf_null_pb.NullTimestamp): Build;
  hasFinishedAt(): boolean;
  clearFinishedAt(): Build;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Build.AsObject;
  static toObject(includeInstance: boolean, msg: Build): Build.AsObject;
  static serializeBinaryToWriter(message: Build, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Build;
  static deserializeBinaryFromReader(message: Build, reader: jspb.BinaryReader): Build;
}

export namespace Build {
  export type AsObject = {
    id: string,
    status: Build.BuildStatus,
    startedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    finishedAt?: neoshowcase_protobuf_null_pb.NullTimestamp.AsObject,
  }

  export enum BuildStatus { 
    BUILDING = 0,
    SUCCEEDED = 1,
    FAILED = 2,
    CANCELLED = 3,
    QUEUED = 4,
    SKIPPED = 5,
  }
}

export class BuildLog extends jspb.Message {
  getOutput(): string;
  setOutput(value: string): BuildLog;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): BuildLog.AsObject;
  static toObject(includeInstance: boolean, msg: BuildLog): BuildLog.AsObject;
  static serializeBinaryToWriter(message: BuildLog, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): BuildLog;
  static deserializeBinaryFromReader(message: BuildLog, reader: jspb.BinaryReader): BuildLog;
}

export namespace BuildLog {
  export type AsObject = {
    output: string,
  }
}

export class GetApplicationsResponse extends jspb.Message {
  getApplicationsList(): Array<Application>;
  setApplicationsList(value: Array<Application>): GetApplicationsResponse;
  clearApplicationsList(): GetApplicationsResponse;
  addApplications(value?: Application, index?: number): Application;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetApplicationsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetApplicationsResponse): GetApplicationsResponse.AsObject;
  static serializeBinaryToWriter(message: GetApplicationsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetApplicationsResponse;
  static deserializeBinaryFromReader(message: GetApplicationsResponse, reader: jspb.BinaryReader): GetApplicationsResponse;
}

export namespace GetApplicationsResponse {
  export type AsObject = {
    applicationsList: Array<Application.AsObject>,
  }
}

export class CreateApplicationRequest extends jspb.Message {
  getRepositoryUrl(): string;
  setRepositoryUrl(value: string): CreateApplicationRequest;

  getBranchName(): string;
  setBranchName(value: string): CreateApplicationRequest;

  getBuildType(): BuildType;
  setBuildType(value: BuildType): CreateApplicationRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateApplicationRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateApplicationRequest): CreateApplicationRequest.AsObject;
  static serializeBinaryToWriter(message: CreateApplicationRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateApplicationRequest;
  static deserializeBinaryFromReader(message: CreateApplicationRequest, reader: jspb.BinaryReader): CreateApplicationRequest;
}

export namespace CreateApplicationRequest {
  export type AsObject = {
    repositoryUrl: string,
    branchName: string,
    buildType: BuildType,
  }
}

export class ApplicationIdRequest extends jspb.Message {
  getId(): string;
  setId(value: string): ApplicationIdRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ApplicationIdRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ApplicationIdRequest): ApplicationIdRequest.AsObject;
  static serializeBinaryToWriter(message: ApplicationIdRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ApplicationIdRequest;
  static deserializeBinaryFromReader(message: ApplicationIdRequest, reader: jspb.BinaryReader): ApplicationIdRequest;
}

export namespace ApplicationIdRequest {
  export type AsObject = {
    id: string,
  }
}

export class GetApplicationBuildsResponse extends jspb.Message {
  getBuildsList(): Array<Build>;
  setBuildsList(value: Array<Build>): GetApplicationBuildsResponse;
  clearBuildsList(): GetApplicationBuildsResponse;
  addBuilds(value?: Build, index?: number): Build;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetApplicationBuildsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetApplicationBuildsResponse): GetApplicationBuildsResponse.AsObject;
  static serializeBinaryToWriter(message: GetApplicationBuildsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetApplicationBuildsResponse;
  static deserializeBinaryFromReader(message: GetApplicationBuildsResponse, reader: jspb.BinaryReader): GetApplicationBuildsResponse;
}

export namespace GetApplicationBuildsResponse {
  export type AsObject = {
    buildsList: Array<Build.AsObject>,
  }
}

export class GetApplicationBuildRequest extends jspb.Message {
  getBuildId(): string;
  setBuildId(value: string): GetApplicationBuildRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetApplicationBuildRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetApplicationBuildRequest): GetApplicationBuildRequest.AsObject;
  static serializeBinaryToWriter(message: GetApplicationBuildRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetApplicationBuildRequest;
  static deserializeBinaryFromReader(message: GetApplicationBuildRequest, reader: jspb.BinaryReader): GetApplicationBuildRequest;
}

export namespace GetApplicationBuildRequest {
  export type AsObject = {
    buildId: string,
  }
}

export class GetApplicationBuildLogRequest extends jspb.Message {
  getBuildId(): string;
  setBuildId(value: string): GetApplicationBuildLogRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetApplicationBuildLogRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetApplicationBuildLogRequest): GetApplicationBuildLogRequest.AsObject;
  static serializeBinaryToWriter(message: GetApplicationBuildLogRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetApplicationBuildLogRequest;
  static deserializeBinaryFromReader(message: GetApplicationBuildLogRequest, reader: jspb.BinaryReader): GetApplicationBuildLogRequest;
}

export namespace GetApplicationBuildLogRequest {
  export type AsObject = {
    buildId: string,
  }
}

export class SetApplicationEnvironmentVariableRequest extends jspb.Message {
  getApplicationId(): string;
  setApplicationId(value: string): SetApplicationEnvironmentVariableRequest;

  getKey(): string;
  setKey(value: string): SetApplicationEnvironmentVariableRequest;

  getValue(): string;
  setValue(value: string): SetApplicationEnvironmentVariableRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SetApplicationEnvironmentVariableRequest.AsObject;
  static toObject(includeInstance: boolean, msg: SetApplicationEnvironmentVariableRequest): SetApplicationEnvironmentVariableRequest.AsObject;
  static serializeBinaryToWriter(message: SetApplicationEnvironmentVariableRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SetApplicationEnvironmentVariableRequest;
  static deserializeBinaryFromReader(message: SetApplicationEnvironmentVariableRequest, reader: jspb.BinaryReader): SetApplicationEnvironmentVariableRequest;
}

export namespace SetApplicationEnvironmentVariableRequest {
  export type AsObject = {
    applicationId: string,
    key: string,
    value: string,
  }
}

export enum BuildType { 
  IMAGE = 0,
  STATIC = 1,
}
