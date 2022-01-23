package request

import readings "backend/business/readings"

type ReadingsAdd struct {
	ModuleId   uint   `json:"moduleId"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Quiz       string `json:"quiz"`
	Attachment string `json:"attachment"`
	Order      int    `json:"order"`
}
type ReadingsUpdate struct {
	ModuleId   uint   `json:"moduleId"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Quiz       string `json:"quiz"`
	Attachment string `json:"attachment"`
	Order      int    `json:"order"`
}

func (rds *ReadingsAdd) ToDomain() readings.Domain {
	return readings.Domain{
		ModuleId:   rds.ModuleId,
		Title:      rds.Title,
		Content:    rds.Content,
		Quiz:       rds.Quiz,
		Attachment: rds.Attachment,
		Order:      rds.Order,
	}
}
func (rds *ReadingsUpdate) ToDomain() readings.Domain {
	return readings.Domain{
		ModuleId:   rds.ModuleId,
		Title:      rds.Title,
		Quiz:       rds.Quiz,
		Attachment: rds.Attachment,
		Content:    rds.Content,
		Order:      rds.Order,
	}
}
