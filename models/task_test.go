package models

import (
	"testing"
	"time"
)

func TestTaskString(t *testing.T) {
	startTime, _ := time.Parse("20060102T150405Z", "20250303T081000Z")
	endTime, _ := time.Parse("20060102T150405Z", "20250303T090511Z")
	task := Task{
		ID:    1,
		Start: ISO8601(startTime),
		End:   ISO8601(endTime),
		Tags:  []string{"WP 123: A test task"},
	}

	expected := "Task 1: 2025-03-03T08:10:00Z -> 2025-03-03T09:05:11Z | WP 123"
	if task.String() != expected {
		t.Errorf("Expected %s, got %s", expected, task.String())
	}
}

func TestTaskWorkPackage(t *testing.T) {
	task := Task{
		ID:    1,
		Start: ISO8601(time.Time{}),
		End:   ISO8601(time.Time{}),
		Tags:  []string{"WP 2291: A test task"},
	}

	if task.WorkPackage() != "WP 2291" {
		t.Errorf("Expected WP 2291, got %s", task.WorkPackage())
	}
}

func TestTaskDuration(t *testing.T) {
	startTime, _ := time.Parse("20060102T150405Z", "20250303T081000Z")
	endTime, _ := time.Parse("20060102T150405Z", "20250303T090511Z")
	task := Task{
		ID:    1,
		Start: ISO8601(startTime),
		End:   ISO8601(endTime),
		Tags:  []string{"WP 123: Another test task"},
	}

	expected := time.Duration(0*time.Hour + 55*time.Minute + 11*time.Second)
	if task.Duration() != expected {
		t.Errorf("Expected %s, got %s", expected, task.Duration())
	}
}

func TestParseTasks(t *testing.T) {
	input := []byte(`[
        {
            "id": 1,
            "start": "20250303T081000Z",
            "end": "20250303T090511Z",
            "tags": ["WP 123: A test task"]
        },
        {
            "id": 2,
            "start": "20250303T091000Z",
            "end": "20250303T101511Z",
            "tags": ["WP 123: Another test task"]
        }
    ]`)

	tasks, err := ParseTasks(input)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if len(tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(tasks))
	}

	if tasks[0].ID != 1 {
		t.Errorf("Expected ID 1, got %d", tasks[0].ID)
	}

	if tasks[1].ID != 2 {
		t.Errorf("Expected ID 2, got %d", tasks[1].ID)
	}
}
