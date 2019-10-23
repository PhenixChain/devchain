package executor

import (
	"github.com/PhenixChain/devchain/types"
	"github.com/PhenixChain/devchain/plugin/dapp/js/types/jsproto"
)

func (c *js) ExecDelLocal_Create(payload *jsproto.Create, tx *types.Transaction, receiptData *types.ReceiptData, index int) (*types.LocalDBSet, error) {
	return &types.LocalDBSet{}, nil
}

func (c *js) ExecDelLocal_Call(payload *jsproto.Call, tx *types.Transaction, receiptData *types.ReceiptData, index int) (*types.LocalDBSet, error) {
	execer := c.userExecName(payload.Name, true)
	r := &types.LocalDBSet{}
	c.prefix = types.CalcLocalPrefix([]byte(execer))
	kvs, err := c.DelRollbackKV(tx, []byte(execer))
	if err != nil {
		return nil, err
	}
	r.KV = kvs
	return r, nil
}
