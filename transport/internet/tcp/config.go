package tcp

import (
	"github.com/sqkam/xray-core/common"
	"github.com/sqkam/xray-core/transport/internet"
)

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
