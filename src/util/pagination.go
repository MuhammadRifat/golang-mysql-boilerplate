package util

import (
	"math"
	"strconv"
)

type Pagination struct {
	TotalRecord int64
	TotalPage   int64
	Limit       int64
	Page        int64
}

type Paginate struct {
	Limit  int   `form:"limit"`
	Count  int64 `form:"count"`
	Page   int   `form:"page"`
	Offset int
}

func PaginateDefault(RequestPage string, RequestLimit string) Paginate {
	limit, err := strconv.Atoi(RequestLimit)
	if err != nil {
		limit = 10
	}
	page, err := strconv.Atoi(RequestPage)
	if err != nil || page <= 0 {
		page = 1
	}
	offset := (page - 1) * limit

	return Paginate{
		Limit:  limit,
		Page:   page,
		Offset: offset,
	}
}

func PaginationMake(paginate Paginate) Pagination {
	// pagination
	pgn := Pagination{}
	pgn.TotalRecord = paginate.Count
	pgn.TotalPage = int64(math.Ceil((float64(paginate.Count) / float64(paginate.Limit))))
	pgn.Limit = int64(paginate.Limit)
	pgn.Page = int64(paginate.Page)

	return pgn
}
