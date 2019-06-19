package models

import "errors"

type Contact struct {
	ContactID      int    `gorm:"primary_key"`
	ContactName    string `gorm:"type:varchar(50)”`
	ContactAddress string `gorm:"type:varchar(500)”`
	ContactEmail   string `gorm:"type:varchar(50)”`
	//CustomerId     string
}

type Customer struct {
	CustomerId   int       `gorm:"primary_key"`
	CustomerName int       `gorm:"type:varchar(50)”`
	Contacts     []Contact `gorm:"ForeignKey:CustomerId"`
}

func (c *Contact) Validate() (err error) {

	var message string

	if c.ContactName == "" {
		message = "ContactName,"
	}

	if c.ContactAddress == "" {
		message = "ContactAddress,"
	}

	if c.ContactEmail == "" {
		message = "ContactEmail,"
	}

	if message != "" {
		message = "Required fields: " + message[:len(message)-1] + " Messages:"
	}
	if message != "" {
		return errors.New(message)
	}

	return

}