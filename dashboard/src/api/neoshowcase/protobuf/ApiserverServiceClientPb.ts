/**
 * @fileoverview gRPC-Web generated client stub for neoshowcase.protobuf
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.4.2
// 	protoc              v4.22.0
// source: neoshowcase/protobuf/apiserver.proto


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';
import * as neoshowcase_protobuf_apiserver_pb from '../../neoshowcase/protobuf/apiserver_pb';


export class ApplicationServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname.replace(/\/+$/, '');
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodDescriptorGetRepositories = new grpcWeb.MethodDescriptor(
    '/neoshowcase.protobuf.ApplicationService/GetRepositories',
    grpcWeb.MethodType.UNARY,
    google_protobuf_empty_pb.Empty,
    neoshowcase_protobuf_apiserver_pb.GetRepositoriesResponse,
    (request: google_protobuf_empty_pb.Empty) => {
      return request.serializeBinary();
    },
    neoshowcase_protobuf_apiserver_pb.GetRepositoriesResponse.deserializeBinary
  );

  getRepositories(
    request: google_protobuf_empty_pb.Empty,
    metadata: grpcWeb.Metadata | null): Promise<neoshowcase_protobuf_apiserver_pb.GetRepositoriesResponse>;

  getRepositories(
    request: google_protobuf_empty_pb.Empty,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: neoshowcase_protobuf_apiserver_pb.GetRepositoriesResponse) => void): grpcWeb.ClientReadableStream<neoshowcase_protobuf_apiserver_pb.GetRepositoriesResponse>;

  getRepositories(
    request: google_protobuf_empty_pb.Empty,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: neoshowcase_protobuf_apiserver_pb.GetRepositoriesResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/neoshowcase.protobuf.ApplicationService/GetRepositories',
        request,
        metadata || {},
        this.methodDescriptorGetRepositories,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/neoshowcase.protobuf.ApplicationService/GetRepositories',
    request,
    metadata || {},
    this.methodDescriptorGetRepositories);
  }

  methodDescriptorCreateRepository = new grpcWeb.MethodDescriptor(
    '/neoshowcase.protobuf.ApplicationService/CreateRepository',
    grpcWeb.MethodType.UNARY,
    neoshowcase_protobuf_apiserver_pb.CreateRepositoryRequest,
    neoshowcase_protobuf_apiserver_pb.Repository,
    (request: neoshowcase_protobuf_apiserver_pb.CreateRepositoryRequest) => {
      return request.serializeBinary();
    },
    neoshowcase_protobuf_apiserver_pb.Repository.deserializeBinary
  );

  createRepository(
    request: neoshowcase_protobuf_apiserver_pb.CreateRepositoryRequest,
    metadata: grpcWeb.Metadata | null): Promise<neoshowcase_protobuf_apiserver_pb.Repository>;

  createRepository(
    request: neoshowcase_protobuf_apiserver_pb.CreateRepositoryRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: neoshowcase_protobuf_apiserver_pb.Repository) => void): grpcWeb.ClientReadableStream<neoshowcase_protobuf_apiserver_pb.Repository>;

  createRepository(
    request: neoshowcase_protobuf_apiserver_pb.CreateRepositoryRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: neoshowcase_protobuf_apiserver_pb.Repository) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/neoshowcase.protobuf.ApplicationService/CreateRepository',
        request,
        metadata || {},
        this.methodDescriptorCreateRepository,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/neoshowcase.protobuf.ApplicationService/CreateRepository',
    request,
    metadata || {},
    this.methodDescriptorCreateRepository);
  }

  methodDescriptorGetApplications = new grpcWeb.MethodDescriptor(
    '/neoshowcase.protobuf.ApplicationService/GetApplications',
    grpcWeb.MethodType.UNARY,
    google_protobuf_empty_pb.Empty,
    neoshowcase_protobuf_apiserver_pb.GetApplicationsResponse,
    (request: google_protobuf_empty_pb.Empty) => {
      return request.serializeBinary();
    },
    neoshowcase_protobuf_apiserver_pb.GetApplicationsResponse.deserializeBinary
  );

  getApplications(
    request: google_protobuf_empty_pb.Empty,
    metadata: grpcWeb.Metadata | null): Promise<neoshowcase_protobuf_apiserver_pb.GetApplicationsResponse>;

  getApplications(
    request: google_protobuf_empty_pb.Empty,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: neoshowcase_protobuf_apiserver_pb.GetApplicationsResponse) => void): grpcWeb.ClientReadableStream<neoshowcase_protobuf_apiserver_pb.GetApplicationsResponse>;

  getApplications(
    request: google_protobuf_empty_pb.Empty,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: neoshowcase_protobuf_apiserver_pb.GetApplicationsResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/neoshowcase.protobuf.ApplicationService/GetApplications',
        request,
        metadata || {},
        this.methodDescriptorGetApplications,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/neoshowcase.protobuf.ApplicationService/GetApplications',
    request,
    metadata || {},
    this.methodDescriptorGetApplications);
  }

  methodDescriptorGetAvailableDomains = new grpcWeb.MethodDescriptor(
    '/neoshowcase.protobuf.ApplicationService/GetAvailableDomains',
    grpcWeb.MethodType.UNARY,
    google_protobuf_empty_pb.Empty,
    neoshowcase_protobuf_apiserver_pb.AvailableDomains,
    (request: google_protobuf_empty_pb.Empty) => {
      return request.serializeBinary();
    },
    neoshowcase_protobuf_apiserver_pb.AvailableDomains.deserializeBinary
  );

  getAvailableDomains(
    request: google_protobuf_empty_pb.Empty,
    metadata: grpcWeb.Metadata | null): Promise<neoshowcase_protobuf_apiserver_pb.AvailableDomains>;

  getAvailableDomains(
    request: google_protobuf_empty_pb.Empty,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: neoshowcase_protobuf_apiserver_pb.AvailableDomains) => void): grpcWeb.ClientReadableStream<neoshowcase_protobuf_apiserver_pb.AvailableDomains>;

  getAvailableDomains(
    request: google_protobuf_empty_pb.Empty,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: neoshowcase_protobuf_apiserver_pb.AvailableDomains) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/neoshowcase.protobuf.ApplicationService/GetAvailableDomains',
        request,
        metadata || {},
        this.methodDescriptorGetAvailableDomains,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/neoshowcase.protobuf.ApplicationService/GetAvailableDomains',
    request,
    metadata || {},
    this.methodDescriptorGetAvailableDomains);
  }

  methodDescriptorAddAvailableDomain = new grpcWeb.MethodDescriptor(
    '/neoshowcase.protobuf.ApplicationService/AddAvailableDomain',
    grpcWeb.MethodType.UNARY,
    neoshowcase_protobuf_apiserver_pb.AvailableDomain,
    google_protobuf_empty_pb.Empty,
    (request: neoshowcase_protobuf_apiserver_pb.AvailableDomain) => {
      return request.serializeBinary();
    },
    google_protobuf_empty_pb.Empty.deserializeBinary
  );

  addAvailableDomain(
    request: neoshowcase_protobuf_apiserver_pb.AvailableDomain,
    metadata: grpcWeb.Metadata | null): Promise<google_protobuf_empty_pb.Empty>;

  addAvailableDomain(
    request: neoshowcase_protobuf_apiserver_pb.AvailableDomain,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: google_protobuf_empty_pb.Empty) => void): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

  addAvailableDomain(
    request: neoshowcase_protobuf_apiserver_pb.AvailableDomain,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: google_protobuf_empty_pb.Empty) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/neoshowcase.protobuf.ApplicationService/AddAvailableDomain',
        request,
        metadata || {},
        this.methodDescriptorAddAvailableDomain,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/neoshowcase.protobuf.ApplicationService/AddAvailableDomain',
    request,
    metadata || {},
    this.methodDescriptorAddAvailableDomain);
  }

  methodDescriptorCreateApplication = new grpcWeb.MethodDescriptor(
    '/neoshowcase.protobuf.ApplicationService/CreateApplication',
    grpcWeb.MethodType.UNARY,
    neoshowcase_protobuf_apiserver_pb.CreateApplicationRequest,
    neoshowcase_protobuf_apiserver_pb.Application,
    (request: neoshowcase_protobuf_apiserver_pb.CreateApplicationRequest) => {
      return request.serializeBinary();
    },
    neoshowcase_protobuf_apiserver_pb.Application.deserializeBinary
  );

  createApplication(
    request: neoshowcase_protobuf_apiserver_pb.CreateApplicationRequest,
    metadata: grpcWeb.Metadata | null): Promise<neoshowcase_protobuf_apiserver_pb.Application>;

  createApplication(
    request: neoshowcase_protobuf_apiserver_pb.CreateApplicationRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: neoshowcase_protobuf_apiserver_pb.Application) => void): grpcWeb.ClientReadableStream<neoshowcase_protobuf_apiserver_pb.Application>;

  createApplication(
    request: neoshowcase_protobuf_apiserver_pb.CreateApplicationRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: neoshowcase_protobuf_apiserver_pb.Application) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/neoshowcase.protobuf.ApplicationService/CreateApplication',
        request,
        metadata || {},
        this.methodDescriptorCreateApplication,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/neoshowcase.protobuf.ApplicationService/CreateApplication',
    request,
    metadata || {},
    this.methodDescriptorCreateApplication);
  }

  methodDescriptorGetApplication = new grpcWeb.MethodDescriptor(
    '/neoshowcase.protobuf.ApplicationService/GetApplication',
    grpcWeb.MethodType.UNARY,
    neoshowcase_protobuf_apiserver_pb.ApplicationIdRequest,
    neoshowcase_protobuf_apiserver_pb.Application,
    (request: neoshowcase_protobuf_apiserver_pb.ApplicationIdRequest) => {
      return request.serializeBinary();
    },
    neoshowcase_protobuf_apiserver_pb.Application.deserializeBinary
  );

  getApplication(
    request: neoshowcase_protobuf_apiserver_pb.ApplicationIdRequest,
    metadata: grpcWeb.Metadata | null): Promise<neoshowcase_protobuf_apiserver_pb.Application>;

  getApplication(
    request: neoshowcase_protobuf_apiserver_pb.ApplicationIdRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: neoshowcase_protobuf_apiserver_pb.Application) => void): grpcWeb.ClientReadableStream<neoshowcase_protobuf_apiserver_pb.Application>;

  getApplication(
    request: neoshowcase_protobuf_apiserver_pb.ApplicationIdRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: neoshowcase_protobuf_apiserver_pb.Application) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/neoshowcase.protobuf.ApplicationService/GetApplication',
        request,
        metadata || {},
        this.methodDescriptorGetApplication,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/neoshowcase.protobuf.ApplicationService/GetApplication',
    request,
    metadata || {},
    this.methodDescriptorGetApplication);
  }

  methodDescriptorDeleteApplication = new grpcWeb.MethodDescriptor(
    '/neoshowcase.protobuf.ApplicationService/DeleteApplication',
    grpcWeb.MethodType.UNARY,
    neoshowcase_protobuf_apiserver_pb.ApplicationIdRequest,
    google_protobuf_empty_pb.Empty,
    (request: neoshowcase_protobuf_apiserver_pb.ApplicationIdRequest) => {
      return request.serializeBinary();
    },
    google_protobuf_empty_pb.Empty.deserializeBinary
  );

  deleteApplication(
    request: neoshowcase_protobuf_apiserver_pb.ApplicationIdRequest,
    metadata: grpcWeb.Metadata | null): Promise<google_protobuf_empty_pb.Empty>;

  deleteApplication(
    request: neoshowcase_protobuf_apiserver_pb.ApplicationIdRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: google_protobuf_empty_pb.Empty) => void): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

  deleteApplication(
    request: neoshowcase_protobuf_apiserver_pb.ApplicationIdRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: google_protobuf_empty_pb.Empty) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/neoshowcase.protobuf.ApplicationService/DeleteApplication',
        request,
        metadata || {},
        this.methodDescriptorDeleteApplication,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/neoshowcase.protobuf.ApplicationService/DeleteApplication',
    request,
    metadata || {},
    this.methodDescriptorDeleteApplication);
  }

  methodDescriptorGetApplicationBuilds = new grpcWeb.MethodDescriptor(
    '/neoshowcase.protobuf.ApplicationService/GetApplicationBuilds',
    grpcWeb.MethodType.UNARY,
    neoshowcase_protobuf_apiserver_pb.ApplicationIdRequest,
    neoshowcase_protobuf_apiserver_pb.GetApplicationBuildsResponse,
    (request: neoshowcase_protobuf_apiserver_pb.ApplicationIdRequest) => {
      return request.serializeBinary();
    },
    neoshowcase_protobuf_apiserver_pb.GetApplicationBuildsResponse.deserializeBinary
  );

  getApplicationBuilds(
    request: neoshowcase_protobuf_apiserver_pb.ApplicationIdRequest,
    metadata: grpcWeb.Metadata | null): Promise<neoshowcase_protobuf_apiserver_pb.GetApplicationBuildsResponse>;

  getApplicationBuilds(
    request: neoshowcase_protobuf_apiserver_pb.ApplicationIdRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: neoshowcase_protobuf_apiserver_pb.GetApplicationBuildsResponse) => void): grpcWeb.ClientReadableStream<neoshowcase_protobuf_apiserver_pb.GetApplicationBuildsResponse>;

  getApplicationBuilds(
    request: neoshowcase_protobuf_apiserver_pb.ApplicationIdRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: neoshowcase_protobuf_apiserver_pb.GetApplicationBuildsResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/neoshowcase.protobuf.ApplicationService/GetApplicationBuilds',
        request,
        metadata || {},
        this.methodDescriptorGetApplicationBuilds,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/neoshowcase.protobuf.ApplicationService/GetApplicationBuilds',
    request,
    metadata || {},
    this.methodDescriptorGetApplicationBuilds);
  }

  methodDescriptorGetApplicationBuild = new grpcWeb.MethodDescriptor(
    '/neoshowcase.protobuf.ApplicationService/GetApplicationBuild',
    grpcWeb.MethodType.UNARY,
    neoshowcase_protobuf_apiserver_pb.GetApplicationBuildRequest,
    neoshowcase_protobuf_apiserver_pb.Build,
    (request: neoshowcase_protobuf_apiserver_pb.GetApplicationBuildRequest) => {
      return request.serializeBinary();
    },
    neoshowcase_protobuf_apiserver_pb.Build.deserializeBinary
  );

  getApplicationBuild(
    request: neoshowcase_protobuf_apiserver_pb.GetApplicationBuildRequest,
    metadata: grpcWeb.Metadata | null): Promise<neoshowcase_protobuf_apiserver_pb.Build>;

  getApplicationBuild(
    request: neoshowcase_protobuf_apiserver_pb.GetApplicationBuildRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: neoshowcase_protobuf_apiserver_pb.Build) => void): grpcWeb.ClientReadableStream<neoshowcase_protobuf_apiserver_pb.Build>;

  getApplicationBuild(
    request: neoshowcase_protobuf_apiserver_pb.GetApplicationBuildRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: neoshowcase_protobuf_apiserver_pb.Build) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/neoshowcase.protobuf.ApplicationService/GetApplicationBuild',
        request,
        metadata || {},
        this.methodDescriptorGetApplicationBuild,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/neoshowcase.protobuf.ApplicationService/GetApplicationBuild',
    request,
    metadata || {},
    this.methodDescriptorGetApplicationBuild);
  }

  methodDescriptorGetApplicationBuildLog = new grpcWeb.MethodDescriptor(
    '/neoshowcase.protobuf.ApplicationService/GetApplicationBuildLog',
    grpcWeb.MethodType.UNARY,
    neoshowcase_protobuf_apiserver_pb.GetApplicationBuildLogRequest,
    neoshowcase_protobuf_apiserver_pb.BuildLog,
    (request: neoshowcase_protobuf_apiserver_pb.GetApplicationBuildLogRequest) => {
      return request.serializeBinary();
    },
    neoshowcase_protobuf_apiserver_pb.BuildLog.deserializeBinary
  );

  getApplicationBuildLog(
    request: neoshowcase_protobuf_apiserver_pb.GetApplicationBuildLogRequest,
    metadata: grpcWeb.Metadata | null): Promise<neoshowcase_protobuf_apiserver_pb.BuildLog>;

  getApplicationBuildLog(
    request: neoshowcase_protobuf_apiserver_pb.GetApplicationBuildLogRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: neoshowcase_protobuf_apiserver_pb.BuildLog) => void): grpcWeb.ClientReadableStream<neoshowcase_protobuf_apiserver_pb.BuildLog>;

  getApplicationBuildLog(
    request: neoshowcase_protobuf_apiserver_pb.GetApplicationBuildLogRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: neoshowcase_protobuf_apiserver_pb.BuildLog) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/neoshowcase.protobuf.ApplicationService/GetApplicationBuildLog',
        request,
        metadata || {},
        this.methodDescriptorGetApplicationBuildLog,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/neoshowcase.protobuf.ApplicationService/GetApplicationBuildLog',
    request,
    metadata || {},
    this.methodDescriptorGetApplicationBuildLog);
  }

  methodDescriptorGetApplicationBuildArtifact = new grpcWeb.MethodDescriptor(
    '/neoshowcase.protobuf.ApplicationService/GetApplicationBuildArtifact',
    grpcWeb.MethodType.UNARY,
    neoshowcase_protobuf_apiserver_pb.ApplicationIdRequest,
    neoshowcase_protobuf_apiserver_pb.ApplicationBuildArtifact,
    (request: neoshowcase_protobuf_apiserver_pb.ApplicationIdRequest) => {
      return request.serializeBinary();
    },
    neoshowcase_protobuf_apiserver_pb.ApplicationBuildArtifact.deserializeBinary
  );

  getApplicationBuildArtifact(
    request: neoshowcase_protobuf_apiserver_pb.ApplicationIdRequest,
    metadata: grpcWeb.Metadata | null): Promise<neoshowcase_protobuf_apiserver_pb.ApplicationBuildArtifact>;

  getApplicationBuildArtifact(
    request: neoshowcase_protobuf_apiserver_pb.ApplicationIdRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: neoshowcase_protobuf_apiserver_pb.ApplicationBuildArtifact) => void): grpcWeb.ClientReadableStream<neoshowcase_protobuf_apiserver_pb.ApplicationBuildArtifact>;

  getApplicationBuildArtifact(
    request: neoshowcase_protobuf_apiserver_pb.ApplicationIdRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: neoshowcase_protobuf_apiserver_pb.ApplicationBuildArtifact) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/neoshowcase.protobuf.ApplicationService/GetApplicationBuildArtifact',
        request,
        metadata || {},
        this.methodDescriptorGetApplicationBuildArtifact,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/neoshowcase.protobuf.ApplicationService/GetApplicationBuildArtifact',
    request,
    metadata || {},
    this.methodDescriptorGetApplicationBuildArtifact);
  }

  methodDescriptorGetApplicationEnvironmentVariables = new grpcWeb.MethodDescriptor(
    '/neoshowcase.protobuf.ApplicationService/GetApplicationEnvironmentVariables',
    grpcWeb.MethodType.UNARY,
    neoshowcase_protobuf_apiserver_pb.ApplicationIdRequest,
    neoshowcase_protobuf_apiserver_pb.ApplicationEnvironmentVariables,
    (request: neoshowcase_protobuf_apiserver_pb.ApplicationIdRequest) => {
      return request.serializeBinary();
    },
    neoshowcase_protobuf_apiserver_pb.ApplicationEnvironmentVariables.deserializeBinary
  );

  getApplicationEnvironmentVariables(
    request: neoshowcase_protobuf_apiserver_pb.ApplicationIdRequest,
    metadata: grpcWeb.Metadata | null): Promise<neoshowcase_protobuf_apiserver_pb.ApplicationEnvironmentVariables>;

  getApplicationEnvironmentVariables(
    request: neoshowcase_protobuf_apiserver_pb.ApplicationIdRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: neoshowcase_protobuf_apiserver_pb.ApplicationEnvironmentVariables) => void): grpcWeb.ClientReadableStream<neoshowcase_protobuf_apiserver_pb.ApplicationEnvironmentVariables>;

  getApplicationEnvironmentVariables(
    request: neoshowcase_protobuf_apiserver_pb.ApplicationIdRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: neoshowcase_protobuf_apiserver_pb.ApplicationEnvironmentVariables) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/neoshowcase.protobuf.ApplicationService/GetApplicationEnvironmentVariables',
        request,
        metadata || {},
        this.methodDescriptorGetApplicationEnvironmentVariables,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/neoshowcase.protobuf.ApplicationService/GetApplicationEnvironmentVariables',
    request,
    metadata || {},
    this.methodDescriptorGetApplicationEnvironmentVariables);
  }

  methodDescriptorSetApplicationEnvironmentVariable = new grpcWeb.MethodDescriptor(
    '/neoshowcase.protobuf.ApplicationService/SetApplicationEnvironmentVariable',
    grpcWeb.MethodType.UNARY,
    neoshowcase_protobuf_apiserver_pb.SetApplicationEnvironmentVariableRequest,
    google_protobuf_empty_pb.Empty,
    (request: neoshowcase_protobuf_apiserver_pb.SetApplicationEnvironmentVariableRequest) => {
      return request.serializeBinary();
    },
    google_protobuf_empty_pb.Empty.deserializeBinary
  );

  setApplicationEnvironmentVariable(
    request: neoshowcase_protobuf_apiserver_pb.SetApplicationEnvironmentVariableRequest,
    metadata: grpcWeb.Metadata | null): Promise<google_protobuf_empty_pb.Empty>;

  setApplicationEnvironmentVariable(
    request: neoshowcase_protobuf_apiserver_pb.SetApplicationEnvironmentVariableRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: google_protobuf_empty_pb.Empty) => void): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

  setApplicationEnvironmentVariable(
    request: neoshowcase_protobuf_apiserver_pb.SetApplicationEnvironmentVariableRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: google_protobuf_empty_pb.Empty) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/neoshowcase.protobuf.ApplicationService/SetApplicationEnvironmentVariable',
        request,
        metadata || {},
        this.methodDescriptorSetApplicationEnvironmentVariable,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/neoshowcase.protobuf.ApplicationService/SetApplicationEnvironmentVariable',
    request,
    metadata || {},
    this.methodDescriptorSetApplicationEnvironmentVariable);
  }

  methodDescriptorGetApplicationOutput = new grpcWeb.MethodDescriptor(
    '/neoshowcase.protobuf.ApplicationService/GetApplicationOutput',
    grpcWeb.MethodType.UNARY,
    neoshowcase_protobuf_apiserver_pb.ApplicationIdRequest,
    neoshowcase_protobuf_apiserver_pb.ApplicationOutput,
    (request: neoshowcase_protobuf_apiserver_pb.ApplicationIdRequest) => {
      return request.serializeBinary();
    },
    neoshowcase_protobuf_apiserver_pb.ApplicationOutput.deserializeBinary
  );

  getApplicationOutput(
    request: neoshowcase_protobuf_apiserver_pb.ApplicationIdRequest,
    metadata: grpcWeb.Metadata | null): Promise<neoshowcase_protobuf_apiserver_pb.ApplicationOutput>;

  getApplicationOutput(
    request: neoshowcase_protobuf_apiserver_pb.ApplicationIdRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: neoshowcase_protobuf_apiserver_pb.ApplicationOutput) => void): grpcWeb.ClientReadableStream<neoshowcase_protobuf_apiserver_pb.ApplicationOutput>;

  getApplicationOutput(
    request: neoshowcase_protobuf_apiserver_pb.ApplicationIdRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: neoshowcase_protobuf_apiserver_pb.ApplicationOutput) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/neoshowcase.protobuf.ApplicationService/GetApplicationOutput',
        request,
        metadata || {},
        this.methodDescriptorGetApplicationOutput,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/neoshowcase.protobuf.ApplicationService/GetApplicationOutput',
    request,
    metadata || {},
    this.methodDescriptorGetApplicationOutput);
  }

  methodDescriptorGetApplicationKeys = new grpcWeb.MethodDescriptor(
    '/neoshowcase.protobuf.ApplicationService/GetApplicationKeys',
    grpcWeb.MethodType.UNARY,
    neoshowcase_protobuf_apiserver_pb.ApplicationIdRequest,
    neoshowcase_protobuf_apiserver_pb.ApplicationKeys,
    (request: neoshowcase_protobuf_apiserver_pb.ApplicationIdRequest) => {
      return request.serializeBinary();
    },
    neoshowcase_protobuf_apiserver_pb.ApplicationKeys.deserializeBinary
  );

  getApplicationKeys(
    request: neoshowcase_protobuf_apiserver_pb.ApplicationIdRequest,
    metadata: grpcWeb.Metadata | null): Promise<neoshowcase_protobuf_apiserver_pb.ApplicationKeys>;

  getApplicationKeys(
    request: neoshowcase_protobuf_apiserver_pb.ApplicationIdRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: neoshowcase_protobuf_apiserver_pb.ApplicationKeys) => void): grpcWeb.ClientReadableStream<neoshowcase_protobuf_apiserver_pb.ApplicationKeys>;

  getApplicationKeys(
    request: neoshowcase_protobuf_apiserver_pb.ApplicationIdRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: neoshowcase_protobuf_apiserver_pb.ApplicationKeys) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/neoshowcase.protobuf.ApplicationService/GetApplicationKeys',
        request,
        metadata || {},
        this.methodDescriptorGetApplicationKeys,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/neoshowcase.protobuf.ApplicationService/GetApplicationKeys',
    request,
    metadata || {},
    this.methodDescriptorGetApplicationKeys);
  }

  methodDescriptorCancelBuild = new grpcWeb.MethodDescriptor(
    '/neoshowcase.protobuf.ApplicationService/CancelBuild',
    grpcWeb.MethodType.UNARY,
    neoshowcase_protobuf_apiserver_pb.CancelBuildRequest,
    google_protobuf_empty_pb.Empty,
    (request: neoshowcase_protobuf_apiserver_pb.CancelBuildRequest) => {
      return request.serializeBinary();
    },
    google_protobuf_empty_pb.Empty.deserializeBinary
  );

  cancelBuild(
    request: neoshowcase_protobuf_apiserver_pb.CancelBuildRequest,
    metadata: grpcWeb.Metadata | null): Promise<google_protobuf_empty_pb.Empty>;

  cancelBuild(
    request: neoshowcase_protobuf_apiserver_pb.CancelBuildRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: google_protobuf_empty_pb.Empty) => void): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

  cancelBuild(
    request: neoshowcase_protobuf_apiserver_pb.CancelBuildRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: google_protobuf_empty_pb.Empty) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/neoshowcase.protobuf.ApplicationService/CancelBuild',
        request,
        metadata || {},
        this.methodDescriptorCancelBuild,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/neoshowcase.protobuf.ApplicationService/CancelBuild',
    request,
    metadata || {},
    this.methodDescriptorCancelBuild);
  }

  methodDescriptorRetryCommitBuild = new grpcWeb.MethodDescriptor(
    '/neoshowcase.protobuf.ApplicationService/RetryCommitBuild',
    grpcWeb.MethodType.UNARY,
    neoshowcase_protobuf_apiserver_pb.RetryCommitBuildRequest,
    google_protobuf_empty_pb.Empty,
    (request: neoshowcase_protobuf_apiserver_pb.RetryCommitBuildRequest) => {
      return request.serializeBinary();
    },
    google_protobuf_empty_pb.Empty.deserializeBinary
  );

  retryCommitBuild(
    request: neoshowcase_protobuf_apiserver_pb.RetryCommitBuildRequest,
    metadata: grpcWeb.Metadata | null): Promise<google_protobuf_empty_pb.Empty>;

  retryCommitBuild(
    request: neoshowcase_protobuf_apiserver_pb.RetryCommitBuildRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: google_protobuf_empty_pb.Empty) => void): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

  retryCommitBuild(
    request: neoshowcase_protobuf_apiserver_pb.RetryCommitBuildRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: google_protobuf_empty_pb.Empty) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/neoshowcase.protobuf.ApplicationService/RetryCommitBuild',
        request,
        metadata || {},
        this.methodDescriptorRetryCommitBuild,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/neoshowcase.protobuf.ApplicationService/RetryCommitBuild',
    request,
    metadata || {},
    this.methodDescriptorRetryCommitBuild);
  }

  methodDescriptorStartApplication = new grpcWeb.MethodDescriptor(
    '/neoshowcase.protobuf.ApplicationService/StartApplication',
    grpcWeb.MethodType.UNARY,
    neoshowcase_protobuf_apiserver_pb.ApplicationIdRequest,
    google_protobuf_empty_pb.Empty,
    (request: neoshowcase_protobuf_apiserver_pb.ApplicationIdRequest) => {
      return request.serializeBinary();
    },
    google_protobuf_empty_pb.Empty.deserializeBinary
  );

  startApplication(
    request: neoshowcase_protobuf_apiserver_pb.ApplicationIdRequest,
    metadata: grpcWeb.Metadata | null): Promise<google_protobuf_empty_pb.Empty>;

  startApplication(
    request: neoshowcase_protobuf_apiserver_pb.ApplicationIdRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: google_protobuf_empty_pb.Empty) => void): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

  startApplication(
    request: neoshowcase_protobuf_apiserver_pb.ApplicationIdRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: google_protobuf_empty_pb.Empty) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/neoshowcase.protobuf.ApplicationService/StartApplication',
        request,
        metadata || {},
        this.methodDescriptorStartApplication,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/neoshowcase.protobuf.ApplicationService/StartApplication',
    request,
    metadata || {},
    this.methodDescriptorStartApplication);
  }

  methodDescriptorStopApplication = new grpcWeb.MethodDescriptor(
    '/neoshowcase.protobuf.ApplicationService/StopApplication',
    grpcWeb.MethodType.UNARY,
    neoshowcase_protobuf_apiserver_pb.ApplicationIdRequest,
    google_protobuf_empty_pb.Empty,
    (request: neoshowcase_protobuf_apiserver_pb.ApplicationIdRequest) => {
      return request.serializeBinary();
    },
    google_protobuf_empty_pb.Empty.deserializeBinary
  );

  stopApplication(
    request: neoshowcase_protobuf_apiserver_pb.ApplicationIdRequest,
    metadata: grpcWeb.Metadata | null): Promise<google_protobuf_empty_pb.Empty>;

  stopApplication(
    request: neoshowcase_protobuf_apiserver_pb.ApplicationIdRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: google_protobuf_empty_pb.Empty) => void): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

  stopApplication(
    request: neoshowcase_protobuf_apiserver_pb.ApplicationIdRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: google_protobuf_empty_pb.Empty) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/neoshowcase.protobuf.ApplicationService/StopApplication',
        request,
        metadata || {},
        this.methodDescriptorStopApplication,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/neoshowcase.protobuf.ApplicationService/StopApplication',
    request,
    metadata || {},
    this.methodDescriptorStopApplication);
  }

}

