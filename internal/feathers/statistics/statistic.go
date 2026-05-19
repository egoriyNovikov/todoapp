package statistics

type SummaryStatistics struct {
	UserID string     `json:"user_id"`
	Counts TaskCounts `json:"counts"`
	Rate   float64    `json:"rate"`
}

type TaskCounts struct {
	Completed int `json:"completed"`
	Pending   int `json:"pending"`
	Total     int `json:"total"`
}
