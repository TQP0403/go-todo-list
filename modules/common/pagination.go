package common

type Pagination struct {
	Page  int   `json:"page" form:"page"`
	Limit int   `json:"limit" form:"limit"`
	Total int64 `json:"total" form:"-"`
}

func (p *Pagination) Validate() {
	if p.Page < 1 {
		p.Page = 1
	}

	if p.Limit < 1 || p.Limit > 100 {
		p.Limit = 10
	}
}
