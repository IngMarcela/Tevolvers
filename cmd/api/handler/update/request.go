package update

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Datetime string `json:"datetime"`
	Address  string `json:"address"`
	Gender   string `json:"gender"`
	Age      int    `json:"age"`
}
