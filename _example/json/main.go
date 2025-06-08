package main

import (
	"encoding/json"
	"log"

	"github.com/investing-kr/bot/rat"
)

func main() {
	// log.Println(rat.Rat("0.95").Ceil())
	// log.Println(rat.Rat("4").Ceil())
	// log.Println(rat.Rat("7.004").Ceil())
	// log.Println(rat.Rat("-7.004").Ceil())
	b := Balance{}
	err := json.Unmarshal([]byte(msg), &b)
	if err != nil {
		panic(err)
	}
	log.Printf("%+v", b)

	log.Println("------")

	rat.DefaultPrecision = 20
	v1 := rat.Rat("1386929.37231066771348207123")
	log.Println("v1", v1)
	v2 := rat.Rat("1386929.3723107")
	log.Println("v2", v2)
	log.Println("v2 is equal", v2.IsEqual(v1))
}

var msg  = `
{
   "balance": 1386929.37231066771348207123,
   "currency": "KRW"
}
`

type Balance struct {
	Balance  *rat.Rational `json:"balance"`
	Currency string  `json:"currency"`
}



