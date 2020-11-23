package library

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

// QueryBuilder ...
type QueryBuilder struct {
	tx     *gorm.DB
	Limit  int
	Offset int
	Sort   string
	Model  interface{}
}

// SetDB ..
func (q *QueryBuilder) SetDB(db *gorm.DB) *QueryBuilder {
	q.tx = db
	return q
}

// SetModel ...
func (q *QueryBuilder) SetModel(model interface{}) *QueryBuilder {
	q.Model = model
	return q
}

// SetWhere ...
func (q *QueryBuilder) SetWhere(conditionStr string, values ...interface{}) *QueryBuilder {
	q.tx = q.tx.Where(conditionStr, values)
	return q
}

// SetPreloads ...
func (q *QueryBuilder) SetPreloads(preloads ...string) *QueryBuilder {
	if len(preloads) > 0 {
		for _, preload := range preloads {
			q.tx = q.tx.Preload(preload)
		}
	}
	return q
}

// List function
func (q *QueryBuilder) List(out interface{}) error {
	q.tx.Limit(q.Limit).Offset(q.Offset)

	// Check column that will be sorted is exists
	if len(q.Sort) > 0 && IsIn(strings.Split(q.Sort, " ")[0], q.Model) {
		q.tx.Order(q.Sort)
	}

	return q.tx.Find(out).Error
}

func checkParam(param []string) (int, bool) {
	if len(param) > 1 {
		return 0, false
	}

	if val, err := strconv.Atoi(param[0]); err == nil {
		return val, true
	}
	return 0, false
}

// Build a query
func (q *QueryBuilder) Build(query url.Values) {
	fmt.Println(query)
	if l, ok := query["limit"]; ok {
		if limit, ok := checkParam(l); ok {
			q.Limit = limit
		}
	}

	if o, ok := query["offset"]; ok {
		if offset, ok := checkParam(o); ok {
			q.Offset = offset
		}
	}

	if s, ok := query["sort"]; ok {
		sort := s[0]
		if strings.HasPrefix(sort, "-") {
			q.Sort = fmt.Sprintf("%s DESC", strings.Split(sort, "-")[1])
		}

		if strings.HasPrefix(sort, "+") {
			q.Sort = fmt.Sprintf("%s ASC", strings.Split(sort, "+")[1])
		}
	}
}
