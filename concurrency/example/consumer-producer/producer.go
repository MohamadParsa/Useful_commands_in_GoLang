package consumerproducer

import "time"

type producer struct {
	data chan order
	quit chan chan error
}
type product struct {
	ID          uint
	qualityRate int
	createTime  time.Time
}

type order struct {
	OrderNo uint
	Count   uint
}

func factory(p *producer) {

}
