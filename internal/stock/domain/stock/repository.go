package domain

import (
	"context"
	"fmt"
	"strings"

	"github.com/NathanHarden/gorder-v2/common/genproto/orderpb"
)

type Repository interface{
	GetItems(ctx context.Context,ids []string)([]*orderpb.Item,error)
}

type NotFoundError struct{
	IDs []string
}

func (e NotFoundError) Error() string{
	return fmt.Sprintf("these items are not found in stock:%s",strings.Join(e.IDs,","))
}