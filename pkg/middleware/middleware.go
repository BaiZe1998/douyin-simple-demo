package middleware

import (
	"context"
	"fmt"
	"github.com/BaiZe1998/douyin-simple-demo/dto"
	"golang.org/x/time/rate"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/BaiZe1998/douyin-simple-demo/pkg/util"
	"github.com/gin-gonic/gin"
)

func whiteList() map[string]string {
	return map[string]string{
		"/douyin/feed/":           "GET",
		"/douyin/user/register/":  "POST",
		"/douyin/user/login/":     "POST",
		"/douyin/publish/action/": "POST",
	}
}

func withinWhiteList(url *url.URL, method string) bool {
	target := whiteList()
	queryUrl := strings.Split(fmt.Sprint(url), "?")[0]
	if _, ok := target[queryUrl]; ok {
		if target[queryUrl] == method {
			return true
		}
		return false
	}
	return false
}

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {

		type QueryToken struct {
			Token string `binding:"required" form:"token"`
		}

		// 当路由不在白名单内时进行token检测
		if !withinWhiteList(c.Request.URL, c.Request.Method) {
			var queryToken QueryToken
			if c.ShouldBindQuery(&queryToken) != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status_code": 40001,
					"status_msg":  "请先登录",
				})
				return
			}
			// 验证token有效性
			userClaims, err := util.ParseToken(queryToken.Token)
			if err != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status_code": 40001,
					"status_msg":  "登陆过期",
				})
				return
			}
			c.Set("token", userClaims)
			c.Set("user_id", userClaims.ID)
		}
		c.Next()
	}
}

func Limiter(r rate.Limit, b int, t time.Duration) gin.HandlerFunc {
	limiters := &sync.Map{}

	return func(c *gin.Context) {
		// 获取限速器
		key := c.ClientIP()
		l, _ := limiters.LoadOrStore(key, rate.NewLimiter(r, b))

		ctx, cancel := context.WithTimeout(c, t)
		defer cancel()

		if err := l.(*rate.Limiter).Wait(ctx); err != nil {
			dto.WriteLog("error", key+"请求过于频繁")
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": err})
		}
		c.Next()
	}
}
