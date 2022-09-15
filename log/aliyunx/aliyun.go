package aliyunx

import (
	"github.com/aesoper101/kratos-utils/protobuf/types/confpb"
	"github.com/go-kratos/kratos/contrib/log/aliyun/v2"
)

func NewAliyunLoggerWithConfig(cfg *confpb.AliyunLog) aliyun.Logger {
	var opts []aliyun.Option
	if cfg != nil {
		opts = append(opts, aliyun.WithEndpoint(cfg.Endpoint),
			aliyun.WithAccessKey(cfg.AccessKey),
			aliyun.WithAccessSecret(cfg.AccessSecret),
			aliyun.WithProject(cfg.Project),
			aliyun.WithProject(cfg.Project))
	}
	return aliyun.NewAliyunLog(opts...)
}
