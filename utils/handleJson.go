package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/ninajika/crud-go/models"
)

func GetDump(id string) (*models.PostType, error) {
	result, err := ReadJson[models.PostType](fmt.Sprintf("dummies/%s/post.json", id))
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

func CreateJson[R any](id string, data *R) error {
	dirPath := fmt.Sprintf("dummies/%s", id)
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return err
	}

	filePath := filepath.Join(dirPath, "post.json")
	if _, err := os.Stat(filePath); err == nil {
		return fmt.Errorf("post.json already exists for ID: %s", id)
	}

	return WriteJson(filePath, data)
}

func UpdateJson[R any](id string, data *R) error {
	dirPath := fmt.Sprintf("dummies/%s", id)
	filePath := filepath.Join(dirPath, "post.json")

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf("post.json does not exist for ID: %s", id)
	}

	return WriteJson(filePath, data)
}

func WriteJson[R any](filePath string, data *R) error {
	jsonValue, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return err
	}

	err = os.WriteFile(filePath, jsonValue, 0644)
	if err != nil {
		fmt.Println("Error writing JSON file:", err)
		return err
	}

	return nil
}

func DeletePost(id string) error {
	log.Printf("Deleting post with ID: %s", id)

	dirPath := fmt.Sprintf("dummies/%s", id)
	err := os.RemoveAll(dirPath)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
