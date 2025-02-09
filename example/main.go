package main

import (
	"log"

	. "github.com/go-stdx/rat"
)

func main() {
	log.Println(Rat("0.1").Add(0.2).IsEqual(Rat(0.3)))
	log.Println(Rat("0.1").Add(0.2).IsEqual(Rat(0.3)))
	log.Println(Rat("0.1").Add("0.2").IsEqual(Rat("0.3")))
}
