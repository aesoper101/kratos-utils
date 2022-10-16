package encoder

import (
	"github.com/aesoper101/kratos-utils/pkg/middleware/requestid"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http"
	stbhttp "net/http"
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

func ApiEncodeResponse() http.EncodeResponseFunc {
	return func(w stbhttp.ResponseWriter, r *stbhttp.Request, v interface{}) error {
		if v == nil {
			return nil
		}
		if rd, ok := v.(http.Redirector); ok {
			url, code := rd.Redirect()
			stbhttp.Redirect(w, r, url, code)
			return nil
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

func ApiErrorEncoder() http.EncodeErrorFunc {
	return func(w http.ResponseWriter, r *http.Request, err error) {
		se := errors.FromError(err)

		reply := &response{
			Code:    se.GetCode(),
			Time:    time.Now().Unix(),
			Message: se.GetMessage(),
			Reason:  se.GetReason(),
		}

		if m := se.GetMetadata(); m != nil {
			reply.Metadata = m
		}

		requestId := w.Header().Get(requestid.HeaderRequestId)
		if len(requestId) > 0 {
			reply.RequestId = requestId
		}

		codec, _ := http.CodecForRequest(r, "Accept")
		body, err := codec.Marshal(reply)
		if err != nil {
			w.WriteHeader(stbhttp.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", contentType(codec.Name()))
		w.WriteHeader(int(se.Code))
		_, _ = w.Write(body)
	}
}
