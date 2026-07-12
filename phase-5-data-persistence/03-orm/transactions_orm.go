package orm

import (
	"fmt"

	"gorm.io/gorm"
)

// GORM provides a very clean transaction API using db.Transaction.
// E.g.:
// db.Transaction(func(tx *gorm.DB) error {
//     // run SQL queries here using 'tx' (not 'db')
//     return nil // Commit transaction automatically
// })
// If the block returns an error, GORM automatically rolls back.
// If a panic occurs, GORM recovers it and rolls back.

// CreateArticleWithTags creates a post and assigns tags to it, ensuring tags exist.
// This is executed inside an atomic transaction block.
func CreateArticleWithTags(db *gorm.DB, authorID uint, title, body string, tagNames []string) (*Post, error) {
	var post Post

	// Start GORM transactional block.
	err := db.Transaction(func(tx *gorm.DB) error {
		// Verify Author exists first
		var author User
		if err := tx.First(&author, authorID).Error; err != nil {
			return fmt.Errorf("author not found: %w", err)
		}

		// Create the Post struct
		post = Post{
			Title:    title,
			Body:     body,
			AuthorID: authorID,
		}

		// Resolve Tags. For each name, we check if it exists, otherwise create it.
		for _, name := range tagNames {
			var t Tag
			// FirstOrCreate fetches the record matching Name, or creates it if not found.
			err := tx.FirstOrCreate(&t, Tag{Name: name}).Error
			if err != nil {
				return fmt.Errorf("failed to get/create tag %s: %w", name, err)
			}
			post.Tags = append(post.Tags, t)
		}

		// Save Post with relations to Tag join table
		if err := tx.Create(&post).Error; err != nil {
			return fmt.Errorf("failed to save post: %w", err)
		}

		return nil // Commits the transaction
	})

	if err != nil {
		return nil, err
	}

	return &post, nil
}
