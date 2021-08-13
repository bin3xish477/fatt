package main

import (
	"fmt"

	"github.com/alexflint/go-arg"
	"github.com/binexisHATT/fatt/models"
)

func main() {
	c := models.Args{}
	arg.MustParse(&c)
	fmt.Printf("%+v", c)
}
