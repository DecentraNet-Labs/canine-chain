package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgPostFile = "post_file"

var _ sdk.Msg = &MsgPostFile{}

func NewMsgPostFile(creator string, account string, hashparent string, hashchild string, contents string, viewers string, editors string,
	trackingNumber string,
) *MsgPostFile {
	return &MsgPostFile{
		Creator:        creator,
		Account:        account,
		HashParent:     hashparent,
		HashChild:      hashchild,
		Contents:       contents,
		Viewers:        viewers,
		Editors:        editors,
		TrackingNumber: trackingNumber,
	}
}

func (msg *MsgPostFile) Route() string {
	return RouterKey
}

func (msg *MsgPostFile) Type() string {
	return TypeMsgPostFile
}

func (msg *MsgPostFile) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgPostFile) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgPostFile) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.Account == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"invalid account: %s", msg.Account)
	}
	if msg.HashParent == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"invalid hash parent: %s", msg.HashParent)
	}
	if msg.HashChild == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"invalid hash child: %s", msg.HashChild)
	}
	if msg.Viewers == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"invalid viewers: %s", msg.Viewers)
	}
	if msg.Editors == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"invalid editors: %s", msg.Editors)
	}
	if msg.TrackingNumber == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"invalid tracking number: %s", msg.TrackingNumber)
	}

	return nil
}
