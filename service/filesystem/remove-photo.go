package filesystem

import (
	"fmt"
	"os"
)

func RemovePhoto(idPhoto uint64) error {
	fileName := fmt.Sprintf("service/filesystem/%d.jpg", idPhoto)

	err := os.Remove(fileName)
	if err != nil {
		return fmt.Errorf("error removing file: %v", err)
	}

	fmt.Printf("Photo %s removed successfully\n", fileName)
	return nil
}
