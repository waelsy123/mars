package mars_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/waelsy123/mars/testutil/keeper"
	"github.com/waelsy123/mars/testutil/nullify"
	"github.com/waelsy123/mars/x/mars"
	"github.com/waelsy123/mars/x/mars/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MarsKeeper(t)
	mars.InitGenesis(ctx, *k, genesisState)
	got := mars.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
