package repo

import (
	"context"
	"fmt"
	"github.com/huage66/chat/interval/keys"
	"github.com/huage66/chat/interval/orm"
	"github.com/huage66/chat/interval/vars"
	"github.com/huage66/chat/model"
	"net"
)

// 登录, 如何用户ip不存在,直接默认新用户
func Login(chat model.ChatMessage, conn net.Conn) bool {
	// 该用户登录已经注册, 不需要走下面流程
	if _, ok := vars.UserMap.Load(chat.Ip); ok {
		return true
	}
	ctx := context.Background()
	var u model.User
	if err := orm.RedisClient.Get(ctx, GetUserKey(chat.Ip)).Scan(&u); err != nil {
		return false
	}
	if len(u.IP) < 1 {
		u.IP = chat.Ip
		if err := orm.RedisClient.Set(ctx, GetUserKey(u.IP), &u, 0).Err(); err != nil {
			return false
		}
	}

	info := model.UserInfo{
		User:   u,
		Conn:   conn,
		OnLine: true,
	}
	vars.UserMap.Store(info.IP, &info)
	return true
}

// 添加用户信息
func AddUserInfo(ctx context.Context, u model.User) error {
	if err := orm.RedisClient.Set(ctx, GetUserKey(u.IP), &u, 0).Err(); err != nil {
		return err
	}
	return nil
}

func GetUserKey(ip string) string {
	return fmt.Sprintf(keys.UserPrefix+"%s", ip)
}
