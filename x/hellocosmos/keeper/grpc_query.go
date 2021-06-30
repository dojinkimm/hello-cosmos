package keeper

import (
	"github.com/dojinkimm/hello-cosmos/x/hellocosmos/types"
)

var _ types.QueryServer = Keeper{}
