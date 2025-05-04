package response

import (
	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/go-kratos/kratos/v2/transport/http"
	"go.opentelemetry.io/otel/trace"
	"strings"
)

// Response 是统一的响应结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	TraceID string      `json:"trace_id"`
}

func EncoderResponse() http.EncodeResponseFunc {
	return func(w http.ResponseWriter, request *http.Request, i any) error {
		if i == nil {
			return nil
		}
		traceID := ""
		trace.SpanContextFromContext(request.Context())
		if id := w.Header().Get("X-Trace-ID"); id != "" {
			traceID = id
		}

		// 2. 如果响应头中没有，尝试从请求头获取
		if traceID == "" {
			if id := request.Header.Get("X-Trace-ID"); id != "" {
				traceID = id
			} else if id := request.Header.Get("uber-trace-id"); id != "" {
				parts := strings.Split(id, ":")
				if len(parts) > 0 {
					traceID = parts[0]
				}
			} else if id := request.Header.Get("X-B3-TraceId"); id != "" {
				traceID = id
			}
		}
		resp := &Response{
			Code:    200,
			Message: "success",
			Data:    i,
			TraceID: traceID,
		}
		codec := encoding.GetCodec("json")
		data, err := codec.Marshal(resp)
		if err != nil {
			return err
		}
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(data)
		if err != nil {
			return err
		}
		return nil
	}
}
