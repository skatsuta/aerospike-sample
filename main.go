package main

import (
	"flag"
	"fmt"

	asc "github.com/aerospike/aerospike-client-go"
)

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var host string
	var port int
	flag.StringVar(&host, "h", "127.0.0.1", "Aerospike host")
	flag.IntVar(&port, "p", 3000, "Aerospike port")
	flag.Parse()

	// define a client to connect to
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
	err = cl.Put(nil, key, bins)
	panicOnErr(err)

	rec, err := cl.Get(nil, key)
	panicOnErr(err)

	fmt.Printf("%#v\n", *rec)

	existed, err := cl.Delete(nil, key)
	panicOnErr(err)
	fmt.Printf("Record existed before delete? %v\n", existed)
}
