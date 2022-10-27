package mid

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ezzycreative1/svc-pokemon/pkg/mlog"
	"github.com/labstack/echo/v4"
)

type reqResTrap struct {
	logger mlog.Logger
}

func NewReqResMiddleware(logger mlog.Logger) *reqResTrap {
	return &reqResTrap{
		logger: logger,
	}
}

func (rr *reqResTrap) Middle(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		req := c.Request()
		res := c.Response()
		start := time.Now()

		// traceID
		traceID := GetID(c)

		// referer
		referer := getFirstValueFromHeader(req, "Referer")
		clientID := getFirstValueFromHeader(req, "X-Client-Id")
		// source ip
		ipAddress := getFirstValueFromHeader(req, "X-Forwarded-For")

		// LOG for incoming request
		rr.logger.Info("request started",
			mlog.String("referer", referer),
			mlog.String("client_id", clientID),
			mlog.String("source_ip", ipAddress),
			mlog.String("path", req.URL.Path),
			mlog.String("trace_id", traceID),
		)

		// LOG after request end
		defer func() {
			rr.logger.Info("request completed",
				mlog.String("referer", referer),
				mlog.String("client_id", clientID),
				mlog.String("source_ip", ipAddress),
				mlog.String("path", req.URL.Path),
				mlog.String("trace_id", traceID),
				mlog.Any("latency", fmt.Sprintf("%d ms", time.Since(start).Milliseconds())),
				mlog.Int("status", res.Status),
			)
		}()

		if err = next(c); err != nil {
			c.Error(err)
		}
		return
	}
}

func getFirstValueFromHeader(req *http.Request, key string) string {
	vs, ok := req.Header[key]
	if ok {
		if len(vs) != 0 {
			return vs[0]
		}
	}
	return ""
}
