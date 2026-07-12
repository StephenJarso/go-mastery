package orm

import (
	"time"

	"gorm.io/gorm"
)

// GORM is the developer-friendly ORM (Object-Relational Mapping) library for Go.
// It maps Go structs to SQL tables automatically.
// gorm.Model is a pre-defined struct containing fields: ID, CreatedAt, UpdatedAt, DeletedAt.
// Embedding it in a struct automatically adds these columns and enables soft-delete!

type User struct {
	gorm.Model
	Name  string  `gorm:"type:varchar(100);not null"`
	Email *string `gorm:"type:varchar(100);uniqueIndex"` // Pointer allows storing NULL in database
	Role  string  `gorm:"type:varchar(20);default:'user'"`
	
	// One-to-One relationship
	// GORM will automatically look for Profile's foreign key (UserID)
	Profile Profile `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	// One-to-Many relationship (User has many Posts)
	Posts []Post `gorm:"foreignKey:AuthorID"`
}

type Profile struct {
	gorm.Model
	UserID int // Foreign key
	Bio    string
	Avatar string
}

type Post struct {
	gorm.Model
	Title    string `gorm:"not null"`
	Body     string `gorm:"type:text"`
	AuthorID uint   // Foreign key (referencing User.ID)

	// Many-to-Many relationship (Post has many Tags, Tag belongs to many Posts)
	// GORM automatically manages the join table "post_tags"
	Tags []Tag `gorm:"many2many:post_tags;"`
}

type Tag struct {
	ID        uint           `gorm:"primaryKey"`
	Name      string         `gorm:"uniqueIndex;not null"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
