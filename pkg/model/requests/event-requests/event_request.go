package event_requests

import "storyly/pkg/model/enums"

type EventRequest struct {
	StoryId   int64           `json:"story_id" validate:"required"`
	EventType enums.EventType `json:"event_type" validate:"oneof=impression close,required"`
	UserId    string          `json:"user_id" validate:"required"`
}
