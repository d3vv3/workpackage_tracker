package models

import (
	"testing"
	"time"
)

func TestParseWorkPackages(t *testing.T) {
	input := []byte(`[{"id":1,"start":"20250303T081000Z","end":"20250303T081500Z","tags":["WP 1 - Test"]}]`)
	tasks, _ := ParseTasks(input)
	workPackages := ParseWorkPackages(tasks)
	if len(workPackages) != 1 {
		t.Errorf("Expected 1, got %d", len(workPackages))
	}
	if workPackages[0].ID != "WP 1" {
		t.Errorf("Expected WP 1, got %s", workPackages[0].ID)
	}
	if workPackages[0].TimeSpent.String() != "5m0s" {
		t.Errorf("Expected 5m0s, got %s", workPackages[0].TimeSpent.String())
	}
}

func TestCondensedWorkPackages(t *testing.T) {
	workPackages := []WorkPackage{
		{ID: "WP 1", TimeSpent: time.Duration(5 * time.Minute)},
		{ID: "WP 1", TimeSpent: time.Duration(10 * time.Minute)},
		{ID: "WP 2", TimeSpent: time.Duration(15 * time.Minute)},
	}
	result, _ := CondensedWorkPackages(workPackages)
	if len(result) != 2 {
		t.Errorf("Expected 2, got %d", len(result))
	}
	if result["WP 1"].String() != "15m0s" {
		t.Errorf("Expected 15m0s, got %s", result["WP 1"].String())
	}
	if result["WP 2"].String() != "15m0s" {
		t.Errorf("Expected 15m0s, got %s", result["WP 2"].String())
	}
}
