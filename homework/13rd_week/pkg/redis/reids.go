package redis

import (
	"context"
	"time"

	redis "github.com/go-redis/redis/v8"
)

type RedisClient = redis.Client

func NewRedisClient(cnf *RedisConfig) (*RedisClient, error) {
	rdb := redis.NewClient(&redis.Options{
		//连接信息
		Network:  cnf.Network,  //网络类型，tcp or unix，默认tcp
		Addr:     cnf.Addr,     //主机名+冒号+端口，默认localhost:6379
		Password: cnf.Password, //密码
		DB:       cnf.DB,       // redis数据库index

		//连接池容量及闲置连接数量
		PoolSize:     cnf.PoolSize,     // 连接池最大socket连接数，默认为4倍CPU数， 4 * runtime.NumCPU
		MinIdleConns: cnf.MinIdleConns, //在启动阶段创建指定数量的Idle连接，并长期维持idle状态的连接数不少于指定数量；。

		//超时
		DialTimeout:  cnf.DialTimeout * time.Second,  //连接建立超时时间，默认5秒。
		ReadTimeout:  cnf.ReadTimeout * time.Second,  //读超时，默认3秒， -1表示取消读超时
		WriteTimeout: cnf.WriteTimeout * time.Second, //写超时，默认等于读超时
		PoolTimeout:  cnf.PoolTimeout * time.Second,  //当所有连接都处在繁忙状态时，客户端等待可用连接的最大等待时长，默认为读超时+1秒。

		//闲置连接检查包括IdleTimeout，MaxConnAge
		IdleCheckFrequency: cnf.IdleCheckFrequency * time.Second, //闲置连接检查的周期，默认为1分钟，-1表示不做周期性检查，只在客户端获取连接时对闲置连接进行处理。
		IdleTimeout:        cnf.IdleTimeout * time.Minute,        //闲置超时，默认5分钟，-1表示取消闲置超时检查
		MaxConnAge:         cnf.MaxConnAge * time.Second,         //连接存活时长，从创建开始计时，超过指定时长则关闭连接，默认为0，即不关闭存活时长较长的连接

		//命令执行失败时的重试策略
		MaxRetries:      cnf.MaxRetries,                         // 命令执行失败时，最多重试多少次，默认为0即不重试
		MinRetryBackoff: cnf.MinRetryBackoff * time.Millisecond, //每次计算重试间隔时间的下限，默认8毫秒，-1表示取消间隔
		MaxRetryBackoff: cnf.MaxRetryBackoff * time.Millisecond, //每次计算重试间隔时间的上限，默认512毫秒，-1表示取消间隔

		//可自定义连接函数

		//钩子函数
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, err
	} else {
		return rdb, nil
	}
}
