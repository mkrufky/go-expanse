// Copyright 2015 The go-expanse Authors
// This file is part of the go-expanse library.
//
// The go-expanse library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-expanse library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-expanse library. If not, see <http://www.gnu.org/licenses/>.

package core

import (
	"math/big"

	"github.com/expanse-project/go-expanse/common"
	"github.com/expanse-project/go-expanse/core/state"
)

var (
	lol     = common.HexToAddress("959c33de5961820567930eccce51ea715c496f85")
)

// Canary will check the 0'd address of the 4 contracts above.
// If two or more are set to anything other than a 0 the canary
// dies a horrible death.
func Canary(statedb *state.StateDB) bool {
	var r int
	if (statedb.GetState(lol, common.Hash{}).Big().Cmp(big.NewInt(0)) > 0) {
		
	}
	return r > 1
}
