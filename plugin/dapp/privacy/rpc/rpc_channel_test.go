// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rpc_test

import (
	"fmt"
	"strings"
	"testing"

	commonlog "github.com/PhenixChain/devchain/common/log"
	"github.com/PhenixChain/devchain/rpc/jsonclient"
	rpctypes "github.com/PhenixChain/devchain/rpc/types"
	"github.com/PhenixChain/devchain/types"
	"github.com/PhenixChain/devchain/util/testnode"
	pty "github.com/PhenixChain/devchain/plugin/dapp/privacy/types"
	"github.com/stretchr/testify/assert"

	_ "github.com/PhenixChain/devchain/system"
	_ "github.com/PhenixChain/devchain/plugin"
)

func init() {
	commonlog.SetLogLevel("error")
}

func TestRPCChannel(t *testing.T) {
	// 启动RPCmocker
	mocker := testnode.New("--notset--", nil)
	defer func() {
		mocker.Close()
	}()
	mocker.Listen()

	rpcCfg := mocker.GetCfg().RPC
	jrpcClient, err := jsonclient.NewJSONClient(fmt.Sprintf("http://%s/", rpcCfg.JrpcBindAddr))
	assert.NoError(t, err)
	assert.NotNil(t, jrpcClient)

	testCases := []struct {
		fn func(*testing.T, *jsonclient.JSONClient) error
	}{
		{fn: testShowPrivacyKey},
		{fn: testShowPrivacyAccountInfo},
		{fn: testShowPrivacyAccountSpend},
		{fn: testShowAmountsOfUTXO},
		{fn: testShowUTXOs4SpecifiedAmount},
		{fn: testListPrivacyTxs},
		{fn: testRescanUtxosOpt},
		{fn: testEnablePrivacy},
	}
	for index, testCase := range testCases {
		err := testCase.fn(t, jrpcClient)
		if err == nil {
			continue
		}
		assert.NotEqualf(t, err, types.ErrActionNotSupport, "test index %d", index)
		if strings.Contains(err.Error(), "rpc: can't find") {
			assert.FailNowf(t, err.Error(), "test index %d", index)
		}
	}
}

func testShowPrivacyKey(t *testing.T, jrpc *jsonclient.JSONClient) error {
	var res pty.ReplyPrivacyPkPair
	params := types.ReqString{
		Data: "1JSRSwp16NvXiTjYBYK9iUQ9wqp3sCxz2p",
	}
	err := jrpc.Call("privacy.ShowPrivacyKey", params, &res)
	return err
}

func testShowPrivacyAccountInfo(t *testing.T, jrpc *jsonclient.JSONClient) error {
	params := pty.ReqPPrivacyAccount{
		Addr:        "1JSRSwp16NvXiTjYBYK9iUQ9wqp3sCxz2p",
		Token:       types.BTY,
		Displaymode: 1,
	}
	var res pty.ReplyPrivacyAccount
	err := jrpc.Call("privacy.ShowPrivacyAccountInfo", params, &res)
	return err
}

func testShowPrivacyAccountSpend(t *testing.T, jrpc *jsonclient.JSONClient) error {
	params := pty.ReqPrivBal4AddrToken{
		Addr:  "1JSRSwp16NvXiTjYBYK9iUQ9wqp3sCxz2p",
		Token: types.BTY,
	}
	var res pty.UTXOHaveTxHashs
	err := jrpc.Call("privacy.ShowPrivacyAccountSpend", params, &res)
	return err
}

func testShowAmountsOfUTXO(t *testing.T, jrpc *jsonclient.JSONClient) error {
	reqPrivacyToken := pty.ReqPrivacyToken{Token: types.BTY}
	var params rpctypes.Query4Jrpc
	params.Execer = pty.PrivacyX
	params.FuncName = "ShowAmountsOfUTXO"
	params.Payload = types.MustPBToJSON(&reqPrivacyToken)

	var res pty.ReplyPrivacyAmounts
	err := jrpc.Call("Chain33.Query", params, &res)
	return err
}

func testShowUTXOs4SpecifiedAmount(t *testing.T, jrpc *jsonclient.JSONClient) error {
	reqPrivacyToken := pty.ReqPrivacyToken{
		Token:  types.BTY,
		Amount: 123456,
	}
	var params rpctypes.Query4Jrpc
	params.Execer = pty.PrivacyX
	params.FuncName = "ShowUTXOs4SpecifiedAmount"
	params.Payload = types.MustPBToJSON(&reqPrivacyToken)

	var res pty.ReplyUTXOsOfAmount
	err := jrpc.Call("Chain33.Query", params, &res)
	return err
}

func testListPrivacyTxs(t *testing.T, jrpc *jsonclient.JSONClient) error {
	params := pty.ReqPrivacyTransactionList{
		Tokenname:    types.BTY,
		SendRecvFlag: 1,
		Direction:    1,
		Count:        16,
		Address:      "13cS5G1BDN2YfGudsxRxr7X25yu6ZdgxMU",
		Seedtxhash:   []byte("0xa64296792f90f364371e0b66fdac622080ceb7b2537ff9152e189aa9e88e61bd"),
	}
	var res rpctypes.WalletTxDetails
	err := jrpc.Call("privacy.GetPrivacyTxByAddr", params, &res)
	return err
}

func testRescanUtxosOpt(t *testing.T, jrpc *jsonclient.JSONClient) error {
	var params pty.ReqRescanUtxos
	var res pty.RepRescanUtxos
	err := jrpc.Call("privacy.RescanUtxos", params, &res)
	return err
}

func testEnablePrivacy(t *testing.T, jrpc *jsonclient.JSONClient) error {
	var params pty.ReqEnablePrivacy
	var res pty.RepEnablePrivacy
	err := jrpc.Call("privacy.EnablePrivacy", params, &res)
	return err
}
