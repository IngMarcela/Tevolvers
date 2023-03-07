package list

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"Name"`
	Lastname string `json:"lastname"`
	Datetime string `json:"datetime"`
	Address  string `json:"address"`
	Gender   string `json:"gender"`
	Age      int    `json:"age"`
}
