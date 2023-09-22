package pkg

import "os"

func CreateOrReplaceDir(path string) error {
	// Step 1: Check if directory exists
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		// Step 2: Remove the directory
		if err := os.RemoveAll(path); err != nil {
			return err
		}
	}

	// Step 3: Create the directory
	return os.Mkdir(path, 0755)
}
