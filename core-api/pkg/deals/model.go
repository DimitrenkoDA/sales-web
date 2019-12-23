package deals

import "time"

type Deal struct {
	ID         uint64     `json:"id"`
	StartedAt  time.Time  `json:"started_at"`
	FinishedAt *time.Time `json:"finished_at"`
	Note       *string    `json:"note"`
	SubId      uint64     `json:"sub_id"`
}
