package resource

type Resource struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	URL           string `yaml:"repo" json:"repo"`
	Image         string `yaml:"container_image"`
	Owner         string `json:"username"`
	Stars         string `json:"stars"`
	StarsCount    int    `json:"stars_count"`
	NameWithOwner string
}
