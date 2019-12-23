package statuses

type Status struct {
	ID          uint64  `json:"status_id"`
	StutusGroup *uint64 `json:"status_group"`
	Name        string  `json:"status_name"`
	Note        *string `json:"note"`
}
