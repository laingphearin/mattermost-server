package model

import (
"encoding/json"
"io"
)

type Friend struct {
	UserId 				string   `json:"userid"`
	FriendId 			string   `json:"friendid"`
	CreateAt    		int64    `json:"create_at"`
	UpdateAt			int64	 `json:"update_at"`
}


// FriendFromJson will decode the input and return a Friend
func FriendFromJson(data io.Reader) *Friend {
	var friend *Friend
	json.NewDecoder(data).Decode(&friend)
	return friend
}
