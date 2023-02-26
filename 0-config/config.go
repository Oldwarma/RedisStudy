package __config

type Config struct {
	DataSourceConf DataSourceConfig
	RedisConfig    RedisConfig
}

type DataSourceConfig struct {
	Mysql   MysqlConfig
	LogMode string // 是否开启Gorm全局日志
}

type MysqlConfig struct {
	Address      string // 服务器地址:端口
	DbName       string // 数据库名
	Username     string // 数据库用户名
	Password     string // 数据库密码
	MaxIdleConns int    // 空闲中的最大连接数
	MaxOpenConns int    // 打开到数据库的最大连接数
}
type RedisConfig struct {
	Host           string //redis ip
	Port           string
	Database       string //数据库名
	Password       string //密码
	ConnectTimeout int    //链接超时时间
	SubscribeChan  string //订阅的通道
	RedisPool      RedisPoolStruct
}

type RedisPoolStruct struct {
	PoolMaxIdle     int //最大空闲数
	PoolMaxActive   int //最大连接个数
	PoolIdleTimeout int //最大的空闲连接等待时间
}
