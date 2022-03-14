package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/beeust/AAA/testutil/keeper"
	"github.com/beeust/AAA/x/aaa/keeper"
	"github.com/beeust/AAA/x/aaa/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.AaaKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
