package structs

type OFQ struct {
	SateliteList []Satelite `json:"satellites"`
}

type Satelite struct {
	Name string `json:"name"`
	*SateliteInfo
}

type SateliteInfo struct {
	Distance float32  `json:"distance"`
	Message  []string `json:"message"`
}

type CoordPos struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

type ResponseMessage struct {
	Position CoordPos `json:"position"`
	Message  string   `json:"message"`
}
