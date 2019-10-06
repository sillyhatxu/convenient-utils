package config

var Conf config

type config struct {
	Project                      string
	Module                       string
	Host                         string
	HttpPort                     int     `toml:"http_port"`
	GRPCPort                     int     `toml:"grpc_port"`
	Log                          LogConf `toml:"log_conf"`
	DBMotherUserName             string  `env:"SILLYHAT.DB.MOTHER.USERNAME"`
	DBMotherPassword             string  `env:"SILLYHAT.DB.MOTHER.PASSWORD"`
	DBMotherHost                 string  `env:"SILLYHAT.DB.MOTHER.HOST"`
	DBMotherPort                 string  `env:"SILLYHAT.DB.MOTHER.PORT"`
	DBMotherSchema               string  `env:"SILLYHAT.DB.MOTHER.SCHEMA"`
	DBUserUserName               string  `env:"SILLYHAT.DB.USER.USERNAME"`
	DBUserPassword               string  `env:"SILLYHAT.DB.USER.PASSWORD"`
	DBUserHost                   string  `env:"SILLYHAT.DB.USER.HOST"`
	DBUserPort                   string  `env:"SILLYHAT.DB.USER.PORT"`
	DBUserSchema                 string  `env:"SILLYHAT.DB.USER.SCHEMA"`
	DBWordUserName               string  `env:"SILLYHAT.DB.WORD.USERNAME"`
	DBWordPassword               string  `env:"SILLYHAT.DB.WORD.PASSWORD"`
	DBWordHost                   string  `env:"SILLYHAT.DB.WORD.HOST"`
	DBWordPort                   string  `env:"SILLYHAT.DB.WORD.PORT"`
	DBWordSchema                 string  `env:"SILLYHAT.DB.WORD.SCHEMA"`
	DBDDLPath                    string  `env:"SILLYHAT.DB.DDL.PATH"`
	LogstashURL                  string  `env:"SILLYHAT.LOGSTASH.URL"`
	OSSImageBucket               string  `env:"SILLYHAT.OSS.IMAGE_BUCKET"`
	OSSEndpoint                  string  `env:"SILLYHAT.OSS.ENDPOINT"`
	OSSAccessKey                 string  `env:"SILLYHAT.OSS.ACCESS_KEY"`
	OSSSecretKey                 string  `env:"SILLYHAT.OSS.SECRET_KEY"`
	MiniMQHost                   string  `env:"SILLYHAT.MINIMQ.HOST"`
	MiniMQApiPort                string  `env:"SILLYHAT.MINIMQ.API.PORT"`
	MiniMQGRPCPort               string  `env:"SILLYHAT.MINIMQ.GRPC.PORT"`
	MiniMQUserName               string  `env:"SILLYHAT.MINIMQ.USERNAME"`
	MiniMQPassword               string  `env:"SILLYHAT.MINIMQ.PASSWORD"`
	HostConsul                   string  `env:"SILLYHAT.HOST.CONSUL"`
	HostMiniMQHttp               string  `env:"SILLYHAT.HOST.MINI_MQ.HTTP"`
	HostMiniMQGRPC               string  `env:"SILLYHAT.HOST.MINI_MQ.GRPC"`
	HostUserGRPC                 string  `env:"SILLYHAT.HOST.USER.GRPC"`
	HostWordGRPC                 string  `env:"SILLYHAT.HOST.WORD.GRPC"`
	HostMessageGRPC              string  `env:"SILLYHAT.HOST.MESSAGE.GRPC"`
	MessageSlackHost             string  `env:"MESSAGE.SLACK.HOST"`
	MessageSlackBirthday         string  `env:"MESSAGE.SLACK.BIRTHDAY"`
	MessageSlackRemind           string  `env:"MESSAGE.SLACK.REMIND"`
	MessageSlackSillyHat         string  `env:"MESSAGE.SLACK.SILLYHAT"`
	MessageSlackSillyHatErrorLog string  `env:"MESSAGE.SLACK.SILLYHAT.ERROR.LOG"`
}

type LogConf struct {
	OpenLogstash bool   `toml:"open_logstash"`
	OpenLogfile  bool   `toml:"open_logfile"`
	FilePath     string `toml:"file_path"`
}
