package langTemplates

import (
	"fmt"
	"io/ioutil"
	"os"
)
/**
creates a template at cwd
*/
func CreateTemplate(options map[string]string) {


}

func writeToFile(filepath string, data string) error {
	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(data)
	if err != nil {
		return err
	}
	return nil
}