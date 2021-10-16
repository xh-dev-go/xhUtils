package httpUtils

import (
	"net/http"
)
func NewClient() *http.Client {
	return http.DefaultClient
}