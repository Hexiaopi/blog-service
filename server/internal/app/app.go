package app

import (
	"encoding/json"
	"net/http"

	"github.com/hexiaopi/blog-service/internal/retcode"
)

type CommResponse struct {
	*retcode.Error
	Data interface{} `json:"data,omitempty"`
}

type ErrResponse struct {
	*retcode.Error
}

type ListResponse struct {
	*retcode.Error
	TotalRows int64       `json:"total_rows"`
	Data      interface{} `json:"data,omitempty"`
}

func ToResponseCode(writer http.ResponseWriter, code *retcode.Error) {
	response := ErrResponse{
		code,
	}
	result, _ := json.Marshal(response)
	writer.Write(result)
}

func ToResponseData(writer http.ResponseWriter, data interface{}) {
	response := CommResponse{
		Error: retcode.Success,
		Data:  data,
	}
	result, _ := json.Marshal(response)
	writer.Write(result)
}

func ToResponseList(writer http.ResponseWriter, total int64, data interface{}) {
	response := ListResponse{
		Error:     retcode.Success,
		TotalRows: total,
		Data:      data,
	}
	result, _ := json.Marshal(response)
	writer.Write(result)
}
