package cli

import (
	"fmt"

	"github.com/tkuchiki/unordiff/internal/diff"
)

type JSONCmd struct {
	File1 string `arg:"" required:"" help:"First JSON file to compare."`
	File2 string `arg:"" required:"" help:"Second JSON file to compare."`
}

func (j *JSONCmd) Run() error {
	fmt.Printf("Comparing JSON files: %s and %s\n", j.File1, j.File2)
	differences, err := diff.CompareJSONFiles(j.File1, j.File2)
	if err != nil {
		return err
	}
	fmt.Println("Differences:", differences)
	return nil
}
