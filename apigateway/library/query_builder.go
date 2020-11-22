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
	Limit  int
	Offset int
	Sort   string
	Model  interface{}
}

// SetModel ...
func (q *QueryBuilder) SetModel(model interface{}) *QueryBuilder {
	q.Model = model
	return q
}

// List function
func (q *QueryBuilder) List(db *gorm.DB, out interface{}, preloads ...string) (*gorm.DB, error) {
	db = db.Limit(q.Limit).Offset(q.Offset)

	// Check column that will be sorted is exists
	if len(q.Sort) > 0 && IsIn(strings.Split(q.Sort, " ")[0], q.Model) {
		db = db.Order(q.Sort)
	}

	if len(preloads) > 0 {
		for _, preload := range preloads {
			db = db.Preload(preload)
		}
	}

	db = db.Find(out)

	return db, db.Error
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
