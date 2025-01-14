package command_test

import (
	"context"
	"testing"

	"github.com/sqkam/xray-core/app/dispatcher"
	"github.com/sqkam/xray-core/app/log"
	. "github.com/sqkam/xray-core/app/log/command"
	"github.com/sqkam/xray-core/app/proxyman"
	_ "github.com/sqkam/xray-core/app/proxyman/inbound"
	_ "github.com/sqkam/xray-core/app/proxyman/outbound"
	"github.com/sqkam/xray-core/common"
	"github.com/sqkam/xray-core/common/serial"
	"github.com/sqkam/xray-core/core"
)

func TestLoggerRestart(t *testing.T) {
	v, err := core.New(&core.Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&log.Config{}),
			serial.ToTypedMessage(&dispatcher.Config{}),
			serial.ToTypedMessage(&proxyman.InboundConfig{}),
			serial.ToTypedMessage(&proxyman.OutboundConfig{}),
		},
	})
	common.Must(err)
	common.Must(v.Start())

	server := &LoggerServer{
		V: v,
	}
	common.Must2(server.RestartLogger(context.Background(), &RestartLoggerRequest{}))
}
