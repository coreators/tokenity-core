package keeper

import (
	"github.com/coreators/tokenity/x/blog/types"
)

var _ types.QueryServer = Keeper{}
