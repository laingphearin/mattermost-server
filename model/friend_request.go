package model

import (
	"encoding/json"
	"io"
)

type FriendRequest struct {
	UserId              string                  `json:"userid"`
	FriendId            string                  `json:"friendid"`
	FriendRequestStatus int 					`json:"friendrequeststatus"`

}

// FriendRequestFromJson will decode the input and return a FriendRequest
func FriendRequestFromJson(data io.Reader) *FriendRequest {
	var friendRequest *FriendRequest
	json.NewDecoder(data).Decode(&friendRequest)
	return friendRequest
}
