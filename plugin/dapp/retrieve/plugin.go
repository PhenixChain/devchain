// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package retrieve

import (
	"github.com/PhenixChain/devchain/pluginmgr"
	"github.com/PhenixChain/devchain/plugin/dapp/retrieve/cmd"
	"github.com/PhenixChain/devchain/plugin/dapp/retrieve/executor"
	"github.com/PhenixChain/devchain/plugin/dapp/retrieve/rpc"
	"github.com/PhenixChain/devchain/plugin/dapp/retrieve/types"
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     types.RetrieveX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      cmd.RetrieveCmd,
		RPC:      rpc.Init,
	})
}
