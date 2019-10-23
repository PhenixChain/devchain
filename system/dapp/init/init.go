// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package init 初始化系统dapp包
package init

import (
	_ "github.com/PhenixChain/devchain/system/dapp/coins"  // register coins package
	_ "github.com/PhenixChain/devchain/system/dapp/manage" // register manage package
	_ "github.com/PhenixChain/devchain/system/dapp/none"   // register none package
)
