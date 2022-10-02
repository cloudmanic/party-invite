package validation

import (
	"net/http"
	"os"
	"path/filepath"
)

//
// FileIsTextFile will take the input of a full path to a file.
// It will read that file and validate that is is a text file.
//
func FileIsTextFile(filePath string) (bool, error) {
	// Read the content to get the type
	buf, err := os.ReadFile(filePath)

	if err != nil {
		return false, err
	}

	filetype := http.DetectContentType(buf)

	if filetype != "text/plain; charset=utf-8" {
		return false, nil
	}

	// Verify the extension to make sure it is .txt
	extension := filepath.Ext(filePath)

	if extension == ".txt" {
		return true, nil
	}

	return false, nil
}
