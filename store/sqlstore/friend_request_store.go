package sqlstore

import (
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/store"
)

type SqlFriendRequestStore struct {
	SqlStore
}

func (fs SqlFriendRequestStore) Save(friendRequest *model.FriendRequest) (*model.FriendRequest, *model.AppError) {
	fs.GetMaster().Insert(friendRequest)
	return friendRequest, nil
}

func (fs SqlFriendRequestStore) Update(friendRequest *model.FriendRequest) (*model.FriendRequest, *model.AppError) {
	fs.GetMaster().Update(friendRequest)
	return friendRequest, nil
}

func newSqlFriendRequestStore(sqlStore SqlStore) store.FriendRequestStore {
	s := &SqlFriendRequestStore{
		SqlStore: sqlStore,
	}

	for _, db := range sqlStore.GetAllConns() {
		db.AddTableWithName(model.FriendRequest{}, "FriendRequest").SetKeys(false, "UserId", "FriendId")
	}

	return s
}
