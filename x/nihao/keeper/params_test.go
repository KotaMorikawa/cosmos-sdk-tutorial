package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "hello/testutil/keeper"
	"hello/x/nihao/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.NihaoKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
