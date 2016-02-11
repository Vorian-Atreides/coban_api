package databases

import "time"

// Model is the interface for the items in the database
type Model interface {
	IsValid() error
	LoadRelated()
}

// Address {
//	"id":"1",
//	"zip":"1040061"
//	"street":"ぎんざ",
//	"city":"ときょ"
//}
//
type Address struct {
	ID uint `gorm:"column:id;primary_key" json:"id"`

	Zip    string `gorm:"column:zip;" sql:"not null" json:"zip"`
	Street string `gorm:"column:street;" sql:"not null" json:"street"`
	City   string `gorm:"column:city;" sql:"not null" json:"city"`

	Company   *Company `json:"-"`
	CompanyID uint     `gorm:"column:company_id;" sql:"not null" json:"company-id"`
}

// Company {
//	"id":"1",
//	"name":"コバン",
//	"address":
//	{
//		"id":"1",
//		"zip":"1040061"
//		"street":"ぎんざ",
//		"city":"ときょ"
//	}
//}
//
type Company struct {
	ID uint `gorm:"column:id; primary_key" json:"id"`

	Name string `gorm:"column:name;" sql:"not null; unique" json:"name"`

	Addresses []Address `json:"-"`
	Employees []User    `json:"-"`
}

// JSON Device
//
//{
//	"is-paired":"false"
//}

type Device struct {
	ID uint `gorm:"column:id; primary_key" json:"-"`

	IsPaired bool `gorm:"column:is_paired;" sql:"not null;" json:"is-paired"`

	User   *User `json:"-"`
	UserID uint  `gorm:"column:user_id;" sql:"not null" json:"-"`
}

// Account {
//	"email":"tatsuya@coban.jp"
//}
//
type Account struct {
	ID uint `gorm:"column:id; primary_key" json:"-"`

	Email    string `gorm:"column:email;" sql:"not null; unique" json:"email"`
	Password string `gorm:"column:password;" sql:"not null" json:"-"`
	Scope    byte   `gorm:"column:scope;" sql:"not null" json:"-"`

	User *User `sql:"not null" json:"-"`
}

// User {
//	"id":"1",
//	"first-name":"たつや",
//	"last-name":"ぜんぶつ",
//	"account":
//	{
//		"email":"tatsuya@coban.jp"
//	},
//	"company":
//	{
//		"id":"1",
//		"name":"コバン"
//	},
//	"device":
//	{
//		"is-paired":"false"
//	}
//}
//
type User struct {
	ID uint `gorm:"column:id; primary_key" json:"id"`

	FirstName string `gorm:"column:first_name;" sql:"not null;" json:"first-name"`
	LastName  string `gorm:"column:last_name;" sql:"not null;" json:"last-name"`

	Account   *Account `json:"account; omitempty"`
	AccountID uint     `gorm:"column:account_id;" sql:"not null;" json:"-"`

	Company   *Company `json:"company; omitempty"`
	CompanyID uint     `gorm:"column:company_id;" sql:"not null;" json:"-"`

	Device *Device `json:"device; omitempty"`
}

// Station {
//	"name":"ときょ"
//	"type":"Metro"
//}
//
type Station struct {
	ID uint `gorm:"column:id; primary_key" json:"-"`

	// Name string `gorm:"column:name;" sql:"not null;" json:"name"`
	// Type string `gorm:"column:type;" sql:"not null;" json:"type"`

	LineCode    uint   `gorm:"column:line_code;" sql:"not null;" json:"-"`
	StationCode uint   `gorm:"column:station_code;" sql:"not null;" json:"-"`
	Company     string `gorm:"column:company;" sql:"not null;" json:"company"`
	Line        string `gorm:"column:line;" sql:"not null;" json:"line"`
	Name        string `gorm:"column:name;" sql:"not null;" json:"station"`
}

// TransportHistory {
//	"id":"1",
//	"date":"2016-01-29T06:03:26+09:00",
//	"stock":"3500",
//	"expense":"193",
//	"entrance":
//	{
//		"name":"ときょ"
//		"type":"Metro"
//	},
//	"exit":
//	{
//		"name":"しんじゅく"
//		"type":"Metro"
//	}
//}
//
type TransportHistory struct {
	ID uint `gorm:"column:id; primary_key" json:"id"`

	Date    time.Time `gorm:"column:date;" sql:"not null;" json:"date"`
	Stock   uint      `gorm:"column:stock;" sql:"not null;" json:"stock"`
	Expense uint      `gorm:"column:expense;" sql:"not null;" json:"expense"`

	Entrance   *Station `json:"entrance; omitempty"`
	EntranceID uint     `gorm:"column:entrance_id;" sql:"not null;" json:"-"`

	Exit   *Station `json:"exit; omitempty"`
	ExitID uint     `gorm:"column:exit_id;" sql:"not null;" json:"-"`

	User   *User `json:"user; omitempty"`
	UserID uint  `gorm:"column:user_id;" sql:"not null;" json:"-"`
}
