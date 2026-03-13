package core

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

const (
	KeyNonce = "nonce"
	KeyCSRF  = "csrf"

	CSRFEntropy = 16 // 128 bits -> 16 bytes
)

// GenerateNonce returns a base64 encoded url safe nonce of the given size in bytes.
func GenerateNonce(size uint) string {
	bytes := make([]byte, size)
	_, _ = rand.Read(bytes)

	return base64.RawURLEncoding.EncodeToString(bytes)
}

// WithCSRF generates a CSRF token and adds it to the request context and a response cookie.
func WithCSRF(w http.ResponseWriter, r *http.Request) *http.Request {
	csrf := GenerateNonce(CSRFEntropy)

	http.SetCookie(w, &http.Cookie{
		Name:     KeyCSRF,
		Value:    csrf,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

	ctx := context.WithValue(r.Context(), KeyCSRF, csrf)
	return r.WithContext(ctx)
}

// CSRF returns a hidden form input with the CSRF token as it's value.
func CSRF(ctx context.Context) templ.Component {
	csrf := ctx.Value(KeyCSRF).(string)
	return templ.Raw(fmt.Sprintf(`<input type="hidden" value="%s"`, csrf))
}
