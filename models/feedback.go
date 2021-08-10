package models

import (
	"github.com/go-ozzo/ozzo-validation/v4"
)

type Feedback struct {
	ID				uint `grom:"primaryKey" json:"id"`
	Subject			string `gorm:"size:60;not null" json:"subject"`
	Description		string `gorm:"not null" json:"Description"`
	OS				string `grom:"size:40;not null" json:"os"`
	Browser			string `gorm:"size:40;not null" json:"browser"`
	BrowserVersion	string `gorm:"size:20;not null" json:"browser_version"`
	WebsiteID		string `gorm:"size:36;not null" json:"website_id"`  
}


func (f *Feedback) Validate() error {
	return validation.ValidateStruct(f,
		validation.Field(&f.Subject, validation.Required, validation.Length(2,60)),
		validation.Field(&f.Description, validation.Required),
		validation.Field(&f.OS, validation.Required, validation.Length(1,40)),
		validation.Field(&f.Browser, validation.Required, validation.Length(1,40)),
		validation.Field(&f.BrowserVersion, validation.Required, validation.Length(1,20)),
		validation.Field(&f.WebsiteID, validation.Required, validation.Length(36,36)),
	)
}