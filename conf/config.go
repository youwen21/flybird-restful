package conf

type AppConfig struct {
	Smtp SmtpCfg `mapstructure:"smtp"` //,squash

	Mysql MysqlCfg

	RedisDefault RedisCfg `toml:"redis_default" mapstructure:"redis_default"` //,
}

var Config *AppConfig
