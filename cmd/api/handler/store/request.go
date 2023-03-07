package store

type User struct {
	Name     string `json:"Name"`
	Lastname string `json:"lastname"`
	Datetime string `json:"datetime"`
	Address  string `json:"address"`
	Gender   string `json:"gender"`
	Age      int    `json:"age"`
}
