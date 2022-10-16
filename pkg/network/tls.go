package network

import (
	"crypto/tls"
	"github.com/aesoper101/kratos-utils/protobuf/types/confpb"
)

func InitTLSConfig(cfg *confpb.TLSConfig) *tls.Config {
	if cfg != nil && cfg.GetEnabled() {
		cert, err := tls.LoadX509KeyPair(cfg.GetCertFile(), cfg.GetKeyFile())
		if err != nil {
			panic(err)
		}
		return &tls.Config{Certificates: []tls.Certificate{cert}}
	}

	return nil
}
