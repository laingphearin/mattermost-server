package model

import (
	"encoding/json"
	"io"
)

type FriendRequest struct {
	UserId              string   `json:"userid"`
	FriendId            string   `json:"friendid"`
	FriendRequestStatus string 	 `json:"friendrequeststatus"`
	CreateAt    		int64    `json:"create_at"`
	UpdateAt			int64	 `json:"update_at"`
}

// FriendRequestFromJson will decode the input and return a FriendRequest
func FriendRequestFromJson(data io.Reader) *FriendRequest {
	var friendRequest *FriendRequest
	json.NewDecoder(data).Decode(&friendRequest)
	return friendRequest
}
