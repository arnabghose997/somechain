package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/arnabghose997/somechain/testutil/keeper"
	"github.com/arnabghose997/somechain/x/somechain/keeper"
	"github.com/arnabghose997/somechain/x/somechain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.SomechainKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
