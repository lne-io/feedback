package models

import (
	"gorm.io/gorm"
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
)

type Website struct {
	ID			string `grom:"size:36;primaryKey" json:"id"`
	Name		string `gorm:"size:40;not null" json:"name"`
	Url			string `gorm:"size:255" json:"url"`
	Feedback	[]Feedback `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"feedback"`  
}


func (w *Website) Validate() error {
	return validation.ValidateStruct(w,
		validation.Field(&w.Name, validation.Required, validation.Length(1,40)),
		validation.Field(&w.Url, validation.Length(0,255), is.URL),
	)
}

func (w *Website) BeforeCreate(tx *gorm.DB) error {
	w.ID = uuid.New().String()
	return nil
  }