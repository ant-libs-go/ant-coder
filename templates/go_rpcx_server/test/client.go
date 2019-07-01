/* ######################################################################
# Author: (__AUTHOR__)
# Created Time: __CREATE_DATETIME__
# File Name: client.go
# Description:
####################################################################### */

package main

import (
	"context"
	"flag"
	"log"
	"time"

	"__PROJECT_NAME__/libs"

	"github.com/smallnest/rpcx/client"
)

var (
	method = flag.String("m", "", "call method")
)

func buildCommonHeader() *libs.Header {
	return &libs.Header{
		Requester: "test-client",
		Timestamp: time.Now().Unix(),
		Version:   100,
		Operator:  20,
		Metadata:  map[string]string{},
	}
}

func run(fun string, req, resp interface{}) (r interface{}) {
	d := client.NewConsulDiscovery("/__PROJECT_NAME___dev", "Server", []string{"127.0.0.1:8500"}, nil)
	cli := client.NewXClient("Server", client.Failover, client.RandomSelect, d, client.DefaultOption)
	defer cli.Close()

	log.Printf("Request: %v", req)
	log.Printf("Response: %v", resp)
	if err := cli.Call(context.Background(), fun, req, resp); err != nil {
		log.Printf("ERROR: %v", err)
	} else {
		log.Printf("RESULT: %v", resp)
	}
	return resp
}

func GetByIds() {
	req := &libs.GetByIdsRequest{
		Header: buildCommonHeader(),
		Body:   []int32{108, 109},
	}
	resp := &libs.GetByIdsResponse{}
	run("GetByIds", req, resp)
}

func main() {
	flag.Parse()
	switch *method {
	case "GetByIds":
		GetByIds()
	default:
		log.Println("not support.")
	}
}
