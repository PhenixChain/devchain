package price

import (
	"github.com/PhenixChain/devchain/common/skiplist"
	"github.com/PhenixChain/devchain/system/mempool"
	"github.com/golang/protobuf/proto"
)

// Queue 价格队列模式(价格=手续费/交易字节数,价格高者优先,同价则时间早优先)
type Queue struct {
	*skiplist.Queue
	subConfig subConfig
}

type priceScore struct {
	*mempool.Item
}

func (item *priceScore) GetScore() int64 {
	txSize := proto.Size(item.Value)
	return item.Value.Fee / int64(txSize)
}

func (item *priceScore) Hash() []byte {
	return item.Value.Hash()
}

func (item *priceScore) Compare(cmp skiplist.Scorer) int {
	it := cmp.(*priceScore)
	//时间越小，权重越高
	if item.EnterTime < it.EnterTime {
		return skiplist.Big
	}
	if item.EnterTime == it.EnterTime {
		return skiplist.Equal
	}
	return skiplist.Small
}

// NewQueue 创建队列
func NewQueue(subcfg subConfig) *Queue {
	return &Queue{
		Queue:     skiplist.NewQueue(subcfg.PoolCacheSize),
		subConfig: subcfg,
	}
}

//GetItem 获取数据通过 key
func (cache *Queue) GetItem(hash string) (*mempool.Item, error) {
	item, err := cache.Queue.GetItem(hash)
	if err != nil {
		return nil, err
	}
	return item.(*priceScore).Item, nil
}

//Push 加入数据到队列
func (cache *Queue) Push(item *mempool.Item) error {
	return cache.Queue.Push(&priceScore{Item: item})
}

//Walk 获取数据通过 key
func (cache *Queue) Walk(count int, cb func(tx *mempool.Item) bool) {
	cache.Queue.Walk(count, func(item skiplist.Scorer) bool {
		return cb(item.(*priceScore).Item)
	})
}

// GetProperFee 获取合适的手续费率,取前100的平均手续费率
func (cache *Queue) GetProperFee() int64 {
	var sumFeeRate int64
	var properFeeRate int64
	if cache.Size() == 0 {
		return cache.subConfig.ProperFee
	}
	i := 0
	var txSize int
	var feeRate int64
	cache.Walk(100, func(item *mempool.Item) bool {
		txSize = proto.Size(item.Value)
		feeRate = item.Value.Fee / int64(txSize/1000+1)
		sumFeeRate += feeRate
		i++
		return true
	})
	properFeeRate = sumFeeRate / int64(i)
	return properFeeRate
}
