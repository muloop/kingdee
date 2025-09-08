package kingdee

import (
	"errors"
)

type Config struct {
	LoginConfig LoginConfig
	RedisConfig RedisConfig
}
type LoginConfig struct {
	Host   string
	AcctID string
	User   string
	Pass   string
	LcId   int64
}

func (c *Config) Validate() error {
	if c.LoginConfig.Host == "" {
		return errors.New("金蝶后台地址不能为空")
	}
	if c.LoginConfig.AcctID == "" {
		return errors.New("acctId不能为空")
	}
	if c.LoginConfig.User == "" {
		return errors.New("用户名不能为空")
	}
	return nil
}

type RedisConfig struct {
	Host      string
	Pass      string
	CookieKey string
	PassKey   string
}
