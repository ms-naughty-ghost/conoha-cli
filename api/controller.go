package api

import (
	"github.com/ms-naughty-ghost/conoha-cli/helper"
)

func (c *ApiContext) Excute(user *helper.User, req map[string]interface{}) error {

	switch c.Endpoints {
	case ArgIdentity:
		return c.ExcuteIdentity(user)
	case ArgDns:
		return c.ExcuteDns(user, req)
	default:
	}
	return nil
}
