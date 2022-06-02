/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../../../nolus/mint/v1beta1/mint";

export const protobufPackage = "nolus.mint.v1beta1";

/** QueryParamsRequest is the request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is the response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params defines the parameters of the module. */
  params: Params | undefined;
}

/** QueryMintStateRequest is the request type for the Query/State RPC method. */
export interface QueryMintStateRequest {}

/**
 * QueryMintStateResponse is the response type for the Query/State RPC
 * method.
 */
export interface QueryMintStateResponse {
  norm_time_passed: Uint8Array;
  total_minted: Uint8Array;
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryMintStateRequest: object = {};

export const QueryMintStateRequest = {
  encode(_: QueryMintStateRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryMintStateRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryMintStateRequest } as QueryMintStateRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryMintStateRequest {
    const message = { ...baseQueryMintStateRequest } as QueryMintStateRequest;
    return message;
  },

  toJSON(_: QueryMintStateRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryMintStateRequest>): QueryMintStateRequest {
    const message = { ...baseQueryMintStateRequest } as QueryMintStateRequest;
    return message;
  },
};

const baseQueryMintStateResponse: object = {};

export const QueryMintStateResponse = {
  encode(
    message: QueryMintStateResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.norm_time_passed.length !== 0) {
      writer.uint32(10).bytes(message.norm_time_passed);
    }
    if (message.total_minted.length !== 0) {
      writer.uint32(18).bytes(message.total_minted);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryMintStateResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryMintStateResponse } as QueryMintStateResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.norm_time_passed = reader.bytes();
          break;
        case 2:
          message.total_minted = reader.bytes();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryMintStateResponse {
    const message = { ...baseQueryMintStateResponse } as QueryMintStateResponse;
    if (
      object.norm_time_passed !== undefined &&
      object.norm_time_passed !== null
    ) {
      message.norm_time_passed = bytesFromBase64(object.norm_time_passed);
    }
    if (object.total_minted !== undefined && object.total_minted !== null) {
      message.total_minted = bytesFromBase64(object.total_minted);
    }
    return message;
  },

  toJSON(message: QueryMintStateResponse): unknown {
    const obj: any = {};
    message.norm_time_passed !== undefined &&
      (obj.norm_time_passed = base64FromBytes(
        message.norm_time_passed !== undefined
          ? message.norm_time_passed
          : new Uint8Array()
      ));
    message.total_minted !== undefined &&
      (obj.total_minted = base64FromBytes(
        message.total_minted !== undefined
          ? message.total_minted
          : new Uint8Array()
      ));
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryMintStateResponse>
  ): QueryMintStateResponse {
    const message = { ...baseQueryMintStateResponse } as QueryMintStateResponse;
    if (
      object.norm_time_passed !== undefined &&
      object.norm_time_passed !== null
    ) {
      message.norm_time_passed = object.norm_time_passed;
    } else {
      message.norm_time_passed = new Uint8Array();
    }
    if (object.total_minted !== undefined && object.total_minted !== null) {
      message.total_minted = object.total_minted;
    } else {
      message.total_minted = new Uint8Array();
    }
    return message;
  },
};

/** Query provides defines the gRPC querier service. */
export interface Query {
  /** Params returns the total set of minting parameters. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** MintState returns the current minting state value. */
  MintState(request: QueryMintStateRequest): Promise<QueryMintStateResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "nolus.mint.v1beta1.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  MintState(request: QueryMintStateRequest): Promise<QueryMintStateResponse> {
    const data = QueryMintStateRequest.encode(request).finish();
    const promise = this.rpc.request(
      "nolus.mint.v1beta1.Query",
      "MintState",
      data
    );
    return promise.then((data) =>
      QueryMintStateResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

const atob: (b64: string) => string =
  globalThis.atob ||
  ((b64) => globalThis.Buffer.from(b64, "base64").toString("binary"));
function bytesFromBase64(b64: string): Uint8Array {
  const bin = atob(b64);
  const arr = new Uint8Array(bin.length);
  for (let i = 0; i < bin.length; ++i) {
    arr[i] = bin.charCodeAt(i);
  }
  return arr;
}

const btoa: (bin: string) => string =
  globalThis.btoa ||
  ((bin) => globalThis.Buffer.from(bin, "binary").toString("base64"));
function base64FromBytes(arr: Uint8Array): string {
  const bin: string[] = [];
  for (let i = 0; i < arr.byteLength; ++i) {
    bin.push(String.fromCharCode(arr[i]));
  }
  return btoa(bin.join(""));
}

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
