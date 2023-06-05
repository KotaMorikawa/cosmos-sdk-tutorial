package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"hello/x/hello/types"
)

func (k msgServer) Greeting(goCtx context.Context, msg *types.MsgGreeting) (*types.MsgGreetingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgGreetingResponse{}, nil
}
