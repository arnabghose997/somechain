package somechain_test

import (
	"testing"

	keepertest "github.com/arnabghose997/somechain/testutil/keeper"
	"github.com/arnabghose997/somechain/testutil/nullify"
	"github.com/arnabghose997/somechain/x/somechain"
	"github.com/arnabghose997/somechain/x/somechain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SomechainKeeper(t)
	somechain.InitGenesis(ctx, *k, genesisState)
	got := somechain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
