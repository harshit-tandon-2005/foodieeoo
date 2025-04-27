package models

// Config struct to hold configuration data
type Config struct {
	ApplicationPort  int    `yaml:"APPLICATION_PORT"`
	DatabasePort     string `yaml:"DATABASE_PORT"`
	DatabaseHost     string `yaml:"DATABASE_HOST"`
	DatabaseName     string `yaml:"DATABASE_NAME"`
	DatabaseUser     string `yaml:"DATABASE_USER"`
	DatabasePassword string `yaml:"DATABASE_PASSWORD"`

	// Add other configuration fields here if needed
}
