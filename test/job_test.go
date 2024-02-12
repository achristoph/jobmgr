package test

import (
	"fmt"
	jobmgr "jobmgr/job"
	"reflect"
	"testing"
)

func TestJob_Run_Success(t *testing.T) {
	commands := []string{"echo me", "echo 10"}
	job := jobmgr.NewJob(commands)
	job.Run()
	expectedOutputs := []string{"me", "10"}
	fmt.Println(job.Outputs, expectedOutputs)
	if !reflect.DeepEqual(job.Outputs, expectedOutputs) {
		t.Errorf("Expected outputs are not equal")
	}
}

func TestJob_Run_Failure(t *testing.T) {
	commands := []string{"unknown command"}
	job := jobmgr.NewJob(commands)
	err := job.Run()
	if err == nil {
		t.Errorf("No error thrown when one is expected")
	}
}
