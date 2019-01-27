package git

import (
	"os/exec"
	"regexp"
)

var changedFiles = regexp.MustCompile(`[ MADRCU]{2}\s+(.*)`)

type StatusOutput struct {
	changedFiles []string
}

func Status() (*StatusOutput, error) {
	out, err := exec.Command("git", "status", "--short").Output()
	if err != nil {
		return nil, err
	}

	status := &StatusOutput{}

	result := changedFiles.FindAllStringSubmatch(string(out), -1)

	for i := range result {
		status.changedFiles = append(status.changedFiles, result[i][1])
	}

	return status, nil
}

func (s *StatusOutput) HasChanges() bool {
	return len(s.changedFiles) > 0
}

func Checkout(options ...string) error {
	args := []string{"checkout"}
	return exec.Command("git", append(args, options...)...).Run()
}

func Fetch() error {
	return exec.Command("git", "fetch").Run()
}

func Reset(mode string, source string) error {
	return exec.Command("git", "reset", mode, source).Run()
}
