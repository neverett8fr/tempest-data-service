package storage

type FileInformation struct {
	Name string `json:"name"`
	Ext  string `json:"extension"`
	Size int64  `json:"size"`
}
