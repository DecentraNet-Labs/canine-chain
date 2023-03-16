package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCancelContract_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCancelContract
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCancelContract{
				Creator: "invalid_address",
				Cid:     "jklc1j3p63s42w7ywaczlju626st55mzu5z39qh6g4g",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid cid",
			msg: MsgCancelContract{
				Creator: "jkl1j3p63s42w7ywaczlju626st55mzu5z399f5n6n",
				Cid:     "invalid_cid",
			},
			err: sdkerrors.ErrInvalidRequest,
		}, {
			name: "valid address",
			msg: MsgCancelContract{
				Creator: "jkl1j3p63s42w7ywaczlju626st55mzu5z399f5n6n",
				Cid:     "jklc1j3p63s42w7ywaczlju626st55mzu5z39qh6g4g",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
