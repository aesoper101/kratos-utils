package encoder

import (
	"github.com/aesoper101/kratos-utils/middleware/requestid"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http"
	http2 "net/http"
	"strings"
	"time"
)

const (
	baseContentType = "application"
)

// ContentType returns the content-type with base prefix.
func contentType(subtype string) string {
	return strings.Join([]string{baseContentType, subtype}, "/")
}

func ApiEncodeResponse() http.EncodeResponseFunc {
	return func(w http2.ResponseWriter, r *http2.Request, v interface{}) error {
		if v == nil {
			return nil
		}
		if rd, ok := v.(http.Redirector); ok {
			url, code := rd.Redirect()
			http2.Redirect(w, r, url, code)
			return nil
		}

		type response struct {
			// 错误码，跟 http-status 一致，并且在 grpc 中可以转换成 grpc-status
			Code int32 `json:"code"`
			// 错误原因，定义为业务判定错误码
			Reason string `json:"reason,omitempty"`
			// 错误信息，为用户可读的信息，可作为用户提示内容
			Message   string      `json:"message"`
			Data      interface{} `json:"data,omitempty"`
			RequestId string      `json:"requestId"`
			Time      int64       `json:"time"`
			// 错误元信息，为错误添加附加可扩展信息
			Metadata map[string]string `json:"metadata,omitempty"`
		}

		reply := &response{
			Code:    200,
			Time:    time.Now().Unix(),
			Message: "OK",
		}

		switch d := v.(type) {
		case error:
			err := errors.FromError(d)
			reply.Code = err.GetCode()
			reply.Message = err.GetMessage()
			reply.Reason = err.GetReason()
			if m := err.GetMetadata(); m != nil {
				reply.Metadata = m
			}
			break
		default:
			reply.Data = v
		}

		requestId := w.Header().Get(requestid.HeaderRequestId)
		if len(requestId) > 0 {
			reply.RequestId = requestId
		}

		codec, _ := http.CodecForRequest(r, "Accept")
		data, err := codec.Marshal(reply)
		if err != nil {
			return err
		}

		w.Header().Set("Content-Type", contentType(codec.Name()))
		_, err = w.Write(data)
		if err != nil {
			return err
		}
		return nil
	}
}