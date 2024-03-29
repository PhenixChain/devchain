// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package relayd

import (
	"fmt"

	"github.com/PhenixChain/devchain/common/db"
	"github.com/PhenixChain/devchain/types"
	ty "github.com/PhenixChain/devchain/plugin/dapp/relay/types"
)

// Store hash and blockHeader
// SPV information

// Namespace keys
var (
	blockHashPrefix = []byte("blockHash")
	heightPrefix    = []byte("height")
	orderPrefix     = []byte("order")
)

type relaydDB struct {
	db db.DB
}

func newRelayDB(name string, dir string, cache int32) *relaydDB {
	d := db.NewDB(name, "goleveldb", dir, cache)
	return &relaydDB{d}
}

func (r *relaydDB) Get(key []byte) ([]byte, error) {
	return r.db.Get(key)
}

func (r *relaydDB) Set(key, value []byte) error {
	return r.db.Set(key, value)
}

func (r *relaydDB) queryOrderByHash(hash []byte) ([]byte, error) {
	return r.db.Get(append(orderPrefix, hash...))
}

func (r *relaydDB) queryOrderByStatus(status uint64) [][]byte {
	iter := r.db.Iterator(append(orderPrefix, fmt.Sprintf("%d", status)...), nil, false)
	var orders [][]byte
	for {
		if iter.Next() {
			orders = append(orders, iter.Value())
		} else {
			break
		}
	}
	return orders
}

func (r *relaydDB) BlockHeader(value interface{}) (*ty.BtcHeader, error) {
	var key []byte
	switch val := value.(type) {
	case uint64:
		key = makeHeightKey(val)
	case []byte:
		key = makeBlockHashKey(val)
	default:
		panic(val)
	}
	var ret ty.BtcHeader
	data, err := r.Get(key)
	if err != nil {
		return nil, err
	}
	err = types.Decode(data, &ret)
	return &ret, err
}

func makeHeightKey(height uint64) []byte {
	return append(heightPrefix, []byte(fmt.Sprintf("%d", height))...)
}

func makeBlockHashKey(hash []byte) []byte {
	return append(blockHashPrefix, hash...)
}
