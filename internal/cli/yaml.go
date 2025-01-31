package cli

import (
	"fmt"

	"github.com/tkuchiki/unordiff/internal/diff"
)

type YAMLCmd struct {
	File1 string `arg:"" required:"" help:"First YAML file to compare."`
	File2 string `arg:"" required:"" help:"Second YAML file to compare."`
}

func (y *YAMLCmd) Run() error {
	fmt.Printf("Comparing YAML files: %s and %s\n", y.File1, y.File2)
	differences, err := diff.CompareYAMLFiles(y.File1, y.File2)
	if err != nil {
		return err
	}
	fmt.Println("Differences:", differences)
	return nil
}
