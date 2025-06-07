package confige

type PortConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"pass"`
	DBName   string `yaml:"dbname"`
	SSlMode  string `yaml:"sslmode"`
}
