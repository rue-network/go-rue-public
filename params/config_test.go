// Copyright 2017 The go-ruereum Authors
// This file is part of the go-ruereum library.
//
// The go-ruereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ruereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ruereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import (
	"math/big"
	"reflect"
	"testing"
)

func TestCheckCompatible(t *testing.T) {
	type test struct {
		stored, new *ChainConfig
		head        uint64
		wantErr     *ConfigCompatError
	}
	tests := []test{
		{stored: AllRuehashProtocolChanges, new: AllRuehashProtocolChanges, head: 0, wantErr: nil},
		{stored: AllRuehashProtocolChanges, new: AllRuehashProtocolChanges, head: 100, wantErr: nil},
		{
			stored:  &ChainConfig{EIP150Block: big.NewInt(10)},
			new:     &ChainConfig{EIP150Block: big.NewInt(20)},
			head:    9,
			wantErr: nil,
		},
		{
			stored: AllRuehashProtocolChanges,
			new:    &ChainConfig{HorizonBlock: nil},
			head:   3,
			wantErr: &ConfigCompatError{
				What:         "Horizon fork block",
				StoredConfig: big.NewInt(0),
				NewConfig:    nil,
				RewindTo:     0,
			},
		},
		{
			stored: AllRuehashProtocolChanges,
			new:    &ChainConfig{HorizonBlock: big.NewInt(1)},
			head:   3,
			wantErr: &ConfigCompatError{
				What:         "Horizon fork block",
				StoredConfig: big.NewInt(0),
				NewConfig:    big.NewInt(1),
				RewindTo:     0,
			},
		},
		{
			stored: &ChainConfig{HorizonBlock: big.NewInt(30), EIP150Block: big.NewInt(10)},
			new:    &ChainConfig{HorizonBlock: big.NewInt(25), EIP150Block: big.NewInt(20)},
			head:   25,
			wantErr: &ConfigCompatError{
				What:         "EIP150 fork block",
				StoredConfig: big.NewInt(10),
				NewConfig:    big.NewInt(20),
				RewindTo:     9,
			},
		},
	}

	for _, test := range tests {
		err := test.stored.CheckCompatible(test.new, test.head)
		if !reflect.DeepEqual(err, test.wantErr) {
			t.Errorf("error mismatch:\nstored: %v\nnew: %v\nhead: %v\nerr: %v\nwant: %v", test.stored, test.new, test.head, err, test.wantErr)
		}
	}
}
