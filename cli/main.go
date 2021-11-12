package main

import (
	"apigo/crypto"
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Printf("rule-engine version: %s\n", 1)
	flag.PrintDefaults()
	os.Exit(0)
}

func main() {
	CfgVer := false
	CfgGen := false

	flag.BoolVar(&CfgGen, "g", false, "generate key")
	flag.BoolVar(&CfgVer, "v", false, "show version")
	flag.Usage = usage

	flag.Parse()

	if CfgVer {
		flag.Usage()
	}

	if CfgGen {
		pub, prv, err := crypto.GenRsaKey()
		if err != nil {
			fmt.Println("GenRsaKey err:", err)
		} else {
			fmt.Println("PublicKey====================")
			fmt.Println(pub)
			fmt.Println("PrivateKey====================")
			fmt.Println(prv)
		}
	}
}
