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


export class AppsServiceClient {
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

  methodDescriptorGetApps = new grpcWeb.MethodDescriptor(
    '/neoshowcase.protobuf.AppsService/GetApps',
    grpcWeb.MethodType.UNARY,
    google_protobuf_empty_pb.Empty,
    neoshowcase_protobuf_apiserver_pb.GetAppsResponse,
    (request: google_protobuf_empty_pb.Empty) => {
      return request.serializeBinary();
    },
    neoshowcase_protobuf_apiserver_pb.GetAppsResponse.deserializeBinary
  );

  getApps(
    request: google_protobuf_empty_pb.Empty,
    metadata: grpcWeb.Metadata | null): Promise<neoshowcase_protobuf_apiserver_pb.GetAppsResponse>;

  getApps(
    request: google_protobuf_empty_pb.Empty,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: neoshowcase_protobuf_apiserver_pb.GetAppsResponse) => void): grpcWeb.ClientReadableStream<neoshowcase_protobuf_apiserver_pb.GetAppsResponse>;

  getApps(
    request: google_protobuf_empty_pb.Empty,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: neoshowcase_protobuf_apiserver_pb.GetAppsResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/neoshowcase.protobuf.AppsService/GetApps',
        request,
        metadata || {},
        this.methodDescriptorGetApps,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/neoshowcase.protobuf.AppsService/GetApps',
    request,
    metadata || {},
    this.methodDescriptorGetApps);
  }

}
