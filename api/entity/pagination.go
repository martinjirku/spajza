package entity

import (
	"net/url"
	"strconv"
)

type Pagination struct {
	Index int64 `json:"index"`
	Size  int64 `json:"size"`
}

func (p Pagination) Validate() bool {
	if p.Index < 0 || p.Size < 0 {
		return false
	}
	if p.Size > 100 {
		return false
	}
	return true
}

func PaginationFromQuery(query url.Values) Pagination {
	pagination := Pagination{
		Index: 0,
		Size:  10,
	}
	if index, ok := query["index"]; ok {
		parsedIndex, err := strconv.ParseInt(index[0], 10, 64)
		if err == nil {
			pagination.Index = parsedIndex
		}
	}
	if size, ok := query["size"]; ok {
		parsedSize, err := strconv.ParseInt(size[0], 10, 64)
		if err == nil {
			pagination.Size = parsedSize
		}
	}
	return pagination
}
