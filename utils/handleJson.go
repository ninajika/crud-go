package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/ninajika/crud-go/controllers"
)

func GetDump(id string) (*controllers.PostType, error) {
	result, err := ReadJson[controllers.PostType](fmt.Sprintf("dummies/%s/post.json", id))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, nil
}

func ReadJson[R any](path string) (*R, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer jsonFile.Close()
	bytesValue, _ := io.ReadAll(jsonFile)
	var result *R
	json.Unmarshal(bytesValue, &result)
	return result, nil
}
