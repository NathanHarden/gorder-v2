package order

import "github.com/NathanHarden/gorder-v2/common/genproto/orderpb"

type Order struct{
	ID string
	CustomerID string
	Status string
	Items []orderpb.Item
	PaymentLink string
}