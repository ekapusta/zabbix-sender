package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"

	. "github.com/blacked/go-zabbix"
	"github.com/jimlawless/cfg"
)

var (
	configFile     string
	parameter      string
	parameterValue string
	stdin          *bufio.Reader
)

const (
	defaultPort = 10051
)

func init() {
	stdin = bufio.NewReader(os.Stdin)
	flag.StringVar(&configFile, "c", "/etc/zabbix/zabbix_agentd.conf", "Zabbix config file")
	flag.StringVar(&parameter, "p", "test", "Parameter")
	flag.StringVar(&parameterValue, "v", "0", "Parameter value")
}

func main() {
	flag.Parse()
	zbxCfg := make(map[string]string)
	err := cfg.Load(configFile, zbxCfg)
	if err != nil {
		panic(err)
	}

	var metrics []*Metric
	metrics = append(metrics, NewMetric(zbxCfg["hostname"], parameter, parameterValue, time.Now().Unix()))
	packet := NewPacket(metrics)
	z := NewSender(zbxCfg["server"], defaultPort)
	z.Send(packet)
	fmt.Println("OK")
}
