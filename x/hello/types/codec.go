package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgGreeting{}, "hello/Greeting", nil)
	cdc.RegisterConcrete(&MsgCreatePeople{}, "hello/CreatePeople", nil)
	cdc.RegisterConcrete(&MsgUpdatePeople{}, "hello/UpdatePeople", nil)
	cdc.RegisterConcrete(&MsgDeletePeople{}, "hello/DeletePeople", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgGreeting{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreatePeople{},
		&MsgUpdatePeople{},
		&MsgDeletePeople{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
