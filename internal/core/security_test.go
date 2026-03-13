package core_test

import (
	"docpile/internal/core"
	"net/http/httptest"
	"testing"
)

func TestSetCSRF(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)

	r = core.WithCSRF(w, r)

	var csrf string
	for _, cookie := range w.Result().Cookies() {
		if cookie.Name == core.KeyCSRF {
			csrf = cookie.Value
		}
	}

	if csrf == "" {
		t.Errorf("failed to set csrf cookie")
	}

	if ctxCsrf := r.Context().Value(core.KeyCSRF).(string); ctxCsrf != csrf {
		t.Errorf("cookie csrf does not match context, cookie: %s, context:%s", csrf, ctxCsrf)
	}
}
