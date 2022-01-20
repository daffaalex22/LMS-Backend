package request

import "backend/business/attachments"

type AttachmentsAdd struct {
	ContentType string `json:"contentType"`
	ContentId   uint   `json:"contentId"`
	VideoId     uint   `json:"videoId"`
	ReadingId   uint   `json:"readingId"`
	Title       string `json:"title"`
	Url         string `json:"url"`
}

type AttachmentsUpdate struct {
	ContentType string `json:"contentType"`
	ContentId   uint   `json:"contentId"`
	VideoId     uint   `json:"videoId"`
	ReadingId   uint   `json:"readingId"`
	Title       string `json:"title"`
	Url         string `json:"url"`
}

func (att *AttachmentsAdd) ToDomain() attachments.Domain {
	return attachments.Domain{
		ContentType: att.ContentType,
		ContentId:   att.ContentId,
		VideoId:     att.VideoId,
		ReadingId:   att.ReadingId,
		Title:       att.Title,
		Url:         att.Url,
	}
}
func (att *AttachmentsUpdate) ToDomain() attachments.Domain {
	return attachments.Domain{
		ContentType: att.ContentType,
		ContentId:   att.ContentId,
		VideoId:     att.VideoId,
		ReadingId:   att.ReadingId,
		Title:       att.Title,
		Url:         att.Url,
	}
}
