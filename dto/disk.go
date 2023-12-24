package dto

type DiskUsage struct {
	Size      string `json:"size"`
	Used      string `json:"used"`
	Available string `json:"available"`
	Usage     string `json:"usage"`
}

type Disk struct {
	FileSystem string    `json:"file_system"`
	Type       string    `json:"type"`
	MountedOn  string    `json:"mounted_on"`
	Usage      DiskUsage `json:"usage"`
}
