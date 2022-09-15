package bootstrap

import (
	"github.com/aesoper101/kratos-utils/protobuf/types/confpb"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/contrib/registry/discovery/v2"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/contrib/registry/polaris/v2"
	"github.com/go-kratos/kratos/contrib/registry/zookeeper/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-zookeeper/zk"
	"github.com/hashicorp/consul/api"
	"github.com/nacos-group/nacos-sdk-go/clients"
	nacosConstant "github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	polarisConfig "github.com/polarismesh/polaris-go/pkg/config"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func NewDiscoveryProvider(cfg *confpb.Discovery) registry.Discovery {
	if cfg.GetEtcd() != nil {
		return NewEtcdDiscovery(cfg.GetEtcd())
	} else if cfg.GetConsul() != nil {
		return NewConsulDiscovery(cfg.GetConsul())
	} else if cfg.GetDiscovery() != nil {
		return NewDiscoveryDiscovery(cfg.GetDiscovery())
	} else if cfg.GetNacos() != nil {
		return NewNacosDiscovery(cfg.GetNacos())
	} else if cfg.GetPolaris() != nil {
		return NewPolarisDiscovery(cfg.GetPolaris())
	} else if cfg.GetZookeeper() != nil {
		return NewZookeeperDiscovery(cfg.GetZookeeper())
	}

	return nil
}

func NewEtcdDiscovery(cfg *confpb.EtcdConfig) registry.Discovery {
	if cfg == nil {
		panic("etcd discovery config must be set.")
	}
	config := clientv3.Config{
		Endpoints: cfg.Endpoints,
	}
	if cfg.Username != nil {
		config.Username = *cfg.Username
	}
	if cfg.Password != nil {
		config.Password = *cfg.Password
	}

	client, err := clientv3.New(config)
	if err != nil {
		panic(err)
	}

	return etcd.New(client)
}

func NewConsulDiscovery(cfg *confpb.ConsulConfig) registry.Discovery {
	if cfg == nil {
		panic("consul discovery config must be set.")
	}

	if len(cfg.GetEndpoints()) == 0 {
		panic("consul endpoints config must be set more than one.")
	}

	clientConfig := &api.Config{
		Address: cfg.GetEndpoints()[0],
	}
	if cfg.Password != nil && cfg.Username != nil {
		clientConfig.HttpAuth = &api.HttpBasicAuth{
			Username: cfg.GetUsername(),
			Password: cfg.GetPassword(),
		}
	}

	client, err := api.NewClient(clientConfig)
	if err != nil {
		panic(err)
	}

	return consul.New(client)
}

func NewDiscoveryDiscovery(cfg *confpb.DiscoveryConfig) registry.Discovery {
	if cfg == nil {
		panic("discovery discovery config must be set.")
	}
	return discovery.New(&discovery.Config{
		Nodes:  cfg.GetEndpoints(),
		Env:    cfg.GetEnv(),
		Region: cfg.GetRegion(),
		Zone:   cfg.GetZone(),
		Host:   cfg.GetHost(),
	})
}

func NewNacosDiscovery(cfg *confpb.NacosConfig) registry.Discovery {
	if cfg == nil {
		panic("nacos discovery config must be set.")
	}

	var sc []nacosConstant.ServerConfig

	for _, configHost := range cfg.GetEndpoints() {
		uri, _ := url.Parse(configHost)
		h := strings.Split(uri.Host, ":")
		addr := h[0]
		port, _ := strconv.Atoi(h[1])

		sc = append(sc, *nacosConstant.NewServerConfig(addr, uint64(port)))
	}

	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ServerConfigs: sc,
		},
	)

	if err != nil {
		log.Panic(err)
	}

	return nacos.New(client)
}

func NewPolarisDiscovery(cfg *confpb.PolarisConfig) registry.Discovery {
	if cfg == nil {
		panic("polaris discovery config must be set.")
	}

	c := polarisConfig.NewDefaultConfiguration(cfg.Endpoints)
	if err := c.Verify(); err != nil {
		panic(err)
	}

	var opts []polaris.Option
	if cfg.Namespace != nil {
		opts = append(opts, polaris.WithNamespace(cfg.GetNamespace()))
	}
	if cfg.RetryCount != nil {
		opts = append(opts, polaris.WithRetryCount(int(cfg.GetRetryCount())))
	}
	if cfg.Weight != nil {
		opts = append(opts, polaris.WithWeight(int(cfg.GetWeight())))
	}
	if cfg.Timeout != nil {
		opts = append(opts, polaris.WithTimeout(cfg.GetTimeout().AsDuration()))
	}
	if cfg.ServiceToken != nil {
		opts = append(opts, polaris.WithServiceToken(cfg.GetServiceToken()))
	}
	if cfg.Ttl != nil {
		opts = append(opts, polaris.WithTTL(int(cfg.GetTtl())))
	}
	if cfg.Healthy != nil {
		opts = append(opts, polaris.WithHealthy(cfg.GetHealthy()))
	}
	if cfg.Heartbeat != nil {
		opts = append(opts, polaris.WithHeartbeat(cfg.GetHeartbeat()))
	}
	if cfg.Isolate != nil {
		opts = append(opts, polaris.WithIsolate(cfg.GetIsolate()))
	}
	if cfg.Protocol != nil {
		opts = append(opts, polaris.WithProtocol(cfg.GetProtocol()))
	}

	return polaris.NewRegistryWithConfig(c, opts...)
}

func NewZookeeperDiscovery(cfg *confpb.ZookeeperConfig) registry.Discovery {
	if cfg == nil {
		panic("zookeeper discovery config must be set.")
	}

	timeout := time.Second * 15
	if cfg.Timeout != nil {
		timeout = cfg.GetTimeout().AsDuration()
	}

	conn, _, err := zk.Connect(cfg.GetEndpoints(), timeout)
	if err != nil {
		panic(err)
	}

	var opts []zookeeper.Option

	if cfg.Password != nil && cfg.Username != nil {
		opts = append(opts, zookeeper.WithDigestACL(cfg.GetUsername(), cfg.GetPassword()))
	}
	if cfg.Namespace != nil {
		opts = append(opts, zookeeper.WithRootPath(cfg.GetNamespace()))
	}

	return zookeeper.New(conn, opts...)
}
