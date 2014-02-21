package percheron

import (
	"os"
)

// Checks if the given path exists
func DoesDirExist(folderPath string) (bool, error) {
	// check if dir exists
	_, err := os.Stat(folderPath)

	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	// no error means the stat returned successfully
	return true, nil
}
