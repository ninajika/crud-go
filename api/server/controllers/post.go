package controllers

import (
	"fmt"
	"strconv"

	"github.com/ninajika/crud-go/api/server/types"
	"github.com/ninajika/crud-go/api/utils"
)

func GetPost(id string) (interface{}, error) {
	result, err := utils.ReadJson[types.PostType](fmt.Sprintf("test/dummies/%s/post.json", id))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, nil
}

func CreatePost(id int64, title string, body string, tags []string) error {
	data := &types.PostType{
		ID:    id,
		Title: title,
		Body:  body,
		Tags:  tags,
		Reactions: types.PostReaction{
			Likes:   0,
			Disikes: 0,
		},
		Views:  0,
		UserId: 0,
	}
	if err := utils.WriteJson(fmt.Sprintf("test/dummies/%d/post.json", id), data); err != nil {
		return err
	}
	return nil
}

func RemovePost(id string) bool {
	if err := utils.DeletePost(id); err != nil {
		return false
	}
	return true
}

func UpdatePost(id string, title string, body string, tags []string) (bool, error) {
	data, err := GetPost(id)
	if err != nil {
		return false, err
	}
	parsedId, _ := strconv.ParseInt(id, 10, 32)
	ok := RemovePost(id)
	if !ok {
		return false, fmt.Errorf("fail to update post: %w", err)
	}
	newData := &types.PostType{
		ID:        parsedId,
		Title:     title,
		Body:      body,
		Tags:      tags,
		Reactions: data.(types.PostType).Reactions,
		Views:     data.(types.PostType).Views,
		UserId:    data.(types.PostType).UserId,
	}
	writtenErr := utils.WriteJson(fmt.Sprintf("test/dummies/%s/post.json", id), newData)
	if writtenErr != nil {
		return false, writtenErr
	}
	return true, nil
}
