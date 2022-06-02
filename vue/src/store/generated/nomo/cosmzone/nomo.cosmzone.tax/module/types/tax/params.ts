/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "nomo.cosmzone.tax";

/** Params defines the parameters for the module. */
export interface Params {
  feeRate: number;
  feeCaps: string;
  contractAddress: string;
}

const baseParams: object = { feeRate: 0, feeCaps: "", contractAddress: "" };

export const Params = {
  encode(message: Params, writer: Writer = Writer.create()): Writer {
    if (message.feeRate !== 0) {
      writer.uint32(8).int32(message.feeRate);
    }
    if (message.feeCaps !== "") {
      writer.uint32(18).string(message.feeCaps);
    }
    if (message.contractAddress !== "") {
      writer.uint32(26).string(message.contractAddress);
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
          message.feeRate = reader.int32();
          break;
        case 2:
          message.feeCaps = reader.string();
          break;
        case 3:
          message.contractAddress = reader.string();
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
    if (object.feeRate !== undefined && object.feeRate !== null) {
      message.feeRate = Number(object.feeRate);
    } else {
      message.feeRate = 0;
    }
    if (object.feeCaps !== undefined && object.feeCaps !== null) {
      message.feeCaps = String(object.feeCaps);
    } else {
      message.feeCaps = "";
    }
    if (
      object.contractAddress !== undefined &&
      object.contractAddress !== null
    ) {
      message.contractAddress = String(object.contractAddress);
    } else {
      message.contractAddress = "";
    }
    return message;
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    message.feeRate !== undefined && (obj.feeRate = message.feeRate);
    message.feeCaps !== undefined && (obj.feeCaps = message.feeCaps);
    message.contractAddress !== undefined &&
      (obj.contractAddress = message.contractAddress);
    return obj;
  },

  fromPartial(object: DeepPartial<Params>): Params {
    const message = { ...baseParams } as Params;
    if (object.feeRate !== undefined && object.feeRate !== null) {
      message.feeRate = object.feeRate;
    } else {
      message.feeRate = 0;
    }
    if (object.feeCaps !== undefined && object.feeCaps !== null) {
      message.feeCaps = object.feeCaps;
    } else {
      message.feeCaps = "";
    }
    if (
      object.contractAddress !== undefined &&
      object.contractAddress !== null
    ) {
      message.contractAddress = object.contractAddress;
    } else {
      message.contractAddress = "";
    }
    return message;
  },
};

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
