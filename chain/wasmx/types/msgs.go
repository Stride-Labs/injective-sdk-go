package types

import (
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const RouterKey = ModuleName

func (msg MsgExecuteContractCompat) Route() string {
	return RouterKey
}

func (msg MsgExecuteContractCompat) Type() string {
	return "executeContractCompat"
}

func (msg MsgExecuteContractCompat) ValidateBasic() error {
	funds, err := sdk.ParseCoinsNormalized(msg.Funds)
	if err != nil {
		return err
	}

	oMsg := &wasmtypes.MsgExecuteContract{
		Sender:   msg.Sender,
		Contract: msg.Contract,
		Msg:      []byte(msg.Msg),
		Funds:    funds,
	}

	if err := oMsg.ValidateBasic(); err != nil {
		return err
	}
	return nil
}

func (msg MsgExecuteContractCompat) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgExecuteContractCompat) GetSigners() []sdk.AccAddress {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil { // should never happen as valid basic rejects invalid addresses
		panic(err.Error())
	}
	return []sdk.AccAddress{senderAddr}
}
