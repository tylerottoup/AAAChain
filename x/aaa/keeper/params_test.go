package keeper_test

import (
	"testing"

	testkeeper "github.com/beeust/AAA/testutil/keeper"
	"github.com/beeust/AAA/x/aaa/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.AaaKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
