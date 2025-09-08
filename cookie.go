package kingdee

import (
	"context"
	"encoding/json"
	"github.com/imroc/req/v3"
	"github.com/muloop/kingdee/object"
	"github.com/redis/go-redis/v9"
	"net/http"
	"time"
)

type Cookie struct {
	debug   bool
	redis   *redis.Client
	PassKey string
	Key     string
	config  LoginConfig
}

func NewCookie(conf Config) *Cookie {

	c := redis.NewClient(&redis.Options{
		Addr:     conf.RedisConfig.Host,
		Password: conf.RedisConfig.Pass,
		DB:       0,
	})
	return &Cookie{
		debug:   conf.Debug,
		config:  conf.LoginConfig,
		Key:     conf.RedisConfig.CookieKey,
		PassKey: conf.RedisConfig.PassKey,
		redis:   c,
	}
}

func (c *Cookie) Cookie() (cookies []*http.Cookie) {
	var err error
	var cookieJson string
	cookieJson, err = c.redis.Get(context.Background(), c.Key).Result()
	if err == nil && cookieJson != "" {
		err = json.Unmarshal([]byte(cookieJson), &cookies)
	}
	if err != nil || len(cookies) == 0 {
		pass, _ := c.redis.Get(context.Background(), c.PassKey).Result()
		var request = object.LoginRequest{
			AcctID:   c.config.AcctID,
			Username: c.config.User,
			LcId:     c.config.LcId,
			Password: pass,
		}
		cli := req.C().SetBaseURL(c.config.Host).SetCommonContentType("application/json")
		if c.debug {
			cli = cli.DevMode()
		}
		resp, _ := cli.R().
			SetBody(request).
			Post(LOGIN_API)

		if resp.IsSuccessState() {
			cookies = resp.Cookies()
		}
		b, _ := json.Marshal(cookies)
		c.redis.Set(context.Background(), c.Key, string(b), time.Minute*19)
	}
	return cookies
}
