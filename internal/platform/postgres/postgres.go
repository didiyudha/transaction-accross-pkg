package postgres

import (
	"database/sql"
	"net/url"
	_ "github.com/lib/pq"
)

type Config struct {
	User       string `mapstructure:"user"`
	Password   string `mapstructure:"password"`
	Host       string `mapstructure:"host"`
	Name       string `mapstructure:"name"`
	DisableTLS bool   `mapstructure:"disableTLS"`
}

func Open(cfg Config) (*sql.DB, error) {

	// Define SSL mode.
	sslMode := "require"
	if cfg.DisableTLS {
		sslMode = "disable"
	}

	// Query parameters.
	q := make(url.Values)
	q.Set("sslmode", sslMode)
	q.Set("timezone", "utc")

	// Construct URL.
	u := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(cfg.User, cfg.Password),
		Host:     cfg.Host,
		Path:     cfg.Name,
		RawQuery: q.Encode(),
	}

	return sql.Open("postgres", u.String())
}