package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/coreators/tokenity/testutil/keeper"
	"github.com/coreators/tokenity/x/blog/keeper"
	"github.com/coreators/tokenity/x/blog/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.BlogKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
