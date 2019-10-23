// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rpc_test

import (
	"fmt"
	"testing"

	commonlog "github.com/PhenixChain/devchain/common/log"
	"github.com/PhenixChain/devchain/rpc/jsonclient"
	rpctypes "github.com/PhenixChain/devchain/rpc/types"
	"github.com/PhenixChain/devchain/types"
	"github.com/PhenixChain/devchain/util/testnode"
	pty "github.com/PhenixChain/devchain/plugin/dapp/pokerbull/types"
	"github.com/stretchr/testify/assert"

	_ "github.com/PhenixChain/devchain/system"
	_ "github.com/PhenixChain/devchain/plugin"
)

func init() {
	commonlog.SetLogLevel("error")
}

func TestJRPCChannel(t *testing.T) {
	// 启动RPCmocker
	mocker := testnode.New("--notset--", nil)
	defer func() {
		mocker.Close()
	}()
	mocker.Listen()

	jrpcClient := mocker.GetJSONC()
	assert.NotNil(t, jrpcClient)

	testCases := []struct {
		fn func(*testing.T, *jsonclient.JSONClient) error
	}{
		{fn: testStartRawTxCmd},
		{fn: testContinueRawTxCmd},
		{fn: testQuitRawTxCmd},
	}
	for _, testCase := range testCases {
		err := testCase.fn(t, jrpcClient)
		assert.Nil(t, err)
	}

	testCases = []struct {
		fn func(*testing.T, *jsonclient.JSONClient) error
	}{
		{fn: testQueryGameByID},
		{fn: testQueryGameByAddr},
		{fn: testQueryGameByStatus},
		{fn: testQueryGameByRound},
	}
	for index, testCase := range testCases {
		err := testCase.fn(t, jrpcClient)
		assert.Equal(t, err, types.ErrNotFound, fmt.Sprint(index))
	}

	testCases = []struct {
		fn func(*testing.T, *jsonclient.JSONClient) error
	}{
		{fn: testQueryGameByIDs},
	}
	for index, testCase := range testCases {
		err := testCase.fn(t, jrpcClient)
		assert.Equal(t, err, nil, fmt.Sprint(index))
	}
}

func testStartRawTxCmd(t *testing.T, jrpc *jsonclient.JSONClient) error {
	payload := &pty.PBGameStart{Value: 123}
	params := &rpctypes.CreateTxIn{
		Execer:     types.ExecName(pty.PokerBullX),
		ActionName: pty.CreateStartTx,
		Payload:    types.MustPBToJSON(payload),
	}
	var res string
	return jrpc.Call("Chain33.CreateTransaction", params, &res)
}

func testContinueRawTxCmd(t *testing.T, jrpc *jsonclient.JSONClient) error {
	payload := &pty.PBGameContinue{GameId: "123"}
	params := &rpctypes.CreateTxIn{
		Execer:     types.ExecName(pty.PokerBullX),
		ActionName: pty.CreateContinueTx,
		Payload:    types.MustPBToJSON(payload),
	}
	var res string
	return jrpc.Call("Chain33.CreateTransaction", params, &res)
}

func testQuitRawTxCmd(t *testing.T, jrpc *jsonclient.JSONClient) error {
	payload := &pty.PBGameQuit{GameId: "123"}
	params := &rpctypes.CreateTxIn{
		Execer:     types.ExecName(pty.PokerBullX),
		ActionName: pty.CreateQuitTx,
		Payload:    types.MustPBToJSON(payload),
	}
	var res string
	return jrpc.Call("Chain33.CreateTransaction", params, &res)
}

func testQueryGameByID(t *testing.T, jrpc *jsonclient.JSONClient) error {
	var rep interface{}
	var params rpctypes.Query4Jrpc
	req := &pty.QueryPBGameInfo{}
	params.Execer = pty.PokerBullX
	params.FuncName = pty.FuncNameQueryGameByID
	params.Payload = types.MustPBToJSON(req)
	rep = &pty.ReplyPBGame{}
	return jrpc.Call("Chain33.Query", params, rep)
}

func testQueryGameByAddr(t *testing.T, jrpc *jsonclient.JSONClient) error {
	var rep interface{}
	var params rpctypes.Query4Jrpc
	req := &pty.QueryPBGameInfo{}
	params.Execer = pty.PokerBullX
	params.FuncName = pty.FuncNameQueryGameByAddr
	params.Payload = types.MustPBToJSON(req)
	rep = &pty.PBGameRecords{}
	return jrpc.Call("Chain33.Query", params, rep)
}

func testQueryGameByIDs(t *testing.T, jrpc *jsonclient.JSONClient) error {
	var rep interface{}
	var params rpctypes.Query4Jrpc
	req := &pty.QueryPBGameInfos{}
	params.Execer = pty.PokerBullX
	params.FuncName = pty.FuncNameQueryGameListByIDs
	params.Payload = types.MustPBToJSON(req)
	rep = &pty.ReplyPBGameList{}
	return jrpc.Call("Chain33.Query", params, rep)
}

func testQueryGameByStatus(t *testing.T, jrpc *jsonclient.JSONClient) error {
	var rep interface{}
	var params rpctypes.Query4Jrpc
	req := &pty.QueryPBGameInfo{}
	params.Execer = pty.PokerBullX
	params.FuncName = pty.FuncNameQueryGameByStatus
	params.Payload = types.MustPBToJSON(req)
	rep = &pty.PBGameRecords{}
	return jrpc.Call("Chain33.Query", params, rep)
}

func testQueryGameByRound(t *testing.T, jrpc *jsonclient.JSONClient) error {
	var rep interface{}
	var params rpctypes.Query4Jrpc
	req := &pty.QueryPBGameByRound{}
	params.Execer = pty.PokerBullX
	params.FuncName = pty.FuncNameQueryGameByRound
	params.Payload = types.MustPBToJSON(req)
	rep = &pty.PBGameRecords{}
	return jrpc.Call("Chain33.Query", params, rep)
}