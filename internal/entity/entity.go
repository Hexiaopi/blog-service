package entity

type Article struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
	CreateTime    string `json:"create_time"`
	UpdateTime    string `json:"update_time"`
	Operator      string `json:"operator"`
	Tags          []Tag  `json:"tags"`
}

type Tag struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	State      uint8  `json:"state"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
	Operator   string `json:"operator"`
}
