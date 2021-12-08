import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "coreators.tokenity.blog";
/** proto/blog/comment.proto */
export interface Comment {
    id: string;
    creator: string;
    content: string;
}
export interface MsgCreateComment {
    creator: string;
    content: string;
}
export declare const Comment: {
    encode(message: Comment, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Comment;
    fromJSON(object: any): Comment;
    toJSON(message: Comment): unknown;
    fromPartial(object: DeepPartial<Comment>): Comment;
};
export declare const MsgCreateComment: {
    encode(message: MsgCreateComment, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateComment;
    fromJSON(object: any): MsgCreateComment;
    toJSON(message: MsgCreateComment): unknown;
    fromPartial(object: DeepPartial<MsgCreateComment>): MsgCreateComment;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
