package data

import (
    "context"
    v1 "github.com/BingguWang/bingfood-client-micro/api/prod/service/v1/pbgo/v1"
    v13 "github.com/BingguWang/bingfood-client-micro/api/nsqSrv/service/v1/pbgo/v1"
    "github.com/bwmarrin/snowflake"
    "github.com/go-kratos/kratos/v2/middleware/auth/jwt"
    "github.com/go-kratos/kratos/v2/middleware/recovery"
    "github.com/go-kratos/kratos/v2/registry"
    "github.com/go-kratos/kratos/v2/transport/grpc"
    "github.com/go-redis/redis/v8"
    jwt2 "github.com/golang-jwt/jwt/v4"
    "github.com/google/wire"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    lg "log"
    "os"
    "time"

    "github.com/BingguWang/bingfood-client-micro/app/order/service/internal/conf"
    "github.com/go-kratos/kratos/v2/log"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
    NewData,
    NewDB,
    NewOrderRepo,
    NewRedis,
    NewSnowFlakeNode,
    NewNsqServiceClient,
    NewProdServiceClient,
)

// Data .
type Data struct {
    // TODO wrapped database client
    db       *gorm.DB
    redisCli *redis.Client
    node     *snowflake.Node
}

// NewData .
func NewData(c *conf.Data, rd *redis.Client, node *snowflake.Node, logger log.Logger, db *gorm.DB) (*Data, func(), error) {
    cleanup := func() {
        log.NewHelper(logger).Info("closing the data resources")
    }
    return &Data{db: db, redisCli: rd, node: node}, cleanup, nil
}

// *conf.Data就是数据的配置结构体
func NewDB(c *conf.Data) *gorm.DB {
    defer func() {
        if e := recover(); e != nil {
            lg.Printf("open mysql failed, err: %v", e)
        }
    }()

    newLogger := logger.New(
        lg.New(os.Stdout, "\r\n", lg.LstdFlags), // io writer
        logger.Config{
            SlowThreshold: time.Second,   // 慢 SQL 阈值
            LogLevel:      logger.Silent, // Log level
            Colorful:      true,          // 彩色打印
        },
    )
    //m := global.GVA_CONFIG.Mysql
    m := Mysql{
        Host:         "1.14.163.5",
        Port:         "3306",
        Username:     "root",
        Password:     "1234",
        Dbname:       "user",
        MaxIdleConns: 64,
        MaxOpenConns: 128,
    }

    db, err := gorm.Open(mysql.New(mysql.Config{
        DSN:                       c.Database.Dsn, // DSN data source name
        DefaultStringSize:         191,            // string 类型字段的默认长度
        SkipInitializeWithVersion: true,           // 根据版本自动配置
    }), &gorm.Config{
        DisableForeignKeyConstraintWhenMigrating: true,
        QueryFields:                              true,                           // 这样查询的时候是用字段名称而不是*
        Logger:                                   newLogger.LogMode(logger.Info), // 开启info级别的日志
    })
    if err != nil {
        panic(err)
    }

    sqlDB, _ := db.DB() // 取出成员SqlDB来配置
    sqlDB.SetMaxIdleConns(m.MaxIdleConns)
    sqlDB.SetMaxOpenConns(m.MaxOpenConns)

    lg.Printf("init mysql successful")
    return db
}

func NewProdServiceClient(r registry.Discovery) v1.ProdServiceClient {
    conn, err := grpc.DialInsecure(
        context.Background(),
        grpc.WithEndpoint("discovery:///bingfood.prod.service"),
        grpc.WithDiscovery(r),
        grpc.WithMiddleware(
            recovery.Recovery(),
        ),
    )
    if err != nil {
        panic(err)
    }
    c := v1.NewProdServiceClient(conn)
    return c
}

func NewNsqServiceClient(ac *conf.JWT, r registry.Discovery) v13.NsqServiceClient {
    conn, err := grpc.DialInsecure(
        context.Background(),
        grpc.WithEndpoint("discovery:///bingfood.nsq.service"),
        grpc.WithDiscovery(r),
        grpc.WithMiddleware(
            recovery.Recovery(),
            jwt.Client(func(token *jwt2.Token) (interface{}, error) {
                return []byte(ac.ServiceSecretKey), nil
            }, jwt.WithSigningMethod(jwt2.SigningMethodHS256)),
        ),
    )
    if err != nil {
        panic(err)
    }
    c := v13.NewNsqServiceClient(conn)
    return c
}
