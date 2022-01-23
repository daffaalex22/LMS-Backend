package request

import videos "backend/business/videos"

type VideosAdd struct {
	ModuleId   uint   `json:"moduleId"`
	Title      string `json:"title"`
	Caption    string `json:"caption"`
	Url        string `json:"url"`
	Order      int    `json:"order"`
	Quiz       string `json:"quiz"`
	Attachment string `json:"attachment"`
}
type VideosUpdate struct {
	ModuleId   uint   `json:"moduleId"`
	Title      string `json:"title"`
	Caption    string `json:"caption"`
	Url        string `json:"url"`
	Order      int    `json:"order"`
	Quiz       string `json:"quiz"`
	Attachment string `json:"attachment"`
}

func (vds *VideosAdd) ToDomain() videos.Domain {
	return videos.Domain{
		ModuleId:   vds.ModuleId,
		Title:      vds.Title,
		Caption:    vds.Caption,
		Url:        vds.Url,
		Quiz:       vds.Quiz,
		Attachment: vds.Attachment,
		Order:      vds.Order,
	}
}
func (vds *VideosUpdate) ToDomain() videos.Domain {
	return videos.Domain{
		ModuleId:   vds.ModuleId,
		Title:      vds.Title,
		Caption:    vds.Caption,
		Url:        vds.Url,
		Quiz:       vds.Quiz,
		Attachment: vds.Attachment,
		Order:      vds.Order,
	}
}
