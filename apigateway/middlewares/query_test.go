package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/s3f4/go-load/apigateway/library"
	"github.com/stretchr/testify/assert"
)

func Test_Query(t *testing.T) {
	m := NewMiddleware(nil, nil, nil, nil, nil)

	var next http.HandlerFunc
	next = func(w http.ResponseWriter, r *http.Request) {
		val, ok := r.Context().Value(QueryCtxKey).(*library.QueryBuilder)
		if !ok {
			t.Error("query not found")
		}
		assert.Equal(t, 1, val.Limit)
		assert.Equal(t, 5, val.Offset)
		assert.Equal(t, "abc ASC", val.Order)
	}

	req := httptest.NewRequest(http.MethodGet, "/?limit=1&order=i__abc&offset=5&order=abc", nil)
	res := httptest.NewRecorder()

	test := m.QueryCtx(next)
	test.ServeHTTP(res, req)
}
