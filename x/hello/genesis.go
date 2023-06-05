package hello

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"hello/x/hello/keeper"
	"hello/x/hello/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the people
	for _, elem := range genState.PeopleList {
		k.SetPeople(ctx, elem)
	}

	// Set people count
	k.SetPeopleCount(ctx, genState.PeopleCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.PeopleList = k.GetAllPeople(ctx)
	genesis.PeopleCount = k.GetPeopleCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
