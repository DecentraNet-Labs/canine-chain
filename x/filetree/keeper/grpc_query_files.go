package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jackalLabs/canine-chain/v4/x/filetree/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// To remove
func (k Keeper) AllFiles(c context.Context, req *types.QueryAllFiles) (*types.QueryAllFilesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var filess []types.Files
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	filesStore := prefix.NewStore(store, types.KeyPrefix(types.FilesKeyPrefix))

	iterator := filesStore.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		value := iterator.Value()
		// Check if the key ends with the owner address
		if len(key) >= len(req.OwnerAddress) && string(key[len(key)-len(req.OwnerAddress):]) == req.OwnerAddress {
			// Extract the filepath from the key
			var files types.Files
			if err := k.cdc.Unmarshal(value, &files); err != nil {
				return nil, status.Error(codes.Internal, err.Error())
			}

		filess = append(filess, files)
	}

	pageRes, err := query.Paginate(filesStore, req.Pagination, func(_ []byte, value []byte) error {
		return nil
	})
	
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllFilesResponse{Files: filess, Pagination: pageRes}, nil
}

func (k Keeper) File(c context.Context, req *types.QueryFile) (*types.QueryFileResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetFiles(
		ctx,
		req.Address,
		req.OwnerAddress)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryFileResponse{File: val}, nil
}
