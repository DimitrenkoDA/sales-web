package salemans

type Saleman struct {
	ID        uint64  `json:"id"`
	Name      string  `json:"name"`
	Code      uint64  `json:"code"`
	DealerID  uint64  `json:"dealer_id"`
	StatusID  uint64  `json:"status_id"`
	Note      *string `json:"note"`
	Condition string  `json:"condition"`
}

type Top5 struct {
	Name string `json:"name"`
	Cash uint64 `json:"cash"`
	Rank uint64 `json:"rank"`
}

type Unsold struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}
