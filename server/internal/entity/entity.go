package entity

type Article struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	State         uint8  `json:"state"`
	Tags          []Tag  `json:"tags"`
}

type Tag struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      uint8  `json:"state"`
}
