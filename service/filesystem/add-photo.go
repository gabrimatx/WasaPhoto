package filesystem

import (
	"encoding/base64"
	"fmt"
	"os"
)

func SaveBase64Photo(base64String, fileName string) error {
	decoded, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return fmt.Errorf("error decoding base64 string: %v", err)
	}

	file, err := os.Create("service/filesystem/" + fileName)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	_, err = file.Write(decoded)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	fmt.Printf("Photo saved successfully as %s\n", fileName)

	return nil
}
