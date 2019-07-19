package envconfig

import (
	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type mysqlDB struct {
	DataSource   string `toml:"data_source"`
	MaxIdleConns int    `toml:"max_idle_conns"`
	MaxOpenConns int    `toml:"max_open_conns"`
}

type http struct {
	Listen string `toml:"listen"`
}

type logConf struct {
	OpenLogstash    bool   `toml:"open_logstash"`
	OpenLogfile     bool   `toml:"open_logfile"`
	FilePath        string `toml:"file_path"`
	Project         string `toml:"project"`
	Module          string `toml:"module"`
	LogstashAddress string `toml:"logstash_address"`
}

type redisConf struct {
	Address  string `toml:"address"`
	Password string `toml:"password"`
	DB       int    `toml:"db"`
}

type rabbitMQ struct {
	URL                                     string `toml:"url"`
	Exchange                                string `toml:"exchange"`
	RoutingKeyAftersalesStatusChange        string `toml:"routing_key_aftersales_status_change"`
	ExchangeSMS                             string `toml:"exchange_sms"`
	RoutingKeySMSMarketing                  string `toml:"routing_key_sms_marketing"`
	AppWebHost                              string `toml:"app_web_host"`
	ExchangePushNotification                string `toml:"exchange_push_notification"`
	RoutingKeyPushNotification              string `toml:"routing_key_push_notification"`
	RoutingKeyRefundBankChange              string `toml:"routing_key_refund_bank_change"`
	QueueValueAftersalesRefundBankStatus    string `toml:"queue_value_aftersales_refund_bank_status"`
	RoutingKeyPaymentStatusChange           string `toml:"routing_key_payment_status_change"`
	QueueValueAftersalesPaymentStatus       string `toml:"queue_value_aftersales_payment_status"`
	RoutingKeyOrderStatusChange             string `toml:"routing_key_order_status_change"`
	QueueValueAftersalesOrderStatus         string `toml:"queue_value_aftersales_order_status"`
	RoutingKeyPurchaseOrderStatusChange     string `toml:"routing_key_purchase_order_status_change"`
	QueueValueAftersalesPurchaseOrderStatus string `toml:"queue_value_aftersales_purchase_order_status"`
}

type jwtToken struct {
	AppName      string `toml:"app_name"`
	SecretKey    string `toml:"secret_key"`
	ExpInSeconds int    `toml:"exp_in_seconds"`
	Enabled      bool   `toml:"enabled"`
}

type oss struct {
	Provider    string `toml:"provider"`
	ImageBucket string `toml:"image_bucket"`
	Endpoint    string `toml:"endpoint"`
	AccessKey   string `toml:"access_key"`
	SecretKey   string `toml:"secret_key"`
}

type internalURL struct {
	InternalShopintar     string `toml:"internal_shopintar"`
	InternalOrder         string `toml:"internal_order"`
	InternalUser          string `toml:"internal_user"`
	InternalProduct       string `toml:"internal_product"`
	InternalPurchaseOrder string `toml:"internal_purchase_order"`
	InternalPayment       string `toml:"internal_payment"`
}

type config struct {
	Http        http        `toml:"http"`
	JWTToken    jwtToken    `toml:"jwt_token"`
	Log         logConf     `toml:"log_conf"`
	MysqlDB     mysqlDB     `toml:"mysql_db"`
	RedisConf   redisConf   `toml:"redis_conf"`
	RabbitMQ    rabbitMQ    `toml:"rabbit_mq"`
	OSS         oss         `toml:"oss"`
	InternalURL internalURL `toml:"internal_url"`
}

func TestParseConfig(t *testing.T) {
	var conf config
	ParseConfig("/Users/cookie/go/gopath/src/go-sillyhat-cloud/config.conf", func(content []byte) {
		err := toml.Unmarshal(content, &conf)
		if err != nil {
			log.Panicf("unmarshal toml object error. %v", err)
		}
	})
	assert.NotNil(t, conf)
	assert.EqualValues(t, conf.Http.Listen, ":8080")
}
