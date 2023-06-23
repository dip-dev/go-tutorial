package chapter1

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEcho(t *testing.T) {
	success := map[string]struct {
		params     map[string]string
		wantStatus int
	}{
		"正常: パラメータあり": {
			params: map[string]string{
				"name": "dip 太郎",
				"age":  "26",
			},
			wantStatus: http.StatusOK,
		},
		"正常: パラメータなし": {
			params:     map[string]string{},
			wantStatus: http.StatusOK,
		},
	}

	fail := map[string]struct {
		method     string
		params     map[string]string
		wantStatus int
	}{
		"異常: Getメソッドではない": {
			method:     http.MethodPost,
			params:     map[string]string{},
			wantStatus: http.StatusMethodNotAllowed,
		},
	}

	for tn, tc := range success {
		t.Run(tn, func(t *testing.T) {
			param := url.Values{}
			for k, v := range tc.params {
				param.Add(k, v)
			}

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "http://localhost/?"+param.Encode(), nil)
			GetEcho(w, r)
			got := w.Body.String()
			assert.Equal(t, tc.wantStatus, w.Code)
			for k, v := range tc.params {
				assert.Contains(t, got, k)
				assert.Contains(t, got, v)
			}
		})
	}
	for tn, tc := range fail {
		t.Run(tn, func(t *testing.T) {
			form := url.Values{}
			for k, v := range tc.params {
				form.Add(k, v)
			}
			w := httptest.NewRecorder()
			r := httptest.NewRequest(tc.method, "http://localhost/", strings.NewReader(form.Encode()))
			GetEcho(w, r)
			assert.Equal(t, tc.wantStatus, w.Code)
		})
	}
}
