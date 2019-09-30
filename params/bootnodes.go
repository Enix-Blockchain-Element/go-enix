// Copyright 2015 The go-ethereum Authors
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

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	"enode://c7bdcc03addd2a7791dbc28277659f7a5bdcf03dccfb63e6f2a7129c9aec4dae199143b1547f2cee55184c5dcc20ccfa5d781d0b440c77c47812951f98883fc8@157.245.204.116:30303",
	"enode://8192f9e5cfa8fb73e914c39e830211ab2b98d6180df1cf009c1b1c9ddad3bb6d09d986db5ee125f4bb59c226014e72975f170c734774cc389ccb6e6f73fcce34@174.138.19.97:1234",
	"enode://a27129508c3ccf2f23a093312f2c5bc951c2b7c0a1363ccae33cea783b684afc01a8a5251c6ae20bcbf37c6aae5835ce12a006575b85f7fff386a6ea36accad1@157.245.194.45:1234",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{}

// GoerliBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Görli test network.
var GoerliBootnodes = []string{}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{}
