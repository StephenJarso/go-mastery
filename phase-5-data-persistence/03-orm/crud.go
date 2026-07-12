package orm

import (
	"fmt"

	"gorm.io/gorm"
)

// GORM simplifies CRUD operations by generating SQL commands automatically
// and exposing methods like Create, First, Find, Updates, and Delete.

// CreateUserAndProfile inserts a user and their profile.
// GORM automatically resolves associations and inserts them in a single step!
func CreateUserAndProfile(db *gorm.DB, name, email, bio string) (*User, error) {
	u := User{
		Name:  name,
		Email: &email,
		Profile: Profile{
			Bio: bio,
		},
	}

	// db.Create saves the model. Because Profile is loaded, GORM automatically
	// sets the correct foreign key UserID on Profile and saves both!
	result := db.Create(&u)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to create user and profile: %w", result.Error)
	}

	return &u, nil
}

// GetUserWithAssociations loads a user, their profile, and their posts.
// In GORM, associations are lazy-loaded. To fetch them, we use Preload.
func GetUserWithAssociations(db *gorm.DB, userID uint) (*User, error) {
	var u User
	
	// Preload eager-loads the relational models in secondary queries.
	err := db.Preload("Profile").Preload("Posts").Preload("Posts.Tags").First(&u, userID).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	
	return &u, nil
}

// UpdateUserRole updates the role of a user using struct updates.
func UpdateUserRole(db *gorm.DB, userID uint, newRole string) error {
	var u User
	u.ID = userID

	// Updates() updates specific columns.
	// Note: When updating with struct, GORM will only update non-zero fields.
	// E.g., if you pass string "", it won't update. To update zero values, use Map.
	err := db.Model(&u).Updates(User{Role: newRole}).Error
	if err != nil {
		return fmt.Errorf("failed to update user role: %w", err)
	}
	return nil
}

// SoftDeleteUser deletes a user temporarily (retains row in DB, setting deleted_at).
func SoftDeleteUser(db *gorm.DB, userID uint) error {
	// GORM deletes the record by ID if primary key is set on struct.
	err := db.Delete(&User{}, userID).Error
	if err != nil {
		return fmt.Errorf("failed to soft delete user: %w", err)
	}
	return nil
}

// HardDeleteUser permanently removes a user from the database.
func HardDeleteUser(db *gorm.DB, userID uint) error {
	// Unscoped() disables soft-deleting behavior, executing a raw DELETE.
	err := db.Unscoped().Delete(&User{}, userID).Error
	if err != nil {
		return fmt.Errorf("failed to hard delete user: %w", err)
	}
	return nil
}
