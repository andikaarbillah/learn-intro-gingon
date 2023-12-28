package middlewares

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupLogFile(){
	f, _ :=  os.Create("LogFile.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

}

func Logger()gin.HandlerFunc{
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] %s %s %d %s\n",
			params.ClientIP,
			params.TimeStamp.Format(time.RFC822),
			params.Method,
			params.Path,
			params.StatusCode,
			params.Latency,

	)
	})
}