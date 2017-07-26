package types

type Holiday struct {
	Name string `json:"name"`
	Date string `json:"date"`
	Type int    `json:"type"`
}

type Db struct {
	Holidays []Holiday `json: "holidays"`
}
