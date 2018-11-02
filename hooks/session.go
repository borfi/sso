package hooks

import (
	"fmt"
	"sso/xengine/xconfig"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

const (
	sessionKey               = "secret" //安全key
	sessionMemoryType        = "memory" //默认基于内存
	sessionRedisType         = "redis"  //基于redis
	sessionDefaultSize       = 1        //默认连接池中自动连接个数
	sessionDefaultCookieName = "mysso"  //默认cookie名称
	sessionDefaultMaxAge     = 3600     //默认1小时
	sessionDefaultHTTPOnly   = true     //是否设置不让js取到cookie，默认是
	sessionDefaultSecure     = false    //是否仅https能取到cookie，默认否
)

// Session session
func Session() gin.HandlerFunc {
	//这里可以做redis、memstore可选项
	var store sessions.Store
	stype := getSessionType()

	if stype == sessionRedisType {
		host := getSessionHost()
		port := getSessionPort()
		auth := getSessionAuth()
		size := getSessionSize()

		var err error
		store, err = redis.NewStore(size, "tcp", fmt.Sprintf("%s:%d", host, port), auth, []byte(sessionKey))
		if err != nil {
			panic(err.Error())
		}
	} else {
		store = memstore.NewStore([]byte(sessionKey))
	}

	sessionDomain := getSessionDomain()
	sessionMaxAge := getSessionMaxAge()
	sessionSecure := getSessionSecure()
	sessionHTTPOnly := getSessionHTTPOnly()
	sessionCookieName := getSessionCookieName()

	store.Options(sessions.Options{
		Path:     "/",
		Domain:   sessionDomain,
		MaxAge:   sessionMaxAge, //1小时
		Secure:   sessionSecure,
		HttpOnly: sessionHTTPOnly,
	})
	return sessions.Sessions(sessionCookieName, store)
}

// get session domain
func getSessionDomain() string {
	domain, err := xconfig.Config().String("session", "domain")
	if err != nil {
		return domain
	}
	return domain
}

// get session cookie name
func getSessionCookieName() string {
	cookiename, err := xconfig.Config().String("session", "cookiename")
	if err != nil {
		return sessionDefaultCookieName
	}
	return cookiename
}

// get session httponly
func getSessionHTTPOnly() bool {
	httponly, err := xconfig.Config().Bool("session", "httponly")
	if err != nil {
		return sessionDefaultHTTPOnly
	}
	return httponly
}

// get session secure
func getSessionSecure() bool {
	secure, err := xconfig.Config().Bool("session", "secure")
	if err != nil {
		return sessionDefaultSecure
	}
	return secure
}

// get session max age
func getSessionMaxAge() int {
	maxage, err := xconfig.Config().Int("session", "maxage")
	if err != nil {
		return sessionDefaultMaxAge
	}
	return maxage
}

// get session type
func getSessionType() string {
	stype, err := xconfig.Config().String("session", "type")
	if err != nil {
		return sessionMemoryType
	}
	return stype
}

// get session host
func getSessionHost() string {
	host, err := xconfig.Config().String("session", "host")
	if err != nil {
		return host
	}
	return host
}

// get session port
func getSessionPort() int {
	port, err := xconfig.Config().Int("session", "port")
	if err != nil {
		return port
	}
	return port
}

// get session auth
func getSessionAuth() string {
	auth, err := xconfig.Config().String("session", "auth")
	if err != nil {
		return auth
	}
	return auth
}

// get session size
func getSessionSize() int {
	size, err := xconfig.Config().Int("session", "size")
	if err != nil {
		size = sessionDefaultSize
	}

	//连接池中的连接数
	if size <= 0 {
		size = sessionDefaultSize
	}
	return size
}
