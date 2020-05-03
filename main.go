package main

import (
	"fmt"
	"github.com/lucasb-eyer/go-colorful"
	"pixri_resource_handler/pkg"
	_ "pixri_resource_handler/pkg"
)

func main(){
color,_ := colorful.Hex("#ff0a2f")
	c := pkg.GetTextColor(color)
	fmt.Println(c.Hex())
	fmt.Println(color.Hsv())

	fmt.Println("Triad")
	c1,c2,c3 := pkg.GetTriadColor(c)
	fmt.Println(c1.Hex())
	fmt.Println(c2.Hex())
	fmt.Println(c3.Hex())


	d1,d2,d3 := pkg.GetTriadColor(color)
	fmt.Println(d1.Hex())
	fmt.Println(d2.Hex())
	fmt.Println(d3.Hex())


	fmt.Println("Complementary")
	c1,c2 = pkg.GetComplementaryColor(c)
	fmt.Println(c1.Hex())
	fmt.Println(c2.Hex())

	d1,d2 = pkg.GetComplementaryColor(color)
	fmt.Println(d1.Hex())
	fmt.Println(d2.Hex())

	pkg.GetSlitComplementaryColor(d1,d2)


}