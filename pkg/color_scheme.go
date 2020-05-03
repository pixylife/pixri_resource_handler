package pkg

import (
	"fmt"
	"github.com/lucasb-eyer/go-colorful"
	"math"
)

func ColorBlend(c1 colorful.Color,c2 colorful.Color,variation int) colorful.Color{
	blocks := 10
	return c1.BlendRgb(c2, float64(variation)/float64(blocks-1))
}

func GetTextColor(c colorful.Color) colorful.Color{
		if(c.B+c.G+c.B)>0.179{
		 color,_ := colorful.Hex("#000000")
		return color
	}else {
		color,_ := colorful.Hex("#FFFFFF")
		return color
	}
}

func GetTriadColor(c colorful.Color) (colorful.Color,colorful.Color,colorful.Color){
	h,s,v := c.Hsv()
	c1 := c
	c2 := colorful.Hsv(math.Mod(h+120,360),s,v)
	c3 := colorful.Hsv(math.Mod(h-120,360),s,v)
	return c1,c2,c3
}

func GetComplementaryColor(c colorful.Color) (colorful.Color,colorful.Color){
	h,s,v := c.Hsv()
	c1 := c
	c2 := colorful.Hsv( math.Mod(h+180,360),s,v)
	fmt.Println( math.Mod(h+180,360))
	fmt.Println(c2.Hsv())
	return c1,c2
}

func GetSlitComplementaryColor(c1 colorful.Color,c2 colorful.Color) (colorful.Color,colorful.Color,colorful.Color){
	h,s,v := c2.Hsv()
	c2 = colorful.Hsv(math.Mod(h-20,360),s,v)
	c3 := colorful.Hsv(math.Mod(h+20,360),s,v)

	return c1,c2,c3
}

func GetMonochromaticSaturationColor(c colorful.Color)(colorful.Color,colorful.Color,colorful.Color){
	h,s,v := c.Hsv()
	c1 := c
	c2 := colorful.Hsv(h,math.Mod(s-0.15,1),v)
	c3 := colorful.Hsv(h,math.Mod(s+0.15,1),v)

	fmt.Println(c2.Hex())
	fmt.Println(c3.Hex())

	return c1,c2,c3
}

func GetMonochromaticValueColor(c colorful.Color)(colorful.Color,colorful.Color,colorful.Color){
	h,s,v := c.Hsv()
	c1 := c
	c2 := colorful.Hsv(h,s,math.Mod(v-0.15,1))
	c3 := colorful.Hsv(h,s,math.Mod(v+0.15,1))

	fmt.Println(c2.Hex())
	fmt.Println(c3.Hex())

	return c1,c2,c3
}