package resource

type Resource struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `yaml:"repo_link" json:"repo_link"`
}
