package shadowsocks

import (
	"github.com/sagernet/sing-box/option"
	"github.com/sagernet/sing/common"
)

func (h *MultiInbound) UpdateUsersByOptions(users []option.ShadowsocksUser) error {
	h.usersAccess.Lock()
	defer h.usersAccess.Unlock()
	err := h.service.UpdateUsersWithPasswords(common.MapIndexed(users, func(index int, user option.ShadowsocksUser) int {
		return index
	}), common.Map(users, func(user option.ShadowsocksUser) string {
		return user.Password
	}))
	if err != nil {
		return err
	}
	h.users = users
	return nil
}
