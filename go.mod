module github.com/samlet/petrel

go 1.16

require (
	github.com/Microsoft/go-winio v0.5.0 // indirect
	github.com/anmitsu/go-shlex v0.0.0-20161002113705-648efa622239 // indirect
	github.com/apache/thrift v0.0.0-20160514152517-9549b25c7758 // indirect
	github.com/bmizerany/perks v0.0.0-20141205001514-d9a9656a3a4b // indirect
	github.com/centrifuge/go-substrate-rpc-client/v3 v3.0.0
	github.com/containerd/containerd v1.5.2 // indirect
	github.com/crossdock/crossdock-go v0.0.0-20160816171116-049aabb0122b
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/docker/docker v20.10.6+incompatible
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/ethereum/go-ethereum v1.10.3 // indirect
	github.com/facebookgo/clock v0.0.0-20150410010913-600d898af40a // indirect
	github.com/fatih/color v1.11.0
	github.com/fatih/structtag v1.0.0 // indirect
	github.com/flynn/go-shlex v0.0.0-20150515145356-3f9db97f8568 // indirect
	github.com/gogo/status v1.1.0 // indirect
	github.com/golang/mock v1.4.1
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/jessevdk/go-flags v1.4.0 // indirect
	github.com/kr/pretty v0.2.1
	github.com/labstack/echo/v4 v4.3.0 // indirect
	github.com/mattn/go-sqlite3 v1.14.7 // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/pborman/uuid v0.0.0-20180906182336-adf5a7427709
	github.com/prashantv/protectmem v0.0.0-20171002184600-e20412882b3a // indirect
	github.com/robfig/cron v1.2.0 // indirect
	github.com/samuel/go-thrift v0.0.0-20190219015601-e8b6b52668fe // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/streadway/quantile v0.0.0-20150917103942-b0c588724d25 // indirect
	github.com/stretchr/testify v1.7.0
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
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a // indirect
	golang.org/x/net v0.0.0-20210525063256-abc453219eb5 // indirect
	golang.org/x/sys v0.0.0-20210525143221-35b2ab0089ea // indirect
	google.golang.org/genproto v0.0.0-20210524171403-669157292da3 // indirect
	google.golang.org/grpc v1.38.0 // indirect
	gopkg.in/yaml.v2 v2.4.0
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.21.10
	k8s.io/apimachinery v0.21.1
	k8s.io/client-go v0.21.1
)

//replace github.com/centrifuge/go-substrate-rpc-client => ../go-substrate-rpc-client-master
