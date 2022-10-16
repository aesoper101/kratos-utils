package cache

import (
	"context"
	"github.com/aesoper101/kratos-utils/protobuf/types/confpb"
	"github.com/go-redis/redis/extra/redisotel/v9"
	"github.com/go-redis/redis/v9"
)

func NewRedisClient(cfg *confpb.Redis, options ...redisotel.TracingOption) (*redis.Client, func(), error) {
	opts := &redis.Options{
		Addr:     cfg.GetAddr(),
		Password: cfg.GetPassword(),
	}

	if cfg.Network != nil {
		opts.Network = cfg.GetNetwork()
	}

	if cfg.Username != nil {
		opts.Username = cfg.GetUsername()
	}

	if cfg.Db != nil {
		opts.DB = int(cfg.GetDb())
	}

	if dialTimeout := cfg.GetDialTimeout(); dialTimeout != nil {
		opts.DialTimeout = dialTimeout.AsDuration()
	}

	if writeTimeout := cfg.GetWriteTimeout(); writeTimeout != nil {
		opts.WriteTimeout = writeTimeout.AsDuration()
	}

	if readTimeout := cfg.GetReadTimeout(); readTimeout != nil {
		opts.ReadTimeout = readTimeout.AsDuration()
	}

	if cfg.PoolSize != nil {
		opts.PoolSize = int(cfg.GetPoolSize())
	}

	if poolTimeout := cfg.GetPoolTimeout(); poolTimeout != nil {
		opts.PoolTimeout = poolTimeout.AsDuration()
	}

	rdb := redis.NewClient(opts)
	err := redisotel.InstrumentTracing(rdb, options...)
	if err != nil {
		return nil, nil, err
	}

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, nil, err
	}

	return rdb, func() {
		_ = rdb.Close()
	}, nil
}

func NewClusterClient(cfg *confpb.RedisCluster) (*redis.ClusterClient, func(), error) {
	opts := &redis.ClusterOptions{
		Addrs:    cfg.GetAddr(),
		Password: cfg.GetPassword(),
	}
	rdb := redis.NewClusterClient(opts)

	if cfg.Username != nil {
		opts.Username = cfg.GetUsername()
	}

	if dialTimeout := cfg.GetDialTimeout(); dialTimeout != nil {
		opts.DialTimeout = dialTimeout.AsDuration()
	}

	if writeTimeout := cfg.GetWriteTimeout(); writeTimeout != nil {
		opts.WriteTimeout = writeTimeout.AsDuration()
	}

	if readTimeout := cfg.GetReadTimeout(); readTimeout != nil {
		opts.ReadTimeout = readTimeout.AsDuration()
	}

	if cfg.PoolSize != nil {
		opts.PoolSize = int(cfg.GetPoolSize())
	}

	if poolTimeout := cfg.GetPoolTimeout(); poolTimeout != nil {
		opts.PoolTimeout = poolTimeout.AsDuration()
	}

	if err := rdb.ForEachShard(context.Background(), func(ctx context.Context, shard *redis.Client) error {
		return shard.Ping(ctx).Err()
	}); err != nil {
		return nil, nil, err
	}

	return rdb, func() {
		_ = rdb.Close()
	}, nil
}
