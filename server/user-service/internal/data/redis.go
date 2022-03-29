package data

import (
	"context"
	"strconv"
	"time"
)

const lastCreateUserID = "lastCreateUserID"

func (ur *userRepo) GetLastCreateUserID(ctx context.Context) (id int64, err error) {
	ret := ur.data.rdb.Get(ctx, lastCreateUserID)
	return strconv.ParseInt(ret.Val(), 10, 64)
}

func (ur *userRepo) SetLastCreateUserID(ctx context.Context, id int64) (err error) {
	ret := ur.data.rdb.Set(ctx, lastCreateUserID, id, time.Millisecond*200)
	return ret.Err()
}
