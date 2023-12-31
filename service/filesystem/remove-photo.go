package filesystem

import (
	"fmt"
	"os"
)

func removePhoto(fileName string) error {
	// Step 1: Remove the file
	err := os.Remove(fileName)
	if err != nil {
		return fmt.Errorf("error removing file: %v", err)
	}

	fmt.Printf("Photo %s removed successfully\n", fileName)

	return nil
}
