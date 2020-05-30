package model

type FriendRequestStatusEnum struct {
	Pending int8
	Refuse int
	Agree int
	Ignore int
}

var FriendRequestStatus = &FriendRequestStatusEnum{
	Pending: 0,
	Refuse: 1,
	Agree: 2,
	Ignore: 3,
}
