package model

type FriendRequestStatusEnum struct {
	PENDING string
	REFUSE string
	AGREE string
	IGNORE string
}

var FriendRequestStatus = &FriendRequestStatusEnum {
	PENDING: "PENDING",
	REFUSE: "REFUSE",
	AGREE: "AGREE",
	IGNORE: "IGNORE",
}
