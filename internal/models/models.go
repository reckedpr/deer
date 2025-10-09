package models

type ImgJson struct {
	ImgURL string `json:"img_url,omitempty"`
}

type Image struct {
	ImgJson ImgJson
	ImgList []ImgJson
	ImgPath string
}

type FactJson struct {
	Text string `json:"fact"`
}

type Fact struct {
	FactJson     FactJson
	FactList     []FactJson
	FactPath     string
	FactFileSize int64
}
