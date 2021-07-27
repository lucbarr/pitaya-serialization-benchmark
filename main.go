package main

import (
	"flag"
	"fmt"
	"pitaya-serialization-benchmark/services"
	"strings"

	"github.com/spf13/viper"
	"github.com/topfreegames/pitaya"
	"github.com/topfreegames/pitaya/acceptor"
	"github.com/topfreegames/pitaya/component"
	"github.com/topfreegames/pitaya/serialize/protobuf"
)

func main() {
	port := flag.Int("port", 3250, "the port to listen")
	svType := flag.String("type", "connector", "the server type")
	isFrontend := flag.Bool("frontend", true, "if server is frontend")
	useProto := flag.Bool("useProto", false, "if server uses protobuf")

	flag.Parse()

	defer pitaya.Shutdown()

	if *useProto {
		ser := protobuf.NewSerializer()
		pitaya.SetSerializer(ser)
	}

	if !*isFrontend {
		configureBackend()
	} else {
		configureFrontend(*port)
	}

	config := viper.New()

	config.SetEnvPrefix("PSB")
	config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	pitaya.Configure(*isFrontend, *svType, pitaya.Cluster, map[string]string{}, config)
	pitaya.Start()
}

func configureBackend() {
	pitaya.Register(&services.PingHandler{},
		component.WithName("pinghandler"),
		component.WithNameFunc(strings.ToLower),
	)

	pitaya.Register(&services.DocsHandler{},
		component.WithName("docs"),
		component.WithNameFunc(strings.ToLower),
	)

	pitaya.Register(&services.BenchmarkHandler{},
		component.WithName("benchmark"),
		component.WithNameFunc(strings.ToLower),
	)
}

func configureFrontend(port int) {
	ws := acceptor.NewWSAcceptor(fmt.Sprintf(":%d", port))
	pitaya.Register(&services.Connector{},
		component.WithName("connector"),
		component.WithNameFunc(strings.ToLower),
	)
	pitaya.RegisterRemote(&services.ConnectorRemote{},
		component.WithName("connectorremote"),
		component.WithNameFunc(strings.ToLower),
	)

	pitaya.AddAcceptor(ws)
}
