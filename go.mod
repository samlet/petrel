module github.com/samlet/petrel

go 1.16

require (
	entgo.io/ent v0.8.0
	github.com/GeertJohan/go.rice v1.0.2
	github.com/Microsoft/go-winio v0.5.0 // indirect
	github.com/alecthomas/kong v0.2.16
	github.com/alecthomas/participle/v2 v2.0.0-alpha6
	github.com/alecthomas/repr v0.0.0-20210301060118-828286944d6a
	github.com/antchfx/htmlquery v1.2.3
	github.com/apache/thrift v0.0.0-20160514152517-9549b25c7758 // indirect
	github.com/beevik/etree v1.1.0 // indirect
	github.com/bmizerany/perks v0.0.0-20141205001514-d9a9656a3a4b // indirect
	github.com/bwmarrin/snowflake v0.3.0
	github.com/casbin/casbin/v2 v2.31.2 // indirect
	github.com/casbin/xorm-adapter/v2 v2.3.2 // indirect
	github.com/centrifuge/go-substrate-rpc-client/v3 v3.0.0
	github.com/containerd/containerd v1.5.2 // indirect
	github.com/crossdock/crossdock-go v0.0.0-20160816171116-049aabb0122b
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/docker/docker v20.10.6+incompatible
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/ethereum/go-ethereum v1.10.3
	github.com/facebookgo/clock v0.0.0-20150410010913-600d898af40a // indirect
	github.com/fatih/camelcase v1.0.0
	github.com/fatih/color v1.11.0
	github.com/fatih/structtag v1.0.0 // indirect
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/go-sql-driver/mysql v1.5.1-0.20200311113236-681ffa848bae // indirect
	github.com/gogo/status v1.1.0 // indirect
	github.com/golang/mock v1.4.4
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/wire v0.5.0 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/hyperjumptech/grule-rule-engine v1.9.0
	github.com/iancoleman/strcase v0.1.3
	github.com/jmoiron/sqlx v1.3.4
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51 // indirect
	github.com/kr/pretty v0.2.1
	github.com/labstack/echo/v4 v4.3.0
	github.com/lib/pq v1.10.0
	github.com/looplab/fsm v0.2.0
	github.com/magefile/mage v1.11.0
	github.com/mattn/go-sqlite3 v1.14.7
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/olivere/elastic/v7 v7.0.24
	github.com/pborman/uuid v0.0.0-20180906182336-adf5a7427709
	github.com/prashantv/protectmem v0.0.0-20171002184600-e20412882b3a // indirect
	github.com/robfig/cron v1.2.0 // indirect
	github.com/samuel/go-thrift v0.0.0-20190219015601-e8b6b52668fe // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/segmentio/kafka-go v0.4.16
	github.com/shomali11/slacker v0.0.0-20210527125624-74795172e04f // indirect
	github.com/shopspring/decimal v1.2.0
	github.com/streadway/amqp v1.0.0
	github.com/streadway/quantile v0.0.0-20150917103942-b0c588724d25 // indirect
	github.com/stretchr/testify v1.7.0
	github.com/tinylib/msgp v1.1.5 // indirect
	github.com/uber-go/mapdecode v1.0.0 // indirect
	github.com/uber-go/tally v3.3.11+incompatible
	github.com/uber/jaeger-client-go v2.18.1+incompatible // indirect
	github.com/uber/jaeger-lib v2.1.1+incompatible // indirect
	github.com/uber/tchannel-go v1.14.0 // indirect
	github.com/urfave/cli/v2 v2.3.0
	go.uber.org/cadence v0.8.9-0.20190711214419-cef99c5ba19e
	go.uber.org/net/metrics v1.1.0 // indirect
	go.uber.org/thriftrw v1.20.0 // indirect
	go.uber.org/tools v0.0.0-20190618225709-2cfd321de3ee // indirect
	go.uber.org/yarpc v1.39.0
	go.uber.org/zap v1.10.0
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a
	golang.org/x/net v0.0.0-20210525063256-abc453219eb5
	golang.org/x/sys v0.0.0-20210525143221-35b2ab0089ea // indirect
	golang.org/x/text v0.3.6
	google.golang.org/genproto v0.0.0-20210524171403-669157292da3 // indirect
	google.golang.org/grpc v1.38.0 // indirect
	google.golang.org/protobuf v1.26.0 // indirect
	gopkg.in/yaml.v2 v2.4.0
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.21.10
	k8s.io/apimachinery v0.21.1
	k8s.io/client-go v0.21.1
	zombiezen.com/go/capnproto2 v2.18.2+incompatible

)

//replace github.com/centrifuge/go-substrate-rpc-client => ../go-substrate-rpc-client-master
