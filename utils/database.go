package utils

import (
	"math"
	"strconv"
	"errors"

	"gorm.io/gorm"
)

type Pagination struct {
	CurrentPage 	int
	TotatPages		int
	PerPage 		int
	Offset 			int
	HasNext			bool
	HasPrevious		bool
	Next 			int
	Previous 		int
	Error			error
}

func NewPagination(totalItems int64, perPage int, currentPage string) *Pagination {
	var paginationError error
	var page int
	var err error

	division := float64(totalItems) / float64(perPage)
	totalPages := int(math.Ceil(division))

	if currentPage == "" {
		page = 1
	} else {
		page, err = strconv.Atoi(currentPage)

		if err != nil {
			paginationError = errors.New("Can't parse page!")
			page = 1
		}

		if page < 1 {
			paginationError = errors.New("Page is less than 1!")
			page = 1
		} else if page > totalPages {
			paginationError = errors.New("Page is greater than total pages!")
			page = 1
		}
	}

	offset := (page - 1) * perPage

	hasNext := page < totalPages
	hasPervious	:= page > 1

	return &Pagination{
		CurrentPage: page,
		TotatPages: totalPages,
		PerPage: perPage,
		Offset: offset,
		HasNext: hasNext,
		HasPrevious: hasPervious,
		Next: (page + 1),
		Previous: (page - 1),
		Error: paginationError,
	}

}

func (p *Pagination) Paginate()  func(db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
		return db.Offset(p.Offset).Limit(p.PerPage)
	}
}