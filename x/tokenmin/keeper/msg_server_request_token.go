package keeper

import (
	"context"

	"github.com/harrybite/tokenmin/x/tokenmin/types"

	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RequestToken(goCtx context.Context, msg *types.MsgRequestToken) (*types.MsgRequestTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	amount, err := sdk.ParseCoinsNormalized(msg.Amount)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrSample, "Cannot parse coins in loan amount")
	}

	sdkError := k.bankKeeper.MintCoins(ctx, types.ModuleName, amount)
	if sdkError != nil {
		return nil, sdkError
	}
	lender, _ := sdk.AccAddressFromBech32(msg.Creator)
	sdkError2 := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, lender, amount)
	if sdkError2 != nil {
		return nil, sdkError2
	}

	return &types.MsgRequestTokenResponse{}, nil
}
