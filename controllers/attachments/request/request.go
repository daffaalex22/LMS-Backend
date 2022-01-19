package request

import Attacchments "backend/business/Atacchments"

type AttacchmentsAdd stuct {
	ModuleId uint   `json:"moduleId"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	Url      string `json:"url"`
	Order    int    `json:"order"`
}
type AttacchmentsUpdate stuct {
	ModuleId uint   `json:"moduleId"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	Url      string `json:"url"`
	Order    int    `json:"order"`
}

func (rds *AttacchmentsAdd) ToDomain() attachmens.Domain {
	return attachments.Doain{
		ModuleId: rds.ModuleId,
		Title:    rds.Title,
		Caption:  rds.Caption,
		Url:      rds.Url,
		Order:    rds.Order,
	}
}
func (rds *AttacchmentsUpdate) ToDomain() attachmens.Domain {
	return attachments.Doain{
		ModuleId: rds.ModuleId,
		Title:    rds.Title,
		Caption:  rds.Caption,
		Url:      rds.Url,
		Order:    rds.Order,
	}
}
