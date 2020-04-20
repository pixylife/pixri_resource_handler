package main

import (
	"fmt"
	"github.com/lucasb-eyer/go-colorful"
	"pixri_resource_handler/pkg"
	_ "pixri_resource_handler/pkg"
)

func main(){
color,_ := colorful.Hex("#9c64a6")
	c := pkg.GetTextColor(color)
	fmt.Println(c.Hex())
}