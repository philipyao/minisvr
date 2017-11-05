package main

//shop简要信息
type ShopProfile struct {
	ID          uint32          `json:"id"`
	Name        string          `json:"name"`
	Location    string          `json:"location"`
	Addr        string          `json:"addr"`
}


type AdminListShopRsp struct {
	Shops       []*ShopProfile  `json:"shops"`
}
