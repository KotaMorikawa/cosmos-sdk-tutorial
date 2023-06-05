package hello_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "hello/testutil/keeper"
	"hello/testutil/nullify"
	"hello/x/hello"
	"hello/x/hello/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		PeopleList: []types.People{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		PeopleCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.HelloKeeper(t)
	hello.InitGenesis(ctx, *k, genesisState)
	got := hello.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.PeopleList, got.PeopleList)
	require.Equal(t, genesisState.PeopleCount, got.PeopleCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
