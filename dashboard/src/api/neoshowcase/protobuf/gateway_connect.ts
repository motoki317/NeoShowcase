// @generated by protoc-gen-connect-es v0.8.4 with parameter "target=ts"
// @generated from file neoshowcase/protobuf/gateway.proto (package neoshowcase.protobuf, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { Empty, MethodKind } from "@bufbuild/protobuf";
import { Application, ApplicationEnvVars, ApplicationIdRequest, ApplicationOutput, ArtifactContent, ArtifactIdRequest, AvailableDomain, AvailableDomains, Build, BuildIdRequest, BuildLog, CreateApplicationRequest, CreateRepositoryRequest, GetApplicationsResponse, GetBuildsResponse, GetOutputRequest, GetOutputResponse, GetOutputStreamRequest, GetRepositoriesResponse, GetSystemPublicKeyResponse, Repository, RepositoryIdRequest, RetryCommitBuildRequest, SetApplicationEnvVarRequest, UpdateApplicationRequest, UpdateRepositoryRequest, User } from "./gateway_pb.js";

/**
 * @generated from service neoshowcase.protobuf.APIService
 */
export const APIService = {
  typeName: "neoshowcase.protobuf.APIService",
  methods: {
    /**
     * System
     *
     * @generated from rpc neoshowcase.protobuf.APIService.GetMe
     */
    getMe: {
      name: "GetMe",
      I: Empty,
      O: User,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc neoshowcase.protobuf.APIService.GetSystemPublicKey
     */
    getSystemPublicKey: {
      name: "GetSystemPublicKey",
      I: Empty,
      O: GetSystemPublicKeyResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc neoshowcase.protobuf.APIService.GetAvailableDomains
     */
    getAvailableDomains: {
      name: "GetAvailableDomains",
      I: Empty,
      O: AvailableDomains,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc neoshowcase.protobuf.APIService.AddAvailableDomain
     */
    addAvailableDomain: {
      name: "AddAvailableDomain",
      I: AvailableDomain,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * Repository CRUD
     *
     * @generated from rpc neoshowcase.protobuf.APIService.CreateRepository
     */
    createRepository: {
      name: "CreateRepository",
      I: CreateRepositoryRequest,
      O: Repository,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc neoshowcase.protobuf.APIService.GetRepositories
     */
    getRepositories: {
      name: "GetRepositories",
      I: Empty,
      O: GetRepositoriesResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc neoshowcase.protobuf.APIService.GetRepository
     */
    getRepository: {
      name: "GetRepository",
      I: RepositoryIdRequest,
      O: Repository,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc neoshowcase.protobuf.APIService.UpdateRepository
     */
    updateRepository: {
      name: "UpdateRepository",
      I: UpdateRepositoryRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc neoshowcase.protobuf.APIService.DeleteRepository
     */
    deleteRepository: {
      name: "DeleteRepository",
      I: RepositoryIdRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * Application CRUD
     *
     * @generated from rpc neoshowcase.protobuf.APIService.CreateApplication
     */
    createApplication: {
      name: "CreateApplication",
      I: CreateApplicationRequest,
      O: Application,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc neoshowcase.protobuf.APIService.GetApplications
     */
    getApplications: {
      name: "GetApplications",
      I: Empty,
      O: GetApplicationsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc neoshowcase.protobuf.APIService.GetApplication
     */
    getApplication: {
      name: "GetApplication",
      I: ApplicationIdRequest,
      O: Application,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc neoshowcase.protobuf.APIService.UpdateApplication
     */
    updateApplication: {
      name: "UpdateApplication",
      I: UpdateApplicationRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc neoshowcase.protobuf.APIService.DeleteApplication
     */
    deleteApplication: {
      name: "DeleteApplication",
      I: ApplicationIdRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * Application info / config
     *
     * @generated from rpc neoshowcase.protobuf.APIService.GetEnvVars
     */
    getEnvVars: {
      name: "GetEnvVars",
      I: ApplicationIdRequest,
      O: ApplicationEnvVars,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc neoshowcase.protobuf.APIService.SetEnvVar
     */
    setEnvVar: {
      name: "SetEnvVar",
      I: SetApplicationEnvVarRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc neoshowcase.protobuf.APIService.GetOutput
     */
    getOutput: {
      name: "GetOutput",
      I: GetOutputRequest,
      O: GetOutputResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc neoshowcase.protobuf.APIService.GetOutputStream
     */
    getOutputStream: {
      name: "GetOutputStream",
      I: GetOutputStreamRequest,
      O: ApplicationOutput,
      kind: MethodKind.ServerStreaming,
    },
    /**
     * @generated from rpc neoshowcase.protobuf.APIService.StartApplication
     */
    startApplication: {
      name: "StartApplication",
      I: ApplicationIdRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc neoshowcase.protobuf.APIService.StopApplication
     */
    stopApplication: {
      name: "StopApplication",
      I: ApplicationIdRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * Application builds
     *
     * @generated from rpc neoshowcase.protobuf.APIService.GetBuilds
     */
    getBuilds: {
      name: "GetBuilds",
      I: ApplicationIdRequest,
      O: GetBuildsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc neoshowcase.protobuf.APIService.GetBuild
     */
    getBuild: {
      name: "GetBuild",
      I: BuildIdRequest,
      O: Build,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc neoshowcase.protobuf.APIService.RetryCommitBuild
     */
    retryCommitBuild: {
      name: "RetryCommitBuild",
      I: RetryCommitBuildRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc neoshowcase.protobuf.APIService.CancelBuild
     */
    cancelBuild: {
      name: "CancelBuild",
      I: BuildIdRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc neoshowcase.protobuf.APIService.GetBuildLog
     */
    getBuildLog: {
      name: "GetBuildLog",
      I: BuildIdRequest,
      O: BuildLog,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc neoshowcase.protobuf.APIService.GetBuildLogStream
     */
    getBuildLogStream: {
      name: "GetBuildLogStream",
      I: BuildIdRequest,
      O: BuildLog,
      kind: MethodKind.ServerStreaming,
    },
    /**
     * @generated from rpc neoshowcase.protobuf.APIService.GetBuildArtifact
     */
    getBuildArtifact: {
      name: "GetBuildArtifact",
      I: ArtifactIdRequest,
      O: ArtifactContent,
      kind: MethodKind.Unary,
    },
  }
} as const;
