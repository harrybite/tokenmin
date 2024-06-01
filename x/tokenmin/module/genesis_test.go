package tokenmin_test

import (
	"testing"

	keepertest "github.com/harrybite/tokenmin/testutil/keeper"
	"github.com/harrybite/tokenmin/testutil/nullify"
	tokenmin "github.com/harrybite/tokenmin/x/tokenmin/module"
	"github.com/harrybite/tokenmin/x/tokenmin/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TokenminKeeper(t)
	tokenmin.InitGenesis(ctx, k, genesisState)
	got := tokenmin.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
