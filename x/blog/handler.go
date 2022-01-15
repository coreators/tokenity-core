package blog

import (
	"fmt"

	"github.com/coreators/tokenity/x/blog/keeper"
	"github.com/coreators/tokenity/x/blog/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gogo/protobuf/proto"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		var (
			res proto.Message
			err error
		)
		switch msg := msg.(type) {
		case *types.MsgCreatePost:
			res, err = msgServer.CreatePost(sdk.WrapSDKContext(ctx), msg)
		case *types.MsgUpdatePost:
			res, err = msgServer.UpdatePost(sdk.WrapSDKContext(ctx), msg)
		case *types.MsgDeletePost:
			res, err = msgServer.DeletePost(sdk.WrapSDKContext(ctx), msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
		return sdk.WrapServiceResult(ctx, res, err)
	}
}
