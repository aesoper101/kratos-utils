package bootstrap

import (
	"fmt"
	"github.com/aesoper101/go-utils/str"
	consulKratos "github.com/go-kratos/kratos/contrib/config/consul/v2"
	etcdKratos "github.com/go-kratos/kratos/contrib/config/etcd/v2"
	nacosKratos "github.com/go-kratos/kratos/contrib/config/nacos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/env"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/hashicorp/consul/api"
	nacosClients "github.com/nacos-group/nacos-sdk-go/clients"
	nacosConstant "github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	etcdV3 "go.etcd.io/etcd/client/v3"
	GRPC "google.golang.org/grpc"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type ConfigProviderData struct {
	ConfigPath  string
	ConfigType  string
	ConfigHost  []string
	ServiceName string
	Env         string
	Username    string
	Password    string
}

// getConfigKey 获取合法的配置名
func getConfigKey(configKey string, useBackslash bool) string {
	if useBackslash {
		return strings.Replace(configKey, `.`, `/`, -1)
	} else {
		return configKey
	}
}

// NewRemoteConfigSource 创建一个远程配置源
func NewRemoteConfigSource(cfg *ConfigProviderData) config.Source {
	switch cfg.ConfigType {
	case ConfigNacos:
		return NewNacosConfigSource(cfg)
	case ConfigConsul:
		return NewConsulConfigSource(cfg)
	case ConfigEtcd:
		return NewEtcdConfigSource(cfg)
	}
	return nil
}

// NewNacosConfigSource 创建一个远程配置源 - Nacos
func NewNacosConfigSource(cfg *ConfigProviderData) config.Source {
	var sc []nacosConstant.ServerConfig

	for _, configHost := range cfg.ConfigHost {
		uri, _ := url.Parse(configHost)
		h := strings.Split(uri.Host, ":")
		addr := h[0]
		port, _ := strconv.Atoi(h[1])

		sc = append(sc, *nacosConstant.NewServerConfig(addr, uint64(port)))
	}

	cc := nacosConstant.ClientConfig{
		TimeoutMs:            10 * 1000, // http请求超时时间，单位毫秒
		BeatInterval:         5 * 1000,  // 心跳间隔时间，单位毫秒
		UpdateThreadNum:      20,        // 更新服务的线程数
		LogLevel:             "debug",
		CacheDir:             "../../configs/cache", // 缓存目录
		LogDir:               "../../configs/log",   // 日志目录
		NotLoadCacheAtStart:  true,                  // 在启动时不读取本地缓存数据，true--不读取，false--读取
		UpdateCacheWhenEmpty: true,                  // 当服务列表为空时是否更新本地缓存，true--更新,false--不更新
		Username:             cfg.Username,
		Password:             cfg.Password,
	}

	nacosClient, err := nacosClients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		panic(err)
	}

	return nacosKratos.NewConfigSource(nacosClient,
		nacosKratos.WithGroup(getConfigKey(cfg.ServiceName, false)),
		nacosKratos.WithDataID("bootstrap.yaml"),
	)
}

// NewEtcdConfigSource 创建一个远程配置源 - Etcd
func NewEtcdConfigSource(cfg *ConfigProviderData) config.Source {
	if len(cfg.ConfigHost) == 0 {
		panic("etcd hosts must be set.")
	}
	etcdClient, err := etcdV3.New(etcdV3.Config{
		Endpoints:   cfg.ConfigHost,
		Username:    cfg.Username,
		Password:    cfg.Password,
		DialTimeout: time.Second, DialOptions: []GRPC.DialOption{GRPC.WithBlock()},
	})
	if err != nil {
		panic(err)
	}

	etcdSource, err := etcdKratos.New(etcdClient, etcdKratos.WithPath(getConfigKey(cfg.ServiceName, true)))
	if err != nil {
		panic(err)
	}

	return etcdSource
}

// NewConsulConfigSource 创建一个远程配置源 - Consul
func NewConsulConfigSource(cfg *ConfigProviderData) config.Source {
	if len(cfg.ConfigHost) == 0 {
		panic("Consul hosts must be set.")
	}

	clientConfig := &api.Config{
		Address: cfg.ConfigHost[0],
	}
	if len(cfg.Password) > 0 && len(cfg.Username) > 0 {
		clientConfig.HttpAuth = &api.HttpBasicAuth{
			Username: cfg.Username,
			Password: cfg.Password,
		}
	}

	consulClient, err := api.NewClient(clientConfig)
	if err != nil {
		panic(err)
	}

	consulSource, err := consulKratos.New(consulClient,
		consulKratos.WithPath(getConfigKey(cfg.ServiceName, true)),
	)
	if err != nil {
		panic(err)
	}

	return consulSource
}

// NewFileConfigSource 创建一个本地文件配置源
func NewFileConfigSource(filePath string) config.Source {
	return file.NewSource(filePath)
}

// NewConfigProvider 创建一个配置
func NewConfigProvider(cfg ConfigProviderData) config.Config {
	envSourcePrefix := "KRATOS_"
	getPrefix := os.Getenv("ENV_SOURCE_PREFIX")
	if getPrefix != "" {
		envSourcePrefix = getPrefix
	}
	options := []config.Source{
		env.NewSource(envSourcePrefix),
		NewFileConfigSource(cfg.ConfigPath),
	}

	configType := strings.ToLower(strings.TrimSpace(cfg.ConfigType))
	if len(configType) != 0 {
		if isSupportConfigProvider(configType) {
			options = append(options, NewRemoteConfigSource(&cfg))
		} else {
			panic(fmt.Errorf("not support remote config provider %s", configType))
		}
	}

	return config.New(
		config.WithSource(options...),
	)
}

const (
	ConfigEtcd   = "etcd"
	ConfigConsul = "consul"
	ConfigNacos  = "nacos"
)

func isSupportConfigProvider(provider string) bool {
	return str.InArray(strings.ToLower(provider), []string{ConfigEtcd, ConfigNacos, ConfigNacos})
}
