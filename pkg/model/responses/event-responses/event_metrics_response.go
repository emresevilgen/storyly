package event_responses

type EventMetricsResponse struct {
	AppId            int64 `json:"app_id"`
	DailyActiveUsers int64 `json:"dailyActiveUsers"`
}
