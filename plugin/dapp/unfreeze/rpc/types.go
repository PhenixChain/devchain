// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rpc

import (
	"github.com/PhenixChain/devchain/rpc/types"
)

// Jrpc json rpc struct
type Jrpc struct {
	cli *channelClient
}

// Grpc grpc struct
type Grpc struct {
	*channelClient
}

type channelClient struct {
	types.ChannelClient
}

// Init init grpc param
func Init(name string, s types.RPCServer) {
	cli := &channelClient{}
	grpc := &Grpc{channelClient: cli}
	cli.Init(name, s, &Jrpc{cli: cli}, grpc)
}
