package types

import (
	"net/url"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/pkg/errors"
)

const (
	MsgTypeCreateProvider = "create-provider"
	MsgTypeUpdateProvider = "update-provider"
	MsgTypeDeleteProvider = "delete-provider"
)

var (
	_, _, _ sdk.Msg = &MsgCreateProvider{}, &MsgUpdateProvider{}, &MsgDeleteProvider{}
)

// MsgCreateProvider defines an SDK message for creating a provider
// type MsgCreateProvider Provider

// NewMsgCreateProvider creates a new MsgCreateProvider instance
func NewMsgCreateProvider(owner sdk.AccAddress, hostURI string, attributes Attributes) *MsgCreateProvider {
	return &MsgCreateProvider{
		Owner:      owner,
		HostURI:    hostURI,
		Attributes: attributes,
	}
}

// Route implements the sdk.Msg interface
func (msg MsgCreateProvider) Route() string { return RouterKey }

// Type implements the sdk.Msg interface
func (msg MsgCreateProvider) Type() string { return MsgTypeCreateProvider }

// ValidateBasic does basic validation of a HostURI
func (msg MsgCreateProvider) ValidateBasic() error {
	if err := validateProviderURI(msg.HostURI); err != nil {
		return err
	}
	if err := sdk.VerifyAddressFormat(msg.Owner); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "MsgCreate: Invalid Provider Address")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgCreateProvider) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgCreateProvider) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

// ToProvider returns Provider from msg
func (msg MsgCreateProvider) ToProvider() Provider {
	return Provider{
		Owner:      msg.Owner,
		HostURI:    msg.HostURI,
		Attributes: msg.Attributes,
	}
}

// MsgUpdateProvider defines an SDK message for updating a provider
// type MsgUpdateProvider Provider

// NewMsgUpdateProvider creates a new MsgUpdateProvider instance
func NewMsgUpdateProvider(owner sdk.AccAddress, hostURI string, attributes Attributes) *MsgUpdateProvider {
	return &MsgUpdateProvider{
		Owner:      owner,
		HostURI:    hostURI,
		Attributes: attributes,
	}
}

// Route implements the sdk.Msg interface
func (msg MsgUpdateProvider) Route() string { return RouterKey }

// Type implements the sdk.Msg interface
func (msg MsgUpdateProvider) Type() string { return MsgTypeUpdateProvider }

// ValidateBasic does basic validation of a ProviderURI
func (msg MsgUpdateProvider) ValidateBasic() error {
	if err := validateProviderURI(msg.HostURI); err != nil {
		return err
	}
	if err := sdk.VerifyAddressFormat(msg.Owner); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "MsgUpdate: Invalid Provider Address")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgUpdateProvider) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgUpdateProvider) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

// ToProvider returns Provider from msg
func (msg MsgUpdateProvider) ToProvider() Provider {
	return Provider{
		Owner:      msg.Owner,
		HostURI:    msg.HostURI,
		Attributes: msg.Attributes,
	}
}

// MsgDeleteProvider defines an SDK message for deleting a provider
// type MsgDeleteProvider struct {
// 	Owner sdk.AccAddress `json:"owner"`
// }

// NewMsgDeleteProvider creates a new MsgDeleteProvider instance
func NewMsgDeleteProvider(owner sdk.AccAddress) *MsgDeleteProvider {
	return &MsgDeleteProvider{
		Owner: owner,
	}
}

// Route implements the sdk.Msg interface
func (msg MsgDeleteProvider) Route() string { return RouterKey }

// Type implements the sdk.Msg interface
func (msg MsgDeleteProvider) Type() string { return MsgTypeDeleteProvider }

// ValidateBasic does basic validation
func (msg MsgDeleteProvider) ValidateBasic() error {
	if err := sdk.VerifyAddressFormat(msg.Owner); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "MsgDelete: Invalid Provider Address")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgDeleteProvider) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgDeleteProvider) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

func validateProviderURI(val string) error {
	u, err := url.Parse(val)
	if err != nil {
		return ErrInvalidProviderURI
	}
	if !u.IsAbs() {
		return errors.Wrapf(ErrNotAbsProviderURI, "validating %q for absolute URI", val)
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return errors.Wrapf(ErrInvalidProviderURI, "scheme in %q should be http or https", val)
	}

	if u.Host == "" {
		return errors.Wrapf(ErrInvalidProviderURI, "validating %q for valid host", val)
	}

	if u.Path != "" {
		return errors.Wrapf(ErrInvalidProviderURI, "path in %q should be empty", val)
	}

	return nil
}
