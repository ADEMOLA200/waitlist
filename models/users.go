package models

import "time"

type Waitlist struct {
    ID        	uint       	`json:"id" gorm:"primaryKey"`
    FullName  	string     	`json:"full_name" gorm:"not null"`
	Email 		string 		`gorm:"varchar(255);uniqueIndex"`
    CreatedAt 	time.Time  	`json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt 	time.Time  	`json:"updated_at" gorm:"autoUpdateTime"`
}

type Campaign struct {
   ID          	uint       `json:"id" gorm:"primaryKey"`
   UID         	string     `json:"uid"`

   WaitlistID   uint `gorm:"foreignKey:WaitlistID"`

   Waitlist  	Waitlist   `json:"waitlist"`
   IsRedeemed  	bool       `json:"is_redeemed"`
   CreatedAt   	time.Time  `json:"created_at"`
   UpdatedAt   	time.Time  `json:"updated_at"`
}


// code for the template renderer
type User struct {
	Name string
	Email string 
  }