// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package executor

import (
	log "github.com/PhenixChain/devchain/common/log/log15"
	drivers "github.com/PhenixChain/devchain/system/dapp"
	"github.com/PhenixChain/devchain/types"
	dty "github.com/PhenixChain/devchain/plugin/dapp/dposvote/types"
)

var logger = log.New("module", "execs.dposvote")

var driverName = dty.DPosX

var (
	dposDelegateNum          int64 = 3 //委托节点个数，从配置读取，以后可以根据投票结果来定
	dposBlockInterval        int64 = 3 //出块间隔，当前按3s
	dposContinueBlockNum     int64 = 6 //一个委托节点当选后，一次性持续出块数量
	dposCycle                      = dposDelegateNum * dposBlockInterval * dposContinueBlockNum
	dposPeriod                     = dposBlockInterval * dposContinueBlockNum
	blockNumToUpdateDelegate int64 = 20000
	registTopNHeightLimit    int64 = 100
	updateTopNHeightLimit    int64 = 200
)

// CycleInfo indicates the start and stop of a cycle
type CycleInfo struct {
	cycle      int64
	cycleStart int64
	cycleStop  int64
}

func calcCycleByTime(now int64) *CycleInfo {
	cycle := now / dposCycle
	cycleStart := now - now%dposCycle
	cycleStop := cycleStart + dposCycle - 1

	return &CycleInfo{
		cycle:      cycle,
		cycleStart: cycleStart,
		cycleStop:  cycleStop,
	}
}

func calcTopNVersion(height int64) (version, left int64) {
	return height / blockNumToUpdateDelegate, height % blockNumToUpdateDelegate
}

func init() {
	ety := types.LoadExecutorType(driverName)
	ety.InitFuncList(types.ListMethod(&DPos{}))
}

// Init DPos Executor
func Init(name string, sub []byte) {
	driverName := GetName()
	if name != driverName {
		panic("system dapp can't be rename")
	}

	drivers.Register(driverName, newDposVote, types.GetDappFork(driverName, "Enable"))

	//读取一下配置项，用于和共识模块一致计算cycle
	dposDelegateNum = types.Conf("config.consensus.sub.dpos").GInt("delegateNum")
	dposBlockInterval = types.Conf("config.consensus.sub.dpos").GInt("blockInterval")
	dposContinueBlockNum = types.Conf("config.consensus.sub.dpos").GInt("continueBlockNum")
	blockNumToUpdateDelegate = types.Conf("config.consensus.sub.dpos").GInt("blockNumToUpdateDelegate")
	registTopNHeightLimit = types.Conf("config.consensus.sub.dpos").GInt("registTopNHeightLimit")
	updateTopNHeightLimit = types.Conf("config.consensus.sub.dpos").GInt("updateTopNHeightLimit")
	dposCycle = dposDelegateNum * dposBlockInterval * dposContinueBlockNum
	dposPeriod = dposBlockInterval * dposContinueBlockNum
}

//DPos 执行器，用于Dpos候选节点注册、投票，VRF信息注册管理等功能
type DPos struct {
	drivers.DriverBase
}

func newDposVote() drivers.Driver {
	t := &DPos{}
	t.SetChild(t)
	t.SetExecutorType(types.LoadExecutorType(driverName))
	return t
}

//GetName 获取DPos执行器的名称
func GetName() string {
	return newDposVote().GetName()
}

//ExecutorOrder Exec 的时候 同时执行 ExecLocal
func (g *DPos) ExecutorOrder() int64 {
	return drivers.ExecLocalSameTime
}

//GetDriverName 获取DPos执行器的名称
func (g *DPos) GetDriverName() string {
	return dty.DPosX
}

// CheckReceiptExecOk return true to check if receipt ty is ok
func (g *DPos) CheckReceiptExecOk() bool {
	return true
}
