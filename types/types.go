package types

type Holiday struct {
	Name      string `json:"name"`
	Date      string `json:"date"`
	Type      int    `json:"type"`
	DayOfWeek int    `json:"day_of_week"`
}

type Holidays []Holiday

func (h Holidays) Len() int {
	return len(h)
}

func (h Holidays) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Holidays) Less(i, j int) bool {
	return h[i].Date < h[j].Date
}

type Db struct {
	Holidays []Holiday `json: "holidays"`
}

type AppConfig struct {
	Port int
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
	App AppConfig
}

type Request struct {
	From  string `json:"from" form:"from" query:"from"`
	To    string `json:"to" form:"to" query:"to"`
	Types []int  `json:"types" form:"types" query:"types"`
}
