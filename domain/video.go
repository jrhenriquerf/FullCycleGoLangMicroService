package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type Video struct {
	ID        string    `json:"encoded_video_folder" valid:"uuid" gorm:"type:uuid;primary_key"`
	ResouceID string    `json:"resource_id" valid:"notnull" gorm:"type:varchar(255)"`
	FilePath  string    `json:"file_path" valid:"notnull" gorm:"type:varchar(255)"`
	CreatedAt time.Time `json:"-" valid:"-"`
	Jobs      []*Job    `json:"-" valid:"-" gorm:"ForeignKey:VideoID"`
}

func init() {
	// Faz a validação sempre que for criada uma nova instância de Video
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewVideo() *Video {
	return &Video{}
}

func (video *Video) Validate() error {
	// Valida a struct de acordo com as tags de validação da estrutura
	_, err := govalidator.ValidateStruct(video)

	if err != nil {
		return err
	}

	return nil
}
