package controller

type Appeal struct {
	Id     string `json:"id`
	Index  string `json:"index"`
	Type   string `json:"type"`
	Source Source `json:"source`
}

type Source struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type Query struct {
	Text      string   `json:"query"`
	Index     string   `json:"index"`
	Parametrs Parametr `json:"parametrs"`
}

type Parametr struct {
	TypeSearch string `json:"typeSearch"`
}
