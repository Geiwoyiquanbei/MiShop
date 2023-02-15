package models

type Cart struct {
	Id           int
	Title        string
	Price        float64
	GoodsVersion string
	Uid          int
	Num          int
	GoodsGift    string
	GoodsFitting string
	GoodsColor   string
	GoodsImg     string
	GoodsAttr    string
	Checked      bool
}
