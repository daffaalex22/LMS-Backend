package request

type AddRequest struct {
	Id           uint   `json:"id"`
	Title        string `json:"title"`
	Thumbnail    string `json:"thumbnail"`
	Description  string `json:"description"`
	CategoryId   uint   `json:"category_id"`
	TeacherId    uint   `json:"teacher_id"`
	DifficultyId uint   `json:"difficulty_id"`
}
