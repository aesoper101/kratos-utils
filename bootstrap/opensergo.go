package bootstrap

import (
	"context"
	"github.com/aesoper101/kratos-utils/protobuf/types/confpb"
	"github.com/go-kratos/kratos/contrib/opensergo/v2"
	"github.com/go-kratos/kratos/v2"
)

func InitOpenSergo(app *kratos.App, cfg *confpb.OpenSergo) error {
	if cfg != nil && len(cfg.Endpoint) > 0 {
		osg, err := opensergo.New(opensergo.WithEndpoint(cfg.Endpoint))
		if err != nil {
			return err
		}
		if err = osg.ReportMetadata(context.Background(), app); err != nil {
			return err
		}
	}
	return nil
}
