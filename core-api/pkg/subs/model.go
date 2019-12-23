package subs

type Sub struct {
	ID       uint64  `json:"id"`
	Name     string  `json:"name"`
	Address  string  `json:"address"`
	Phone    string  `json:"phone"`
	StatusID uint64  `json:"status_id"`
	Note     *string `json:"note"`
}
