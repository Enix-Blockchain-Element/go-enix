// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package core

// Constants containing the genesis allocation of built-in genesis blocks.
// Their content is an RLP-encoded list of (address, balance) tuples.
// Use mkalloc.go to create/update them.

// nolint: misspell
const mainnetAllocData = "\xe2\xe1\x94\xfa\xc3OQ\x02\xa1\xbb\xf0\xc47>\xc0]\xb6V<r\xbe\xe3\xfd\x8b\x18\u043fB<\x03\xd8\xde\x00\x00\x00"
const testnetAllocData = ""
const rinkebyAllocData = ""
const goerliAllocData = ""
