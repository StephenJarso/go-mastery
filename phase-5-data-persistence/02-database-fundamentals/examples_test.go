package dbfundamentals

import (
	"testing"
)

func TestStudentDatabaseWorkflow(t *testing.T) {
	// DSN ":memory:" creates a temporary database fully inside the system RAM.
	db, err := ConnectDatabase(":memory:")
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}
	defer db.Close()

	// 1. Create table
	err = CreateStudentsTable(db)
	if err != nil {
		t.Fatalf("failed to create students table: %v", err)
	}

	// 2. Insert student
	sID, err := InsertStudent(db, "Stephen", 22, "A")
	if err != nil {
		t.Fatalf("failed to insert student: %v", err)
	}

	// 3. Get student by ID
	s, err := GetStudentByID(db, sID)
	if err != nil {
		t.Fatalf("failed to get student: %v", err)
	}
	if s.Name != "Stephen" || s.Grade != "A" {
		t.Errorf("unexpected student data: %+v", s)
	}

	// 4. Prepared batch insertion
	batch := []Student{
		{Name: "Jacob", Age: 20, Grade: "B"},
		{Name: "Alex", Age: 21, Grade: "B"},
	}
	err = InsertStudentsPrepared(db, batch)
	if err != nil {
		t.Fatalf("prepared insert failed: %v", err)
	}

	// 5. Query multiple rows
	bStudents, err := GetStudentsByGrade(db, "B")
	if err != nil {
		t.Fatalf("failed to query students by grade: %v", err)
	}
	if len(bStudents) != 2 {
		t.Errorf("expected 2 students with grade B, got %d", len(bStudents))
	}

	// 6. Transaction execution
	err = TransferStudentGradeTransaction(db, sID, "A+")
	if err != nil {
		t.Fatalf("transaction failed: %v", err)
	}

	// Verify updated grade
	sUpdated, _ := GetStudentByID(db, sID)
	if sUpdated.Grade != "A+" {
		t.Errorf("expected grade 'A+', got %q", sUpdated.Grade)
	}

	// Verify audit log entry
	var logCount int
	err = db.QueryRow("SELECT COUNT(*) FROM audit_log WHERE student_id = ?", sID).Scan(&logCount)
	if err != nil || logCount != 1 {
		t.Errorf("expected 1 audit log entry, got %d (err: %v)", logCount, err)
	}
}
