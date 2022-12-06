package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyMaxXSize            = []byte("MaxXSize")
	DefaultMaxXSize uint64 = 1000
)

var (
	KeyMaxYSize            = []byte("MaxYSize")
	DefaultMaxYSize uint64 = 1000
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	maxXSize uint64,
	maxYSize uint64,
) Params {
	return Params{
		MaxXSize: maxXSize,
		MaxYSize: maxYSize,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultMaxXSize,
		DefaultMaxYSize,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMaxXSize, &p.MaxXSize, validateMaxXSize),
		paramtypes.NewParamSetPair(KeyMaxYSize, &p.MaxYSize, validateMaxYSize),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateMaxXSize(p.MaxXSize); err != nil {
		return err
	}

	if err := validateMaxYSize(p.MaxYSize); err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// validateMaxXSize validates the MaxXSize param
func validateMaxXSize(v interface{}) error {
	maxXSize, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = maxXSize

	return nil
}

// validateMaxYSize validates the MaxYSize param
func validateMaxYSize(v interface{}) error {
	maxYSize, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = maxYSize

	return nil
}
