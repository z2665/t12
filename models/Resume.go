package models

type Resume struct {
	TrueName  string            `json:"trueName"`
	Specialty string            `json:"specialty"`
	Content   string            `json:"content"`
	Contact   map[string]string `json:"contact"`
}
