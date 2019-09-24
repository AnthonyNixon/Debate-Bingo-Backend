package types

type Field struct {
	Content string `json:"content"`
}

type Config struct {
	Fields []Field `json:"fields"`
}

type Coordinates struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Box struct {
	Content string `json:"content"`
	Checked bool `json:"checked"`
	Coordinates Coordinates `json:"coordinates"`
}