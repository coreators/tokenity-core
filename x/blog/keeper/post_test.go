package keeper_test

import (
	"testing"

	keepertest "github.com/coreators/tokenity/testutil/keeper"
	"github.com/coreators/tokenity/testutil/nullify"
	"github.com/coreators/tokenity/x/blog/keeper"
	"github.com/coreators/tokenity/x/blog/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNPost(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Post {
	items := make([]types.Post, n)
	for i := range items {
		items[i].Id = keeper.AppendPost(ctx, items[i])
	}
	return items
}

func TestPostGet(t *testing.T) {
	keeper, ctx := keepertest.BlogKeeper(t)
	items := createNPost(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetPost(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestPostRemove(t *testing.T) {
	keeper, ctx := keepertest.BlogKeeper(t)
	items := createNPost(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePost(ctx, item.Id)
		_, found := keeper.GetPost(ctx, item.Id)
		require.False(t, found)
	}
}

func TestPostGetAll(t *testing.T) {
	keeper, ctx := keepertest.BlogKeeper(t)
	items := createNPost(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPost(ctx)),
	)
}

func TestPostCount(t *testing.T) {
	keeper, ctx := keepertest.BlogKeeper(t)
	items := createNPost(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetPostCount(ctx))
}
