package jobmgr

import (
	"os/exec"
	"strings"
)

type Manager interface {
	Run()
}

type Job struct {
	Commands []string
	Outputs  []string
	Manager
}

func NewJob(commands []string) *Job {
	return &Job{Commands: commands}
}

func (j *Job) Run() (err error) {
	for _, command := range j.Commands {
		arr := strings.Split(command, " ")
		cmd := exec.Command(arr[0], arr[1:]...)
		output, err := cmd.Output()
		if err != nil {
			return err
		}
		// Need to trim the output since a newline character is added
		j.Outputs = append(j.Outputs, strings.TrimSpace(string(output)))
	}
	return err
}
