package db

import (
	"fmt"
	"net/url"
)

// Config - represents db configurations.
type Config struct {
	User         string
	Password     string
	Name         string
	Host         string
	Port         string
	MaxIdleConns int
	MaxOpenConns int
}

func (c Config) dbString() string {
	q := make(url.Values)
	q.Set("sslmode", "disable")
	q.Set("timezone", "utc")

	u := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(c.User, c.Password),
		Host:     fmt.Sprintf("%s:%s", c.Host, c.Port),
		Path:     c.Name,
		RawQuery: q.Encode(),
	}
	return u.String()
}
