package validation

import (
	"testing"

	"github.com/nbio/st"
)

//
// TestFileIsTextFile - test if a file validates as a text file.
//
func TestFileIsTextFile(t *testing.T) {
	testFile1 := "../../data/customers.txt"
	testFile2 := "../../data/customers.json"

	test, err := FileIsTextFile(testFile1)
	st.Expect(t, err, nil)
	st.Expect(t, test, true)

	test, err = FileIsTextFile(testFile2)
	st.Expect(t, err, nil)
	st.Expect(t, test, false)

	test, err = FileIsTextFile("")
	st.Expect(t, err.Error(), "open : no such file or directory")
	st.Expect(t, test, false)
}
