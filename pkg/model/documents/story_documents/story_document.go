package story_documents

type StoryDocument struct {
	AppId    int64
	Id       int64
	Metadata MetadataDocument
}

type MetadataDocument struct {
	Image string
}
