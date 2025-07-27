package config

type Config struct {
	Routes []Route `yaml:"routes"`
}

type Route struct {
	Path      string     `yaml:"path"`
	Target    string     `yaml:"target"`
	RateLimit *RateLimit `yaml:"rate_limit,omitempty"`
}

type RateLimit struct {
	Rate  int64 `yaml:"rate"`
	Burst int64 `yaml:"burst"`
}
