package orm

import (
	"testing"
)

func TestORMStudentWorkflow(t *testing.T) {
	// Initialize GORM in-memory SQLite
	db, err := InitORM(":memory:")
	if err != nil {
		t.Fatalf("failed to initialize GORM: %v", err)
	}

	// 1. Create User and Profile
	u, err := CreateUserAndProfile(db, "StephenJarso", "stephen@example.com", "Golang developer")
	if err != nil {
		t.Fatalf("CreateUserAndProfile failed: %v", err)
	}

	if u.ID == 0 {
		t.Error("expected non-zero ID for saved user")
	}

	// Verify profile is saved
	var profile Profile
	err = db.Where("user_id = ?", u.ID).First(&profile).Error
	if err != nil {
		t.Fatalf("failed to find associated profile: %v", err)
	}
	if profile.Bio != "Golang developer" {
		t.Errorf("expected bio 'Golang developer', got %q", profile.Bio)
	}

	// 2. Transactional Post creation with tags
	post, err := CreateArticleWithTags(db, u.ID, "GORM Mastery", "Learning database/sql and GORM in Go", []string{"go", "database", "orm"})
	if err != nil {
		t.Fatalf("CreateArticleWithTags failed: %v", err)
	}

	if post.ID == 0 {
		t.Error("expected non-zero ID for saved post")
	}

	// 3. Eager Loading Associations (Preload)
	uLoaded, err := GetUserWithAssociations(db, u.ID)
	if err != nil {
		t.Fatalf("GetUserWithAssociations failed: %v", err)
	}

	if uLoaded.Profile.Bio != "Golang developer" {
		t.Errorf("eager load failed for profile: got bio %q", uLoaded.Profile.Bio)
	}

	if len(uLoaded.Posts) != 1 {
		t.Errorf("expected 1 preloaded post, got %d", len(uLoaded.Posts))
	} else {
		postLoaded := uLoaded.Posts[0]
		if len(postLoaded.Tags) != 3 {
			t.Errorf("expected 3 preloaded tags for post, got %d", len(postLoaded.Tags))
		}
	}

	// 4. Update Role
	err = UpdateUserRole(db, u.ID, "admin")
	if err != nil {
		t.Fatalf("UpdateUserRole failed: %v", err)
	}

	var uCheck User
	db.First(&uCheck, u.ID)
	if uCheck.Role != "admin" {
		t.Errorf("expected role 'admin', got %q", uCheck.Role)
	}

	// 5. Soft Delete
	err = SoftDeleteUser(db, u.ID)
	if err != nil {
		t.Fatalf("SoftDeleteUser failed: %v", err)
	}

	// User should not be retrievable by normal queries
	var uNotFound User
	err = db.First(&uNotFound, u.ID).Error
	if err == nil {
		t.Error("expected record to not be found after soft delete")
	}

	// But remains in database
	var uSoftDeleted User
	err = db.Unscoped().First(&uSoftDeleted, u.ID).Error
	if err != nil {
		t.Errorf("failed to retrieve soft-deleted user via Unscoped: %v", err)
	}

	// 6. Hard Delete
	err = HardDeleteUser(db, u.ID)
	if err != nil {
		t.Fatalf("HardDeleteUser failed: %v", err)
	}

	// Even Unscoped search should fail now
	var uHardDeleted User
	err = db.Unscoped().First(&uHardDeleted, u.ID).Error
	if err == nil {
		t.Error("expected unscoped search to fail after hard delete")
	}
}
