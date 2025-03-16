package models

import "time"

type WorkPackage struct {
	ID        string
	TimeSpent time.Duration
}

func ParseWorkPackages(tasks []Task) []WorkPackage {
	var workPackages []WorkPackage
	for _, task := range tasks {
		if task.WorkPackage() == "" {
			continue
		}
		workPackages = append(workPackages, WorkPackage{ID: task.WorkPackage(), TimeSpent: task.Duration()})
	}
	return workPackages
}

func CondensedWorkPackages(workPackages []WorkPackage) (map[string]time.Duration, error) {
	result := make(map[string]time.Duration)

	for _, wp := range workPackages {
		result[wp.ID] += wp.TimeSpent
	}

	return result, nil
}
