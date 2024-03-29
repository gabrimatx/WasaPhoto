package filesystem

import (
	"fmt"
	"os"
)

func RemovePhoto(idPhoto uint64) error {
	fileName := fmt.Sprintf("/tmp/filesystem/%d.jpg", idPhoto)

	err := os.Remove(fileName)
	if err != nil {
		return err
	}

	return nil
}
