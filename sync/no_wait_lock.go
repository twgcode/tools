/**
@Author: wei-g
@Date:   2021/7/20 7:39 下午
@Description:
*/

package sync

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis"
	"github.com/google/uuid"
)

var (
	ErrLockedState   = fmt.Errorf("the current lock is locked")
	ErrNotHolderLock = fmt.Errorf("it is not the holder can not unlock")
)

// NoWaitLock 基于redis 非等待锁
type NoWaitLock struct {
	rdb     *redis.ClusterClient
	key     string
	timeOut time.Duration
	mutex   sync.Mutex
	uuid    string
}

func NewNoWaitLock(client *redis.ClusterClient, key string, timeOut time.Duration) (locker *NoWaitLock, err error) {
	var (
		id uuid.UUID
	)
	locker = &NoWaitLock{
		rdb:     client,
		key:     key,
		timeOut: timeOut + 1,
		mutex:   sync.Mutex{},
	}
	if id, err = uuid.NewRandom(); err != nil {
		return
	}
	locker.uuid = id.String()
	return
}
func (n *NoWaitLock) Lock() (err error) {
	var (
		ok bool
	)
	defer n.mutex.Unlock()
	n.mutex.Lock()
	if ok, err = n.rdb.SetNX(n.key, n.uuid, n.timeOut).Result(); err != nil {
		return
	}
	if !ok { // 别的 worker 已经持有锁
		return fmt.Errorf("key: %s, %w", n.key, ErrLockedState)
	}
	return
}
func (n *NoWaitLock) UnLock() (err error) {
	var (
		id string
	)
	defer n.mutex.Unlock()
	n.mutex.Lock()
	// 如果 key不存在 可能就是超时造成,所以直接返回成功
	id, err = n.rdb.Get(n.key).Result()
	if err == redis.Nil {
		err = nil
		return
	}
	if id != n.uuid { // value 内容不相等, 代表别人正在持有锁
		return ErrNotHolderLock
	}
	err = n.rdb.Del(n.key).Err() // 进行删除
	return
}

// Do 对执行函数使用 NoWaitLock 进行同步, err1 返回 回调函数有关的错误, err2 返回 锁有关的错误,
func (n *NoWaitLock) Do(f func() error) (err1 error, err2 error) {
	if err2 = n.Lock(); err2 != nil {
		return
	}
	defer func() {
		err2 = n.UnLock()
	}()
	err1 = f()
	return
}
