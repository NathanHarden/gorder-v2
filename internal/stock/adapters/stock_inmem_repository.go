package adapters

import (
	"context"
	"sync"

	"github.com/NathanHarden/gorder-v2/common/genproto/orderpb"
	domain "github.com/NathanHarden/gorder-v2/stock/domain/stock"
)

type MemoryStockRepository struct{
	lock *sync.RWMutex
	store map[string]*orderpb.Item
}

var stub=map[string]*orderpb.Item{
	"item_id":{
		ID:"foo_item",
		Name:"stub item",
		Quantity: 10000,
		PriceID: "stub_price_id",
	},
}

func NewMemoryStockRepository() *MemoryStockRepository{
	return &MemoryStockRepository{
		lock:&sync.RWMutex{},
		store:stub,
	}
}

func (m MemoryStockRepository)GetItems(ctx context.Context,ids []string)([]*orderpb.Item,error){
	 m.lock.RLock()
	 defer m.lock.RUnlock()
	 res:=make([]*orderpb.Item,0)
	 var mis domain.NotFoundError
	 for _,id:=range ids{
		if item,ok:=m.store[id];ok{
			res=append(res,item)
		}else{
			mis.IDs=append(mis.IDs, id)
		}
	 }
	 if len(ids)==len(res){
		return res,nil
	 }
	 return res,mis
}

