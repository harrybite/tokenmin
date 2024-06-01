package keeper

import (
	"github.com/harrybite/tokenmin/x/tokenmin/types"
)

var _ types.QueryServer = Keeper{}
