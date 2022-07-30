package utils

import (
	"errors"
	"net/http"
)

type CookieBuilder struct {
	Name     string
	Value    string
	Path     string
	Domain   string
	MaxAge   int
	Secure   *bool
	HttpOnly *bool
}

func (builder *CookieBuilder) Build() (*http.Cookie, error) {
	if builder.Name == "" {
		return nil, errors.New("Cookie name cannot be empty")
	}
	if builder.Value == "" {
		return nil, errors.New("Cookie value cannot be empty")
	}
	var path string = builder.Path

	if builder.Path == "" {
		path = "/"
	}

	var domain string = builder.Domain

	if builder.Domain == "" {
		// this should be set on the basis of some environment variable
		domain = "localhost"
	}

	if builder.MaxAge == 0 {
		builder.MaxAge = 60 * 60 * 24 * 365
	}

	if builder.Secure == nil {
		var trueVal = true
		builder.Secure = &trueVal
	}

	if builder.HttpOnly == nil {
		var trueVal = true
		builder.HttpOnly = &trueVal
	}

	cookie := http.Cookie{
		Path:     path,
		Domain:   domain,
		MaxAge:   builder.MaxAge,
		Secure:   *builder.Secure,
		HttpOnly: *builder.HttpOnly,
		Name:     builder.Name,
		Value:    builder.Value,
		SameSite: http.SameSiteLaxMode,
	}
	return &cookie, nil

}

func NewCookieBuilder() *CookieBuilder {
	return &CookieBuilder{}
}

func (builder *CookieBuilder) SetName(name string) *CookieBuilder {
	builder.Name = name
	return builder
}

func (builder *CookieBuilder) SetValue(value string) *CookieBuilder {
	builder.Value = value
	return builder
}

func (builder *CookieBuilder) SetPath(path string) *CookieBuilder {
	builder.Path = path
	return builder
}

func (builder *CookieBuilder) SetDomain(domain string) *CookieBuilder {
	builder.Domain = domain
	return builder
}

func (builder *CookieBuilder) SetMaxAge(maxAge int) *CookieBuilder {
	builder.MaxAge = maxAge
	return builder
}

func (builder *CookieBuilder) SetHttpOnly(httpOnly bool) *CookieBuilder {
	builder.HttpOnly = &httpOnly
	return builder
}

func (builder *CookieBuilder) SetSecure(secure bool) *CookieBuilder {
	builder.Secure = &secure
	return builder
}
