package story_documents

type StoryDocument struct {
	AppId    int64            `json:"app_id"`
	Id       int64            `json:"story_id"`
	Metadata MetadataDocument `json:"metadata"`
}

type MetadataDocument struct {
	Image string `json:"img"`
}
