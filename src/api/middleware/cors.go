package middleware
import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/cors"
)

func InitCORSMiddleware() app.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"Content-Type"},
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE"},
	})
}
