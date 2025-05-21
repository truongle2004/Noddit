package utils

type Pagination struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}

func NewPagination(page, limit int) Pagination {
	return Pagination{
		Page:  page,
		Limit: limit,
	}
}

func (p *Pagination) GetOffset() int {
	if p.Page < 1 {
		p.Page = 1
	}
	if p.Limit <= 0 {
		p.Limit = 10
	}
	return (p.Page - 1) * p.Limit
}
