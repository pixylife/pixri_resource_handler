package pkg

import (
	"github.com/lucasb-eyer/go-colorful"
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
	colorOne := c
	colorTwo := colorful.Hsv(h+120.0,s,v)
	colorThree := colorful.Hsv(h-120.0,s,v)
	return colorOne,colorTwo,colorThree
}


