package story_responses

type StoryListResponse struct {
	AppId     int64           `json:"app_id"`
	Timestamp int64           `json:"ts"`
	Metadata  []StoryResponse `json:"metadata"`
}
