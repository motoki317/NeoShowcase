// @generated by protoc-gen-connect-es v1.1.4 with parameter "target=ts"
// @generated from file neoshowcase/protobuf/gateway.proto (package neoshowcase.protobuf, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { Empty, MethodKind } from "@bufbuild/protobuf";
import { Application, ApplicationEnvVars, ApplicationIdRequest, ApplicationMetrics, ApplicationOutput, ApplicationOutputs, ArtifactContent, ArtifactIdRequest, AvailableMetrics, Build, BuildIdRequest, BuildLog, CreateApplicationRequest, CreateRepositoryRequest, CreateUserKeyRequest, DeleteApplicationEnvVarRequest, DeleteUserKeyRequest, GenerateKeyPairResponse, GetAllBuildsRequest, GetApplicationMetricsRequest, GetApplicationsRequest, GetApplicationsResponse, GetBuildsResponse, GetOutputRequest, GetRepositoriesRequest, GetRepositoriesResponse, GetRepositoryRefsResponse, GetUserKeysResponse, GetUsersResponse, Repository, RepositoryIdRequest, RetryCommitBuildRequest, SetApplicationEnvVarRequest, SystemInfo, UpdateApplicationRequest, UpdateRepositoryRequest, User, UserKey } from "./gateway_pb.js";

/**
 * General / System
 *
 * @generated from service neoshowcase.protobuf.APIService
 */
export const APIService = {
  typeName: "neoshowcase.protobuf.APIService",
  methods: {
    /**
     * GetSystemInfo システム固有情報を取得します
     *
     * @generated from rpc neoshowcase.protobuf.APIService.GetSystemInfo
     */
    getSystemInfo: {
      name: "GetSystemInfo",
      I: Empty,
      O: SystemInfo,
      kind: MethodKind.Unary,
    },
    /**
     * GenerateKeyPair リポジトリ登録で使用する鍵ペアを一時的に生成します
     *
     * @generated from rpc neoshowcase.protobuf.APIService.GenerateKeyPair
     */
    generateKeyPair: {
      name: "GenerateKeyPair",
      I: Empty,
      O: GenerateKeyPairResponse,
      kind: MethodKind.Unary,
    },
    /**
     * GetMe 自身の情報を取得します プロキシ認証のため常に成功します
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
     * GetUsers 全てのユーザーの情報を取得します
     *
     * @generated from rpc neoshowcase.protobuf.APIService.GetUsers
     */
    getUsers: {
      name: "GetUsers",
      I: Empty,
      O: GetUsersResponse,
      kind: MethodKind.Unary,
    },
    /**
     * CreateUserKey アプリコンテナSSH用の公開鍵を登録します
     *
     * @generated from rpc neoshowcase.protobuf.APIService.CreateUserKey
     */
    createUserKey: {
      name: "CreateUserKey",
      I: CreateUserKeyRequest,
      O: UserKey,
      kind: MethodKind.Unary,
    },
    /**
     * GetUserKeys 登録した公開鍵一覧を取得します
     *
     * @generated from rpc neoshowcase.protobuf.APIService.GetUserKeys
     */
    getUserKeys: {
      name: "GetUserKeys",
      I: Empty,
      O: GetUserKeysResponse,
      kind: MethodKind.Unary,
    },
    /**
     * DeleteUserKey 登録した公開鍵を削除します
     *
     * @generated from rpc neoshowcase.protobuf.APIService.DeleteUserKey
     */
    deleteUserKey: {
      name: "DeleteUserKey",
      I: DeleteUserKeyRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * CreateRepository リポジトリを登録します
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
     * GetRepositories リポジトリ一覧を取得します
     *
     * @generated from rpc neoshowcase.protobuf.APIService.GetRepositories
     */
    getRepositories: {
      name: "GetRepositories",
      I: GetRepositoriesRequest,
      O: GetRepositoriesResponse,
      kind: MethodKind.Unary,
    },
    /**
     * GetRepository リポジトリを取得します
     *
     * @generated from rpc neoshowcase.protobuf.APIService.GetRepository
     */
    getRepository: {
      name: "GetRepository",
      I: RepositoryIdRequest,
      O: Repository,
      kind: MethodKind.Unary,
    },
    /**
     * GetRepositoryRefs リポジトリの現在の有効なref一覧を取得します
     *
     * @generated from rpc neoshowcase.protobuf.APIService.GetRepositoryRefs
     */
    getRepositoryRefs: {
      name: "GetRepositoryRefs",
      I: RepositoryIdRequest,
      O: GetRepositoryRefsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * UpdateRepository リポジトリ情報を更新します
     *
     * @generated from rpc neoshowcase.protobuf.APIService.UpdateRepository
     */
    updateRepository: {
      name: "UpdateRepository",
      I: UpdateRepositoryRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * RefreshRepository 自動更新間隔を待たず、手動でリモートリポジトリの最新情報に追従させます
     *
     * @generated from rpc neoshowcase.protobuf.APIService.RefreshRepository
     */
    refreshRepository: {
      name: "RefreshRepository",
      I: RepositoryIdRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * DeleteRepository リポジトリを削除します 関連する全てのアプリケーションの削除が必要です
     *
     * @generated from rpc neoshowcase.protobuf.APIService.DeleteRepository
     */
    deleteRepository: {
      name: "DeleteRepository",
      I: RepositoryIdRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * CreateApplication アプリを作成します
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
     * GetApplications アプリ一覧を取得します
     *
     * @generated from rpc neoshowcase.protobuf.APIService.GetApplications
     */
    getApplications: {
      name: "GetApplications",
      I: GetApplicationsRequest,
      O: GetApplicationsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * GetApplication アプリを取得します
     *
     * @generated from rpc neoshowcase.protobuf.APIService.GetApplication
     */
    getApplication: {
      name: "GetApplication",
      I: ApplicationIdRequest,
      O: Application,
      kind: MethodKind.Unary,
    },
    /**
     * UpdateApplication アプリ情報を更新します
     *
     * @generated from rpc neoshowcase.protobuf.APIService.UpdateApplication
     */
    updateApplication: {
      name: "UpdateApplication",
      I: UpdateApplicationRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * DeleteApplication アプリを削除します 先にアプリのシャットダウンが必要です
     *
     * @generated from rpc neoshowcase.protobuf.APIService.DeleteApplication
     */
    deleteApplication: {
      name: "DeleteApplication",
      I: ApplicationIdRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * GetAvailableMetrics 取得可能メトリクス一覧を取得します
     *
     * @generated from rpc neoshowcase.protobuf.APIService.GetAvailableMetrics
     */
    getAvailableMetrics: {
      name: "GetAvailableMetrics",
      I: Empty,
      O: AvailableMetrics,
      kind: MethodKind.Unary,
    },
    /**
     * GetApplicationMetrics アプリのメトリクスを取得します
     *
     * @generated from rpc neoshowcase.protobuf.APIService.GetApplicationMetrics
     */
    getApplicationMetrics: {
      name: "GetApplicationMetrics",
      I: GetApplicationMetricsRequest,
      O: ApplicationMetrics,
      kind: MethodKind.Unary,
    },
    /**
     * GetOutput アプリの出力を取得します
     *
     * @generated from rpc neoshowcase.protobuf.APIService.GetOutput
     */
    getOutput: {
      name: "GetOutput",
      I: GetOutputRequest,
      O: ApplicationOutputs,
      kind: MethodKind.Unary,
    },
    /**
     * GetOutputStream アプリの出力をストリーム形式で取得します
     *
     * @generated from rpc neoshowcase.protobuf.APIService.GetOutputStream
     */
    getOutputStream: {
      name: "GetOutputStream",
      I: ApplicationIdRequest,
      O: ApplicationOutput,
      kind: MethodKind.ServerStreaming,
    },
    /**
     * GetEnvVars アプリの環境変数を取得します
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
     * SetEnvVar アプリの環境変数をセットします システムによって設定された環境変数は上書きできません
     *
     * @generated from rpc neoshowcase.protobuf.APIService.SetEnvVar
     */
    setEnvVar: {
      name: "SetEnvVar",
      I: SetApplicationEnvVarRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * DeleteEnvVar アプリの環境変数を削除します システムによって設定された環境変数は削除できません
     *
     * @generated from rpc neoshowcase.protobuf.APIService.DeleteEnvVar
     */
    deleteEnvVar: {
      name: "DeleteEnvVar",
      I: DeleteApplicationEnvVarRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * StartApplication アプリを起動します 起動中の場合は再起動します
     *
     * @generated from rpc neoshowcase.protobuf.APIService.StartApplication
     */
    startApplication: {
      name: "StartApplication",
      I: ApplicationIdRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * StopApplication アプリをシャットダウンします
     *
     * @generated from rpc neoshowcase.protobuf.APIService.StopApplication
     */
    stopApplication: {
      name: "StopApplication",
      I: ApplicationIdRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * GetAllBuilds すべてのアプリケーションのビルドキューを取得します
     *
     * @generated from rpc neoshowcase.protobuf.APIService.GetAllBuilds
     */
    getAllBuilds: {
      name: "GetAllBuilds",
      I: GetAllBuildsRequest,
      O: GetBuildsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * GetBuilds アプリのビルド一覧を取得します
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
     * GetBuild アプリのビルド情報を取得します
     *
     * @generated from rpc neoshowcase.protobuf.APIService.GetBuild
     */
    getBuild: {
      name: "GetBuild",
      I: BuildIdRequest,
      O: Build,
      kind: MethodKind.Unary,
    },
    /**
     * RetryCommitBuild アプリの該当コミットのビルドをやり直します
     *
     * @generated from rpc neoshowcase.protobuf.APIService.RetryCommitBuild
     */
    retryCommitBuild: {
      name: "RetryCommitBuild",
      I: RetryCommitBuildRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * CancelBuild 該当ビルドが進行中の場合キャンセルします
     *
     * @generated from rpc neoshowcase.protobuf.APIService.CancelBuild
     */
    cancelBuild: {
      name: "CancelBuild",
      I: BuildIdRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * GetBuildLog 終了したビルドのログを取得します
     *
     * @generated from rpc neoshowcase.protobuf.APIService.GetBuildLog
     */
    getBuildLog: {
      name: "GetBuildLog",
      I: BuildIdRequest,
      O: BuildLog,
      kind: MethodKind.Unary,
    },
    /**
     * GetBuildLogStream ビルド中のログをストリーム形式で取得します
     *
     * @generated from rpc neoshowcase.protobuf.APIService.GetBuildLogStream
     */
    getBuildLogStream: {
      name: "GetBuildLogStream",
      I: BuildIdRequest,
      O: BuildLog,
      kind: MethodKind.ServerStreaming,
    },
    /**
     * GetBuildArtifact 静的サイトアプリの場合ビルド成果物（静的ファイルのtar）を取得します
     *
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

