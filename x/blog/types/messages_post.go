package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreatePost = "create_post"
	TypeMsgUpdatePost = "update_post"
	TypeMsgDeletePost = "delete_post"
)

var _ sdk.Msg = &MsgCreatePost{}

func NewMsgCreatePost(creator string, jsonContents string, description string, isStory bool) *MsgCreatePost {
	var contents []*Content

	err := json.Unmarshal([]byte(jsonContents), &contents)
	if err != nil {
		println("Contents json parsing error : ", err)
	}

	newPost := &MsgCreatePost{
		Creator:     creator,
		Contents:    contents,
		Description: description,
		IsStory:     isStory,
	}

	return newPost
}

func (msg *MsgCreatePost) Route() string {
	return RouterKey
}

func (msg *MsgCreatePost) Type() string {
	return TypeMsgCreatePost
}

func (msg *MsgCreatePost) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreatePost) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreatePost) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdatePost{}

func NewMsgUpdatePost(creator string, id uint64, jsonContents string, description string, isStory bool) *MsgUpdatePost {
	var contents []*Content
	err := json.Unmarshal([]byte(jsonContents), &contents)
	if err != nil {
		println("Contents json parsing error : ", err)
	}

	return &MsgUpdatePost{
		Id:          id,
		Creator:     creator,
		Contents:    contents,
		Description: description,
		IsStory:     isStory,
	}
}

func (msg *MsgUpdatePost) Route() string {
	return RouterKey
}

func (msg *MsgUpdatePost) Type() string {
	return TypeMsgUpdatePost
}

func (msg *MsgUpdatePost) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdatePost) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdatePost) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeletePost{}

func NewMsgDeletePost(creator string, id uint64) *MsgDeletePost {
	return &MsgDeletePost{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeletePost) Route() string {
	return RouterKey
}

func (msg *MsgDeletePost) Type() string {
	return TypeMsgDeletePost
}

func (msg *MsgDeletePost) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeletePost) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeletePost) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
