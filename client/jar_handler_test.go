package client

import (
	"net/http"
	"net/url"
)

type MockJarHandler struct {
	Id string
}

func (self MockJarHandler) SetCookies(u *url.URL, cookies []*http.Cookie) {
}

func (self MockJarHandler) Cookies(u *url.URL) []*http.Cookie {
	return nil
}
