package nihao

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"hello/testutil/sample"
	nihaosimulation "hello/x/nihao/simulation"
	"hello/x/nihao/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = nihaosimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateTask = "op_weight_msg_task"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateTask int = 100

	opWeightMsgUpdateTask = "op_weight_msg_task"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateTask int = 100

	opWeightMsgDeleteTask = "op_weight_msg_task"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteTask int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	nihaoGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		TaskList: []types.Task{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		TaskCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&nihaoGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateTask int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateTask, &weightMsgCreateTask, nil,
		func(_ *rand.Rand) {
			weightMsgCreateTask = defaultWeightMsgCreateTask
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateTask,
		nihaosimulation.SimulateMsgCreateTask(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateTask int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateTask, &weightMsgUpdateTask, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateTask = defaultWeightMsgUpdateTask
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateTask,
		nihaosimulation.SimulateMsgUpdateTask(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteTask int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteTask, &weightMsgDeleteTask, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteTask = defaultWeightMsgDeleteTask
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteTask,
		nihaosimulation.SimulateMsgDeleteTask(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
