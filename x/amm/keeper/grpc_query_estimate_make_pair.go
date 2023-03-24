package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackalLabs/canine-chain/x/amm/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) EstimateContribution(
	goCtx context.Context,
	req *types.QueryEstimateContributionRequest,
) (
	*types.QueryEstimateContributionResponse,
	error,
) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	pool, found := k.GetPool(ctx, req.PoolName)

	if !found {
		return nil, types.ErrLiquidityPoolNotFound
	}

	desiredAmt, ok := sdk.NewIntFromString(req.DesiredAmount)

	if !ok {
		return nil, sdkerrors.Wrapf(
			sdkerrors.ErrInvalidRequest,
			"cannot convert desired amount into type sdk.Int: %s",
			req.DesiredAmount,
		)
	}

	result, err := CoinsToDepositForPoolToken(pool, desiredAmt)

	if err != nil {
		return nil, sdkerrors.Wrapf(
			sdkerrors.ErrInvalidRequest,
			"Failed to calculate deposit coins for pool token",
		)
	}

	return &types.QueryEstimateContributionResponse{Coins: result}, nil
}
