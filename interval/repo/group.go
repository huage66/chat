package repo

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/huage66/chat/interval/keys"
	"github.com/huage66/chat/interval/orm"
	"github.com/huage66/chat/interval/utils"
	"github.com/huage66/chat/interval/vars"
)

// 加载group群聊到内存
func InitGroup() {
	ctx := context.Background()
	keyList, err := ScanKey(ctx, keys.GroupPrefix+":*")
	if err != nil {
		panic(err)
		return
	}
	pipeLine := orm.RedisClient.Pipeline()
	for _, k := range keyList {
		pipeLine.Get(ctx, k)
	}
	exec, err := pipeLine.Exec(ctx)
	if err != nil && err != redis.Nil {
		fmt.Println(err)
		return
	}

	groupList := make([]string, len(keyList))
	for _, cmder := range exec {
		cmd, ok := cmder.(*redis.StringCmd)
		if !ok {
			continue
		}
		fmt.Println(cmd.Val())
		groupList = append(groupList, cmd.Val())
	}
	for _, group := range groupList {
		maps := make(map[string]bool)
		vars.ChatMap.Store(group, maps)
	}
}

func ScanKey(ctx context.Context, match string) ([]string, error) {
	var (
		groupList []string
		keyMap    = make(map[string]bool)
	)

	for {
		scan := orm.RedisClient.Scan(ctx, 0, match, 10000)
		if scan.Err() != nil {
			fmt.Println(scan.Err())
			return groupList, scan.Err()
		}
		keys, cursor := scan.Val()
		for _, k := range keys {
			keyMap[k] = true
		}
		if cursor == 0 {
			break
		}
	}
	return utils.Map2Strings(keyMap), nil
}
