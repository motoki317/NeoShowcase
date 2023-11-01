// @generated by protoc-gen-es v1.4.1 with parameter "target=ts"
// @generated from file neoshowcase/protobuf/null.proto (package neoshowcase.protobuf, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, Timestamp } from "@bufbuild/protobuf";

/**
 * @generated from message neoshowcase.protobuf.NullTimestamp
 */
export class NullTimestamp extends Message<NullTimestamp> {
  /**
   * @generated from field: google.protobuf.Timestamp timestamp = 1;
   */
  timestamp?: Timestamp;

  /**
   * @generated from field: bool valid = 2;
   */
  valid = false;

  constructor(data?: PartialMessage<NullTimestamp>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "neoshowcase.protobuf.NullTimestamp";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "timestamp", kind: "message", T: Timestamp },
    { no: 2, name: "valid", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): NullTimestamp {
    return new NullTimestamp().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): NullTimestamp {
    return new NullTimestamp().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): NullTimestamp {
    return new NullTimestamp().fromJsonString(jsonString, options);
  }

  static equals(a: NullTimestamp | PlainMessage<NullTimestamp> | undefined, b: NullTimestamp | PlainMessage<NullTimestamp> | undefined): boolean {
    return proto3.util.equals(NullTimestamp, a, b);
  }
}

