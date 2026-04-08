package config

import "strings"

type Config struct {
	Port           string
	DatabaseURL    string
	AllowedOrigins []string
}

func FromEnv(env []string) Config {
	m := make(map[string]string, len(env))
	for _, e := range env {
		parts := strings.SplitN(e, "=", 2)
		if len(parts) == 2 {
			m[parts[0]] = parts[1]
		}
	}

	port := m["PORT"]
	if port == "" {
		port = "8080"
	}

	origins := m["ALLOWED_ORIGINS"]
	if origins == "" {
		origins = "http://localhost:3000"
	}

	return Config{
		Port:           port,
		DatabaseURL:    m["DATABASE_URL"],
		AllowedOrigins: strings.Split(origins, ","),
	}
}
