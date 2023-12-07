package models

type Media struct {
	Base
	URL string `gorm:"not null;type:text" json:"url"`
}

type GetPresignedUrlRequestDto struct {
	FileName string `json:"fileName"`
	FileType string `json:"fileType"`
	Path     string `json:"path"`
}
