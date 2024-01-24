package file

import (
    "fmt"
    "io/ioutil"
)

func ReadFile(filePath string) ([]byte, error) {
    content, err := ioutil.ReadFile(filePath)
    if err != nil {
        return nil, fmt.Errorf("error reading file: %v", err)
    }
    return content, nil
}

func SaveToFile(filePath string, data []byte) error {
    outputFilePath := filePath + "_output.txt"
    if err := ioutil.WriteFile(outputFilePath, data, 0644); err != nil {
        return fmt.Errorf("error writing output file: %v", err)
    }
    return nil
}
