package events

type Event struct {
	StoryId   int64  `json:"story_id" gorm:"primaryKey"`
	AppId     int64  `json:"app_id"`
	EventType string `json:"event_type"`
	UserId    string `json:"user_id"`
	Count     int64  `json:"count"`
	Date      string `json:"date"`
}

func (e *Event) IsMissing() bool {
	return e.StoryId == 0
}
