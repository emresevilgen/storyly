package stories

type Story struct {
	StoryId  int64  `json:"story_id"  gorm:"primaryKey"`
	AppId    int64  `json:"app_id"`
	Metadata string `json:"metadata"`
}
