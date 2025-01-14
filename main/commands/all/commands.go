package all

import (
	"github.com/sqkam/xray-core/main/commands/all/api"
	"github.com/sqkam/xray-core/main/commands/all/convert"
	"github.com/sqkam/xray-core/main/commands/all/tls"
	"github.com/sqkam/xray-core/main/commands/base"
)

func init() {
	base.RootCommand.Commands = append(
		base.RootCommand.Commands,
		api.CmdAPI,
		convert.CmdConvert,
		tls.CmdTLS,
		cmdUUID,
		cmdX25519,
		cmdWG,
	)
}
