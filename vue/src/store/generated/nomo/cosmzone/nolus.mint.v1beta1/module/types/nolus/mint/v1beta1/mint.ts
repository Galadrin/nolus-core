/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "nolus.mint.v1beta1";

/** Minter represents the minting state. */
export interface Minter {
  norm_time_passed: string;
  total_minted: string;
  prev_block_timestamp: number;
}

export interface Params {
  /** type of coin to mint */
  mint_denom: string;
  max_mintable_nanoseconds: number;
}

const baseMinter: object = {
  norm_time_passed: "",
  total_minted: "",
  prev_block_timestamp: 0,
};

export const Minter = {
  encode(message: Minter, writer: Writer = Writer.create()): Writer {
    if (message.norm_time_passed !== "") {
      writer.uint32(18).string(message.norm_time_passed);
    }
    if (message.total_minted !== "") {
      writer.uint32(26).string(message.total_minted);
    }
    if (message.prev_block_timestamp !== 0) {
      writer.uint32(32).int64(message.prev_block_timestamp);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Minter {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMinter } as Minter;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 2:
          message.norm_time_passed = reader.string();
          break;
        case 3:
          message.total_minted = reader.string();
          break;
        case 4:
          message.prev_block_timestamp = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Minter {
    const message = { ...baseMinter } as Minter;
    if (
      object.norm_time_passed !== undefined &&
      object.norm_time_passed !== null
    ) {
      message.norm_time_passed = String(object.norm_time_passed);
    } else {
      message.norm_time_passed = "";
    }
    if (object.total_minted !== undefined && object.total_minted !== null) {
      message.total_minted = String(object.total_minted);
    } else {
      message.total_minted = "";
    }
    if (
      object.prev_block_timestamp !== undefined &&
      object.prev_block_timestamp !== null
    ) {
      message.prev_block_timestamp = Number(object.prev_block_timestamp);
    } else {
      message.prev_block_timestamp = 0;
    }
    return message;
  },

  toJSON(message: Minter): unknown {
    const obj: any = {};
    message.norm_time_passed !== undefined &&
      (obj.norm_time_passed = message.norm_time_passed);
    message.total_minted !== undefined &&
      (obj.total_minted = message.total_minted);
    message.prev_block_timestamp !== undefined &&
      (obj.prev_block_timestamp = message.prev_block_timestamp);
    return obj;
  },

  fromPartial(object: DeepPartial<Minter>): Minter {
    const message = { ...baseMinter } as Minter;
    if (
      object.norm_time_passed !== undefined &&
      object.norm_time_passed !== null
    ) {
      message.norm_time_passed = object.norm_time_passed;
    } else {
      message.norm_time_passed = "";
    }
    if (object.total_minted !== undefined && object.total_minted !== null) {
      message.total_minted = object.total_minted;
    } else {
      message.total_minted = "";
    }
    if (
      object.prev_block_timestamp !== undefined &&
      object.prev_block_timestamp !== null
    ) {
      message.prev_block_timestamp = object.prev_block_timestamp;
    } else {
      message.prev_block_timestamp = 0;
    }
    return message;
  },
};

const baseParams: object = { mint_denom: "", max_mintable_nanoseconds: 0 };

export const Params = {
  encode(message: Params, writer: Writer = Writer.create()): Writer {
    if (message.mint_denom !== "") {
      writer.uint32(10).string(message.mint_denom);
    }
    if (message.max_mintable_nanoseconds !== 0) {
      writer.uint32(16).int64(message.max_mintable_nanoseconds);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Params {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseParams } as Params;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.mint_denom = reader.string();
          break;
        case 2:
          message.max_mintable_nanoseconds = longToNumber(
            reader.int64() as Long
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Params {
    const message = { ...baseParams } as Params;
    if (object.mint_denom !== undefined && object.mint_denom !== null) {
      message.mint_denom = String(object.mint_denom);
    } else {
      message.mint_denom = "";
    }
    if (
      object.max_mintable_nanoseconds !== undefined &&
      object.max_mintable_nanoseconds !== null
    ) {
      message.max_mintable_nanoseconds = Number(
        object.max_mintable_nanoseconds
      );
    } else {
      message.max_mintable_nanoseconds = 0;
    }
    return message;
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    message.mint_denom !== undefined && (obj.mint_denom = message.mint_denom);
    message.max_mintable_nanoseconds !== undefined &&
      (obj.max_mintable_nanoseconds = message.max_mintable_nanoseconds);
    return obj;
  },

  fromPartial(object: DeepPartial<Params>): Params {
    const message = { ...baseParams } as Params;
    if (object.mint_denom !== undefined && object.mint_denom !== null) {
      message.mint_denom = object.mint_denom;
    } else {
      message.mint_denom = "";
    }
    if (
      object.max_mintable_nanoseconds !== undefined &&
      object.max_mintable_nanoseconds !== null
    ) {
      message.max_mintable_nanoseconds = object.max_mintable_nanoseconds;
    } else {
      message.max_mintable_nanoseconds = 0;
    }
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

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

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
