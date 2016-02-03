package databases

import (
	"net/http"
	"time"
)

type Model interface {
	IsValid(forCreation bool) error
	LoadRelated()
	FromBody(r *http.Request) error
}

// JSON Address
//
//{
//	"id":"1",
//	"zip":"1040061"
//	"street":"ぎんざ",
//	"city":"ときょ"
//}

type Address 			struct {
	ID					uint			`gorm:"primary_key" json:"id"`

	Zip					string			`sql:"not null" json:"zip"`
	Street				string			`sql:"not null" json:"street"`
	City				string			`sql:"not null" json:"city"`

	CompanyID			uint			`sql:"not null" json:"company-id"`
}

// JSON Company
//
//{
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

type Company 			struct {
	ID					uint			`gorm:"primary_key" json:"id"`

	Name				string			`sql:"not null; unique" json:"name"`

	Addresses			[]Address		`json:"addresses"`
	Employees			[]User			`json:"-"`
}

// JSON Device
//
//{
//	"is-paired":"false"
//}

type Device 			struct {
	ID					uint 			`gorm:"primary_key" json:"-"`

	IsPaired			bool			`sql:"not null;" json:"is-paired"`

	User				*User
	UserID				uint			`sql:"not null" json:"-"`
}

// JSON Account
//
//{
//	"email":"tatsuya@coban.jp"
//}

type Account 			struct {
	ID					uint			`gorm:"primary_key" json:"-"`

	Email				string			`sql:"not null; unique" json:"email"`
	Password			string			`sql:"not null" json:"-"`
	Scope				byte

	User				*User
}

// JSON User
//
//{
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

type User 				struct {
	ID					uint			`gorm:"primary_key" json:"id"`

	FirstName			string			`sql:"not null;" json:"first-name"`
	LastName			string			`sql:"not null;" json:"last-name"`
	IsManager			bool			`sql:"not null;" json:"-"`

	Account				*Account		`json:"account"`
	AccountID			uint			`sql:"not null;" json:"-"`

	Company				*Company		`json:"company"`
	CompanyID			uint			`sql:"not null;" json:"-"`

	Device				*Device			`json:"device"`
}

// JSON TransportType
//
//{
//	"name":"Metro"
//}

type TransportType 		struct {
	ID					uint 			`gorm:"primary_key" json:"-"`

	Name				string			`sql:"not null; unique" json:"name"`
}

// JSON Station
//
//{
//	"name":"ときょ"
//	"type":
//	{
//		"name":"Metro"
//	}
//}

type Station 			struct {
	ID					uint 			`gorm:"primary_key" json:"-"`

	Name				string			`sql:"not null;" json:"name"`

	Type				TransportType	`json:"type"`
	TypeID				uint			`sql:"not null;" json:"-"`
}

// JSON TransportHistory
//
//{
//	"id":"1",
//	"date":"2016-01-29T06:03:26+09:00",
//	"stock":"3500",
//	"expense":"193",
//	"entrance":
//	{
//		"name":"ときょ"
//		"type":
//		{
//			"name":"Metro"
//		}
//	},
//	"exit":
//	{
//		"name":"しんじゅく"
//		"type":
//		{
//			"name":"Metro"
//		}
//	}
//}

type TransportHistory	struct {
	ID					uint 			`gorm:"primary_key" json:"id"`

	Date				time.Time		`sql:"not null;" json:"date"`
	Stock				uint			`sql:"not null;" json:"stock"`
	Expense				uint			`sql:"not null;" json:"expense"`

	Entrance			Station			`json:"entrance"`
	EntranceID			uint			`sql:"not null;" json:"-"`

	Exit				Station			`json:"exit"`
	ExitID				uint			`sql:"not null;" json:"-"`
}