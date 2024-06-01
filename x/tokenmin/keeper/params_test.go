package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/harrybite/tokenmin/testutil/keeper"
	"github.com/harrybite/tokenmin/x/tokenmin/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.TokenminKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
