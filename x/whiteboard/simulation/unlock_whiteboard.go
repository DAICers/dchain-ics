package simulation

import (
	"math/rand"

	"github.com/DAICers/dchain-ics/x/whiteboard/keeper"
	"github.com/DAICers/dchain-ics/x/whiteboard/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgUnlockWhiteboard(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgUnlockWhiteboard{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the UnlockWhiteboard simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "UnlockWhiteboard simulation not implemented"), nil, nil
	}
}
