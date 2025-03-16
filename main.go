package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"text/tabwriter"
	"time"

	"github.com/d3vv3/workpackage_tracker/models"
)

func readInput() []byte {
	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}

func parseInput(stdInput []byte) ([]byte, []byte) {
	config, after, _ := bytes.Cut(stdInput, []byte("\n\n"))
	return config, after
}

func main() {
	bytes := readInput()
	// NOTE: We do not use the config
	_, input := parseInput(bytes)
	tasks, err := models.ParseTasks(input)
	if err != nil {
		log.Fatalf("Error parsing tasks: %v", err)
	}
	// NOTE: In case we wanted to
	// log.Printf("Config: \n%s", string(config))
	workPackages := models.ParseWorkPackages(tasks)
	summary, err := models.CondensedWorkPackages(workPackages)
	if err != nil {
		log.Fatalf("Error summarizing work packages: %v", err)
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 1, 4, ' ', 0)
	fmt.Fprintf(w, "Work Package\tDuration\n")
	var total time.Duration
	for id, duration := range summary {
		total += duration
		fmt.Fprintf(w, "%s\t%s\n", id, duration)
	}
	w.Flush()
	fmt.Printf("\nTotal\t%s\n", total)
}
