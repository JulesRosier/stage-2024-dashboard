package settings

type Database struct {
	User     string `yaml:"user"`
	Paddword string `yaml:"password"`
	Database string `yaml:"database"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}

func (d *Database) SetDefault() {
	d.Port = 5432
	d.Host = "127.0.0.1"
	d.Paddword = "postgres"
	d.User = "postgres"
	d.Database = "event-viewer"
}
