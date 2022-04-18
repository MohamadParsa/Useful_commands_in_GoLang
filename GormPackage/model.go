package gormpackage

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint // field named `ID` will be used as a primary field by default
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time //For models having CreatedAt field, the field will be set to the current time when the record is first created if its value is zero
	UpdatedAt    time.Time //For models having UpdatedAt field, the field will be set to the current time when the record is updated or created if its value is zero
}

// GORM defined a gorm.Model struct, which includes fields ID, CreatedAt, UpdatedAt, DeletedAt
// gorm.Model definition
type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
type UserWithOtherColName struct {
	CreatedAt time.Time // Set to current time if it is zero on creating
	UpdatedAt int       // Set to current unix seconds on updating or if it is zero on creating
	Updated   int64     `gorm:"autoUpdateTime:nano"`  // Use unix nano seconds as updating time
	Updated1  int64     `gorm:"autoUpdateTime:milli"` // Use unix milli seconds as updating time
	Created   int64     `gorm:"autoCreateTime"`       // Use unix seconds as creating time
}

// Set field `UUID` as primary field
type Animal struct {
	ID   int64
	UUID string `gorm:"primaryKey"`
	Name string
	Age  int64
}

type AnimalCustome struct {
	AnimalID int64     `gorm:"column:beast_id"`         // set name to `beast_id`
	Birthday time.Time `gorm:"column:day_of_the_beast"` // set name to `day_of_the_beast`
	Age      int64     `gorm:"column:age_of_the_beast"` // set name to `age_of_the_beast`
}

// GORM pluralizes struct name to snake_cases as table name,
// for struct User, its table name is users by convention

//You can change the default table name by implementing the Tabler interface, for example:
type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (User) TableName() string {
	return "profiles"
}

// Field-Level Permission
// Exported fields have all permission when doing CRUD with GORM, and GORM allows you to change the
// field-level permission with tag, so you can make a field to be read-only, write-only, create-only,
// update-only or ignored

type UserWithPermission struct {
	Name1 string `gorm:"<-:create"`          // allow read and create
	Name2 string `gorm:"<-:update"`          // allow read and update
	Name3 string `gorm:"<-"`                 // allow read and write (create and update)
	Name4 string `gorm:"<-:false"`           // allow read, disable write permission
	Name5 string `gorm:"->"`                 // readonly (disable write permission unless it configured)
	Name6 string `gorm:"->;<-:create"`       // allow read and create
	Name7 string `gorm:"->:false;<-:create"` // createonly (disabled read from db)
	Name8 string `gorm:"-"`                  // ignore this field when write and read with struct
	Name9 string `gorm:"-:all"`              // ignore this field when write, read and migrate with struct
	Name0 string `gorm:"-:migration"`        // ignore this field when migrate with struct
}
