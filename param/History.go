package param

// HistoryParam
// @Description: 前端发来的参数结构体模型
type HistoryParam struct {
	ID        uint  `json:"id"`
	StartTime int64 `json:"start_time"`
	EndTime   int64 `json:"end_time"`
	Page      int   `json:"page"`
	PageSize  int   `json:"page_size"`
}
