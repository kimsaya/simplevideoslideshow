package assets

type Content struct {
	Name     string `json:"name"`
	ItemType string `json:"item_type"`
	Url      string `json:"url"`
	Duration int    `json:"duration"`
}
