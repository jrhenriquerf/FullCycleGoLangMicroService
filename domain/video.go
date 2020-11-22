package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type Video struct {
	ID        string    `valid:"uuid"`
	ResouceID string    `valid:"notnull"`
	FilePath  string    `valid:"notnull"`
	CreatedAt time.Time `valid:"-"`
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
