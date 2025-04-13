package config

const (
	BaseURL   = "https://apix.snappshop.ir"
	Latitude  = "35.77331"
	Longitude = "51.418591"
	Worker    = 8
)

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}
