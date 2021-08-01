package upload

type Uploaded struct {
	Filename string `json:"filename"`
	Status   int    `json:"status"`
}
