package param

import (
	"flag"
)

var Debug bool
var BrokerIP string
var BrokerPort string

func HandleFlags() {
	flag.StringVar(&BrokerIP, "brokerIP", "localhost", "Sets broker IP")
	flag.StringVar(&BrokerPort, "brokerPort", "1883", "Sets broker Port")
	flag.BoolVar(&Debug,"debug", false, "Debug Mode")

	flag.Parse()
}