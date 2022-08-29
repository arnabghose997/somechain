package keeper

import (
	"github.com/arnabghose997/somechain/x/somechain/types"
)

var _ types.QueryServer = Keeper{}
