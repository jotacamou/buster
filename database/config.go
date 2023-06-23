package database

// DatabaseConfig represents the database configuration as
// defined by key in the config/database.json file.
type DatabaseConfig struct {
	Host     string `json:"host"`
	Database string `json:"database"`
	User     string `json:"user"`
	Password string `json:"password"`
	Driver   string `json:"driver"`
	Port     int    `json:"port"`
}
