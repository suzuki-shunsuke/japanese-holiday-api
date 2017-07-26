package types

type Holiday struct {
	Name string `json:"name"`
	Date string `json:"date"`
	Type int    `json:"type"`
}

type Db struct {
	Holidays []Holiday `json: "holidays"`
}

type RDBConfig struct {
	Dbms     string
	User     string
	Password string
	Host     string
	Port     int
	DBName   string
	Protocol string
}

type Config struct {
	RDB RDBConfig
}
