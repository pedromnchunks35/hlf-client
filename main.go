package main

import (
	ch "basic/chaincode"
	c "basic/connection"
	"log"
)

func main() {
	gateway, err := c.CreateConnection(
		"./msp/tlscacerts/tls-hlf-premise-tls-ca-30000.pem",
		"hlf-network:30007",
		"./msp/signcerts/cert.pem",
		"OrgxMSP",
		"./msp/keystore/",
	)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer gateway.Close()
	channel := gateway.GetNetwork("channel1")
	contract := channel.GetContract("basic")
	ch.CreateAsset(contract, "psaihfoihasfhaoifhoi", "35")
}
