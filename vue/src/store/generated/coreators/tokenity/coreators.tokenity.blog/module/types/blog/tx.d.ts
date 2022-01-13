import { Reader, Writer } from "protobufjs/minimal";
export declare const protobufPackage = "coreators.tokenity.blog";
export interface MsgCreatePost {
    creator: string;
    contents: string;
    description: string;
}
export interface MsgCreatePostResponse {
    id: number;
}
export interface MsgUpdatePost {
    creator: string;
    id: number;
    contents: string;
    description: string;
}
export interface MsgUpdatePostResponse {
}
export interface MsgDeletePost {
    creator: string;
    id: number;
}
export interface MsgDeletePostResponse {
}
export declare const MsgCreatePost: {
    encode(message: MsgCreatePost, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreatePost;
    fromJSON(object: any): MsgCreatePost;
    toJSON(message: MsgCreatePost): unknown;
    fromPartial(object: DeepPartial<MsgCreatePost>): MsgCreatePost;
};
export declare const MsgCreatePostResponse: {
    encode(message: MsgCreatePostResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreatePostResponse;
    fromJSON(object: any): MsgCreatePostResponse;
    toJSON(message: MsgCreatePostResponse): unknown;
    fromPartial(object: DeepPartial<MsgCreatePostResponse>): MsgCreatePostResponse;
};
export declare const MsgUpdatePost: {
    encode(message: MsgUpdatePost, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdatePost;
    fromJSON(object: any): MsgUpdatePost;
    toJSON(message: MsgUpdatePost): unknown;
    fromPartial(object: DeepPartial<MsgUpdatePost>): MsgUpdatePost;
};
export declare const MsgUpdatePostResponse: {
    encode(_: MsgUpdatePostResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdatePostResponse;
    fromJSON(_: any): MsgUpdatePostResponse;
    toJSON(_: MsgUpdatePostResponse): unknown;
    fromPartial(_: DeepPartial<MsgUpdatePostResponse>): MsgUpdatePostResponse;
};
export declare const MsgDeletePost: {
    encode(message: MsgDeletePost, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeletePost;
    fromJSON(object: any): MsgDeletePost;
    toJSON(message: MsgDeletePost): unknown;
    fromPartial(object: DeepPartial<MsgDeletePost>): MsgDeletePost;
};
export declare const MsgDeletePostResponse: {
    encode(_: MsgDeletePostResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeletePostResponse;
    fromJSON(_: any): MsgDeletePostResponse;
    toJSON(_: MsgDeletePostResponse): unknown;
    fromPartial(_: DeepPartial<MsgDeletePostResponse>): MsgDeletePostResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    CreatePost(request: MsgCreatePost): Promise<MsgCreatePostResponse>;
    UpdatePost(request: MsgUpdatePost): Promise<MsgUpdatePostResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpc */
    DeletePost(request: MsgDeletePost): Promise<MsgDeletePostResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    CreatePost(request: MsgCreatePost): Promise<MsgCreatePostResponse>;
    UpdatePost(request: MsgUpdatePost): Promise<MsgUpdatePostResponse>;
    DeletePost(request: MsgDeletePost): Promise<MsgDeletePostResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
