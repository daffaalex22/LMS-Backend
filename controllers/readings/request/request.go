package request

import readings "backend/business/readings"

type ReadingsAdd struct {
	ModuleId uint   `json:"moduleId"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Order    int    `json:"order"`
}
type ReadingsUpdate struct {
	ModuleId uint   `json:"moduleId"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Order    int    `json:"order"`
}

func (rds *ReadingsAdd) ToDomain() readings.Domain {
	return readings.Domain{
		ModuleId: rds.ModuleId,
		Title:    rds.Title,
		Content:  rds.Content,
		Order:    rds.Order,
	}
}
func (rds *ReadingsUpdate) ToDomain() readings.Domain {
	return readings.Domain{
		ModuleId: rds.ModuleId,
		Title:    rds.Title,
		Content:  rds.Content,
		Order:    rds.Order,
	}
}
