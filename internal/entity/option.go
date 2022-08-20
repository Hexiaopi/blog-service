package entity

type OneOption struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	State    uint8  `json:"state"`
	Operator string `json:"operator"`
}

type ListOption struct {
	Name     string `json:"name"`
	Sort     string `json:"sort"`
	State    uint8  `json:"state"`
	PageNum  int    `json:"page_num"`
	PageSize int    `json:"page_size"`
}

func (option ListOption) GetPageOffset() int {
	return (option.PageNum - 1) * option.PageSize
}
