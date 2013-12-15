package client

import (
	"errors"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

type ClientBuilder struct {
	jar      http.CookieJar
	redirect func(req *http.Request, via []*http.Request) error
	err      error
}

// New is the entry point method that allows us to begin creating a new http.Client with
// the properties that we'd like
func New() *ClientBuilder {
	return &ClientBuilder{
		jar:      nil,
		redirect: nil,
		err:      nil,
	}
}

// instantiated client will have a default CookieJar attached.  options is an optional
// parameter that defines a pre-loaded set of cookies
func (self *ClientBuilder) WithCookieJar(options *cookiejar.Options) *ClientBuilder {
	jar, err := cookiejar.New(options)
	if err == nil {
		err = self.err
	}

	return &ClientBuilder{
		jar: jar,
		err: err,
	}
}

// allows for a custom CookieJar to be used when constructing the http.Client
func (self *ClientBuilder) WithCookieJarHandler(jar http.CookieJar) *ClientBuilder {
	return &ClientBuilder{
		jar:      jar,
		redirect: self.redirect,
		err:      self.err,
	}
}

// instruct the client to not follow http redirects.  primarily useful for testing
func (self *ClientBuilder) WithNoFollowRedirects() *ClientBuilder {
	redirect := func(req *http.Request, via []*http.Request) error {
		return error(&url.Error{
			Op:  req.Method,
			URL: req.URL.String(),
			Err: errors.New("not following redirects"),
		})
	}

	return &ClientBuilder{
		jar:      self.jar,
		redirect: redirect,
		err:      self.err,
	}
}

// factory method to take all the options we've received so far an turn them into
// an http.Client instance
func (self *ClientBuilder) Build() (*http.Client, error) {
	return &http.Client{
			Jar:           self.jar,
			CheckRedirect: self.redirect,
		},
		self.err
}
