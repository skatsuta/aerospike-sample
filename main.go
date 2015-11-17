package main

import (
	"fmt"
	"os"

	asc "github.com/aerospike/aerospike-client-go"
	"github.com/k0kubun/pp"
)

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// define a client to connect to
	host := os.Getenv("AEROSPIKE_PORT_3000_TCP_ADDR")
	port := 3000
	cl, err := asc.NewClient(host, port)
	panicOnErr(err)

	key, err := asc.NewKey("test", "aerospike", "key")
	panicOnErr(err)

	// define some bins with data
	bins := asc.BinMap{
		"bin1": 42,
		"bin2": "elephant",
		"bin3": []interface{}{"Go", 2009},
	}

	// write the bins
	wp := asc.NewWritePolicy(0, 0)
	wp.SendKey = true
	err = cl.Put(wp, key, bins)
	panicOnErr(err)

	rec, err := cl.Get(nil, key)
	panicOnErr(err)

	_, _ = pp.Printf("key: %v\nbins: %v\n", *rec.Key, rec.Bins)

	// scan all data
	rs, err := cl.ScanAll(nil, "test", "aerospike")
	panicOnErr(err)
	defer rs.Close()

	for r := range rs.Results() {
		_, _ = pp.Println(*r.Record.Key, r.Record.Bins)
	}

	existed, err := cl.Delete(nil, key)
	panicOnErr(err)
	fmt.Printf("Record existed before delete? %v\n", existed)
}
