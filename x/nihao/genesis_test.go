package nihao_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "hello/testutil/keeper"
	"hello/testutil/nullify"
	"hello/x/nihao"
	"hello/x/nihao/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		TaskList: []types.Task{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		TaskCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.NihaoKeeper(t)
	nihao.InitGenesis(ctx, *k, genesisState)
	got := nihao.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.TaskList, got.TaskList)
	require.Equal(t, genesisState.TaskCount, got.TaskCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
