package httpsign

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
)

const (
	headerXAppID     = "x-md-global-appId"
	headerXSignature = "x-md-global-signature"
	headerTimestamp  = "x-md-global-timestamp"
)

func defaultAppIdGetter(ctx context.Context) string {
	if tr, ok := transport.FromServerContext(ctx); ok {
		requestHeader := tr.RequestHeader()
		return requestHeader.Get(headerXAppID)
	}
	return ""
}

func defaultSignatureGetter(ctx context.Context) string {
	if tr, ok := transport.FromServerContext(ctx); ok {
		requestHeader := tr.RequestHeader()
		return requestHeader.Get(headerXSignature)
	}
	return ""
}

func defaultTimestampGetter(ctx context.Context) string {
	if tr, ok := transport.FromServerContext(ctx); ok {
		requestHeader := tr.RequestHeader()
		return requestHeader.Get(headerTimestamp)
	}
	return ""
}

func defaultSignatureValidator(ctx context.Context, appId, secret, signature, timestamp string) error {
	if req, ok := http.RequestFromServerContext(ctx); ok {
		stringToSign := fmt.Sprintf("%v\n%v\n%v\n", appId, req.URL.Path, timestamp)

		hash := hmac.New(sha256.New, []byte(secret))
		hash.Write([]byte(stringToSign))

		if signature != base64.StdEncoding.EncodeToString(hash.Sum(nil)) {
			return errors.Forbidden("ERROR_INCORRECT_SIGNATURE", "authentication failed")
		}
	}

	return nil
}
