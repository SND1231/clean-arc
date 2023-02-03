package setting

type Setting struct {
	DB DB
}

type DB struct {
	Type     string
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}
