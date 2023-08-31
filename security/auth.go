package security

import (
	"myadmin/helper"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader("Authorization")

		if authorizationHeader == "" {
			resp := helper.ErrorJSON(ctx, "Missing Header", http.StatusUnauthorized, nil)

			ctx.JSON(http.StatusUnauthorized, resp)

			ctx.Abort()

			return
		}

		if !strings.Contains(authorizationHeader, "Bearer") {
			resp := helper.ErrorJSON(ctx, "Invalid token", http.StatusUnauthorized, nil)

			ctx.JSON(http.StatusUnauthorized, resp)

			ctx.Abort()

			return
		}

		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
		err := ValidateToken(tokenString)

		if err != nil {
			resp := helper.ErrorJSON(ctx, err.Error(), http.StatusUnauthorized, nil)

			ctx.JSON(http.StatusUnauthorized, resp)

			ctx.Abort()

			return
		}

		ctx.Next()
	}
}
