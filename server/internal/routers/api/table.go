package api

import (
	"encoding/json"
	"net/http"
	"time"
)

type TableResponse struct {
	RetCode string    `json:"ret_code"`
	RetDesc string    `json:"ret_desc"`
	Data    TableData `json:"data"`
}

type TableData struct {
	Items []Item `json:"items"`
}

type Item struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Status      string    `json:"status"`
	Author      string    `json:"author"`
	PageViews   int       `json:"pageviews"`
	DisplayTime time.Time `json:"display_time"`
}

func TableList(writer http.ResponseWriter, request *http.Request) {
	item1 := Item{
		Id:          "1",
		Title:       "11",
		Status:      "published",
		Author:      "11@11",
		PageViews:   10,
		DisplayTime: time.Now().Add(time.Hour * 24),
	}
	item2 := Item{
		Id:          "2",
		Title:       "22",
		Status:      "draft",
		Author:      "22@22",
		PageViews:   40,
		DisplayTime: time.Now(),
	}
	res := TableResponse{
		RetCode: "000000",
		RetDesc: "Success",
		Data:    TableData{Items: make([]Item, 0)},
	}
	res.Data.Items = append(res.Data.Items, item1, item2)
	data, _ := json.Marshal(res)
	writer.Write(data)
}
