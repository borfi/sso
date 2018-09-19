package hooks

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
)

// Session session
func Session() gin.HandlerFunc {
	//以后这里可以做redis、memstore可选项
	store := memstore.NewStore([]byte("secret"))
	store.Options(sessions.Options{
		Path:     "/",
		Domain:   "",
		MaxAge:   3600, //1小时
		Secure:   false,
		HttpOnly: true,
	})
	return sessions.Sessions("mysession", store)
}
