package diff

import (
	"fmt"
	"sort"

	"github.com/google/go-cmp/cmp"
)

type Differ struct {
}

func NewDiffer() *Differ {
	return &Differ{}
}

func (d *Differ) Compare(x, y []interface{}) string {

}

func (d *Differ) CompareSlicesUnordered(x, y []interface{}) string {
	if len(x) != len(y) {
		return false
	}

	// スライスをコピーしてソート
	xCopy, yCopy := make([]interface{}, len(x)), make([]interface{}, len(y))
	copy(xCopy, x)
	copy(yCopy, y)

	sort.Slice(xCopy, func(i, j int) bool {
		return fmt.Sprintf("%v", xCopy[i]) < fmt.Sprintf("%v", xCopy[j])
	})
	sort.Slice(yCopy, func(i, j int) bool {
		return fmt.Sprintf("%v", yCopy[i]) < fmt.Sprintf("%v", yCopy[j])
	})

	return cmp.Diff(xCopy, yCopy)
}

type DiffResult struct {
	Source      interface{}
	Destination interface{}
	DiffType    int
}

const (
	diffTypeDiff                  = 0
	diffTypeExistsOnlySource      = 1
	diffTypeExistsOnlyDestination = 2
)

func NewDiffResult(src, dest interface{}, diffType int) *DiffResult {
	return &DiffResult{
		Source:      src,
		Destination: dest,
		DiffType:    diffType,
	}
}

type GenericObject map[string]interface{}

func compareObjectArrays(src, dst []GenericObject) ([]GenericObject, []GenericObject) {
	srcKeysMap := make(map[string]struct{})
	srcKeys := make([]string, 0)
	for _, obj := range src {
		for key, _ := range obj {
			if _, ok := srcKeysMap[key]; !ok {
				srcKeysMap[key] = struct{}{}
				srcKeys = append(srcKeys, key)
			}
		}
	}
}
