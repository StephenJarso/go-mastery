package advanced_testing

import (
	"errors"
	"testing"
)

// MockUserStore is a manual mock implementation of UserStore.
// It tracks calls and allows configuring return values or errors.
type MockUserStore struct {
	SavedUser *User
	SaveError error
	GetResult *User
	GetError  error
	SaveCalls int
	GetCalls  int
}

func (m *MockUserStore) Get(id int) (*User, error) {
	m.GetCalls++
	return m.GetResult, m.GetError
}

func (m *MockUserStore) Save(user *User) error {
	m.SaveCalls++
	if m.SaveError != nil {
		return m.SaveError
	}
	m.SavedUser = user
	user.ID = 100 // Simulate DB auto-generating an ID
	return nil
}

func TestRegisterUser(t *testing.T) {
	t.Run("successful registration", func(t *testing.T) {
		mockStore := &MockUserStore{}
		service := &UserService{Store: mockStore}

		user, err := service.RegisterUser("Alice", "alice@example.com")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if mockStore.SaveCalls != 1 {
			t.Errorf("expected 1 Save call, got %d", mockStore.SaveCalls)
		}

		if user.ID != 100 {
			t.Errorf("expected generated ID 100, got %d", user.ID)
		}

		if mockStore.SavedUser.Name != "Alice" || mockStore.SavedUser.Email != "alice@example.com" {
			t.Errorf("saved user fields mismatch: %+v", mockStore.SavedUser)
		}
	})

	t.Run("validation failure", func(t *testing.T) {
		mockStore := &MockUserStore{}
		service := &UserService{Store: mockStore}

		_, err := service.RegisterUser("Alice", "invalid-email")
		if err == nil {
			t.Fatal("expected validation error, got nil")
		}

		if mockStore.SaveCalls != 0 {
			t.Errorf("expected 0 Save calls on validation failure, got %d", mockStore.SaveCalls)
		}
	})

	t.Run("storage failure", func(t *testing.T) {
		dbErr := errors.New("db connection failure")
		mockStore := &MockUserStore{SaveError: dbErr}
		service := &UserService{Store: mockStore}

		_, err := service.RegisterUser("Alice", "alice@example.com")
		if !errors.Is(err, dbErr) {
			t.Fatalf("expected error wrapping db failure, got: %v", err)
		}

		if mockStore.SaveCalls != 1 {
			t.Errorf("expected 1 Save call, got %d", mockStore.SaveCalls)
		}
	})
}
