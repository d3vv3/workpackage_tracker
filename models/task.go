package models

import (
	"encoding/json"
	"fmt"
	"regexp"
	"time"
)

type Task struct {
	ID    int      `json:"id"`
	Start ISO8601  `json:"start"`
	End   ISO8601  `json:"end"`
	Tags  []string `json:"tags"`
}

func (t Task) String() string {
	return fmt.Sprintf("Task %d: %s -> %s | %s", t.ID, t.Start, t.End, t.WorkPackage())
}

func (t Task) WorkPackage() string {
	re := regexp.MustCompile(`(WP\s+\d+)`)
	// TODO: Handle multiple tags
	return re.FindString(t.Tags[0])
}

func (t Task) Duration() time.Duration {
	return t.End.Time().Sub(t.Start.Time())
}

func ParseTasks(input []byte) ([]Task, error) {
	var tasks []Task
	err := json.Unmarshal(input, &tasks)
	return tasks, err
}
