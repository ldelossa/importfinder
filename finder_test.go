package importfinder_test

import (
	"fmt"
	"os"
	"testing"

	i "github.com/ldelossa/importfinder"
)

// test table for TestFindImports
var FindImportsTT = []struct {
	Dir string
	Out string
}{
	{"./testdata", `{"main.go":["\"encoding/json\"","\"fmt\"","\"math\""],"testpackage/packagefile.go":["\"encoding/json\"","\"fmt\"","\"sync\""]}`},
}

func TestFindImports(t *testing.T) {
	for _, tt := range FindImportsTT {
		// change dir test table's dir
		err := os.Chdir(tt.Dir)
		if err != nil {
			t.Fatalf("failed to chdir to provide test table directory: %v", err)
		}

		f := i.NewFinder()
		err = f.FindImports()
		if err != nil {
			t.Fatalf("failed to find imports: %v", err)
		}

		s := fmt.Sprint(f)
		if s != tt.Out {
			t.Fatalf("import finder ouput: \n%s\n expected test table output: \n%s\n", s, tt.Out)
		}

	}
}
