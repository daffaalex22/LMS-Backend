package request

import videos "backend/business/videos"

type VideosAdd struct {
	ModuleId uint   `json:"moduleId"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	Url      string `json:"url"`
	Order    int    `json:"order"`
}
type VideosUpdate struct {
	ModuleId uint   `json:"moduleId"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	Url      string `json:"url"`
	Order    int    `json:"order"`
}

func (rds *VideosAdd) ToDomain() videos.Domain {
	return videos.Domain{
		ModuleId: rds.ModuleId,
		Title:    rds.Title,
		Caption:  rds.Caption,
		Url:      rds.Url,
		Order:    rds.Order,
	}
}
func (rds *VideosUpdate) ToDomain() videos.Domain {
	return videos.Domain{
		ModuleId: rds.ModuleId,
		Title:    rds.Title,
		Caption:  rds.Caption,
		Url:      rds.Url,
		Order:    rds.Order,
	}
}
