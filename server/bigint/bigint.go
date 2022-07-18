package bigint

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
)

type BigInt struct {
	big.Int
}

func (b BigInt) MarshalJSON() ([]byte, error) {
	return []byte(b.String()), nil
}

func (b *BigInt) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return nil
	}
	if _, ok := b.SetString(s, 10); ok {
		return nil
	}
	return fmt.Errorf("not a valid bigint: %s", data)
}

func (b *BigInt) Unwrap() *big.Int {
	return &b.Int
}

func FromString(s string) (*BigInt, error) {
	n := new(BigInt)
	if _, ok := n.SetString(s, 10); !ok {
		return nil, errors.New(fmt.Sprintf("not a valid number: %s", s))
	}
	return n, nil
}

func Wrap(b *big.Int) *BigInt {
	return &BigInt{Int: *b}
}
