// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package trade

import (
	"github.com/PhenixChain/devchain/pluginmgr"
	_ "github.com/PhenixChain/devchain/plugin/dapp/trade/autotest" // register autotest package
	"github.com/PhenixChain/devchain/plugin/dapp/trade/commands"
	"github.com/PhenixChain/devchain/plugin/dapp/trade/executor"
	"github.com/PhenixChain/devchain/plugin/dapp/trade/rpc"
	"github.com/PhenixChain/devchain/plugin/dapp/trade/types"
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     types.TradeX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      commands.TradeCmd,
		RPC:      rpc.Init,
	})
}
