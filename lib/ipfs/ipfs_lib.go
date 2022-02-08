package ipfs

import (
	"errors"
	"github.com/gotechbook/application-notebook/logger"
	api "github.com/ipfs/go-ipfs-api"
	"strings"
)

func Client(url string) (cli *api.Shell, err error) {
	cli = api.NewShell(url)
	if cli.IsUp() {
		return cli, nil
	}
	err = errors.New("ipfs url is not connection")
	logger.Error("ipfs create client failed", err, url)
	return nil, err
}

func Upload(cli *api.Shell, meta string) (cid string, err error) {
	return cli.Add(strings.NewReader(meta))
}
