package core

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
)

const (
	KeyNonce = "nonce"
	KeyCSRF  = "csrf"
)

// GenerateNonce returns a base64 encoded url safe nonce of the given size in bytes.
func GenerateNonce(size uint) string

// SetCSRF generates a CSRF token and adds it to the request context and a response cookie.
func SetCSRF(w http.ResponseWriter, r *http.Request)

// CSRF returns a hidden form input with the CSRF token as it's value.
func CSRF(ctx context.Context) templ.Component
