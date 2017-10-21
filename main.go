package main

import (
	"bufio"
	"bytes"
	"strconv"

	"github.com/hashicorp/go-plugin"
	"github.com/hexbotio/hex-plugin"
	"github.com/hexbotio/winrm"
)

type HexWinrm struct {
}

func (g *HexWinrm) Perform(args hexplugin.Arguments) (resp hexplugin.Response) {

	// initialize return values
	var output = ""
	var success = true

	port := 5985
	if args.Config["port"] != "" {
		port, _ = strconv.Atoi(args.Config["port"])
	}

	endpoint := winrm.NewEndpoint(args.Config["server"], port, false, false, nil, nil, nil, 0)
	rmclient, err := winrm.NewClient(endpoint, args.Config["login"], args.Config["pass"])
	if err != nil {
		output = "ERROR - Cannot connect to server " + args.Config["server"] + " (" + err.Error() + ")"
		success = false
	}

	var in bytes.Buffer
	var out bytes.Buffer
	var e bytes.Buffer

	stdin := bufio.NewReader(&in)
	stdout := bufio.NewWriter(&out)
	stderr := bufio.NewWriter(&e)

	_, err = rmclient.RunWithInput(args.Command, stdout, stderr, stdin)
	if err != nil {
		output = "ERROR - Error running command on server " + args.Config["server"] + " (" + err.Error() + ")"
		success = false
	}

	if e.String() != "" {
		output = e.String()
		success = false
	} else {
		output = out.String()
	}

	resp = hexplugin.Response{
		Output:  output,
		Success: success,
	}
	return resp
}

func main() {
	var pluginMap = map[string]plugin.Plugin{
		"action": &hexplugin.HexPlugin{Impl: &HexWinrm{}},
	}
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: hexplugin.GetHandshakeConfig(),
		Plugins:         pluginMap,
	})
}
