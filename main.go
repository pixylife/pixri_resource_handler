package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lucasb-eyer/go-colorful"
	_ "image"
	_ "image/draw"
	_ "pixri_resource_handler/pkg"
	"pixri_resource_handler/pkg/controller"
	_ "pixri_resource_handler/resources"
)

func main(){
/*color,_ := colorful.Hex(resources.Friendliness)
	c := pkg.GetTextColor(color)
	fmt.Println(c.Hex())
	fmt.Println(color.Hsv())

	fmt.Println("Triad")



	d1,d2,d3 := pkg.GetTriadColor(color)
	fmt.Println(d1.Hex())
	fmt.Println(d2.Hex())
	fmt.Println(d3.Hex())


	fmt.Println("Complementary")


	d1,d2 = pkg.GetComplementaryColor(color)
	fmt.Println(d1.Hex())
	fmt.Println(d2.Hex())

	pkg.GetSlitComplementaryColor(d1,d2)

	fmt.Println("Mono")
	pkg.GetMonochromaticSaturationColor(color)
	pkg.GetMonochromaticValueColor(color)

	clr, _ := colorful.Hex("#FFFF")
	clr2, _ := colorful.Hex("#0000")


	fmt.Println("SS")
	fmt.Println(color.Hex())
	fmt.Println(color.BlendLab(clr,0.144192442297936).Hex())
	fmt.Println(color.BlendLab(clr2,0.144192442297936).Hex())

	blocks := 4
	blockw := 40
	img := image.NewRGBA(image.Rect(0,0,blocks*blockw,200))*/
/*
	draw.Draw(img, image.Rect(0*blockw,  40,(0+1)*blockw, 80), &image.Uniform{C: color.BlendLab(clr,0.144192442297936)}, image.ZP, draw.Src)
	draw.Draw(img, image.Rect(1*blockw,  40,(1+1)*blockw, 80), &image.Uniform{C: color.BlendLab(clr,0.256376892328262)}, image.ZP, draw.Src)
	draw.Draw(img, image.Rect(2*blockw,  40,(2+1)*blockw, 80), &image.Uniform{C: color.BlendLab(clr,0.412043422460556)}, image.ZP, draw.Src)
	draw.Draw(img, image.Rect(3*blockw,  40,(3+1)*blockw, 80), &image.Uniform{C: color.BlendLab(clr,0.195936158299446)}, image.ZP, draw.Src)
	draw.Draw(img, image.Rect(4*blockw,  40,(4+1)*blockw, 80), &image.Uniform{C: color.BlendLab(clr,0.184556514024734)}, image.ZP, draw.Src)

*/
	/*draw.Draw(img, image.Rect(0*blockw,  0,(0+1)*blockw, 0), &image.Uniform{C: color.BlendLab(clr2,0.0891487151384354)}, image.ZP, draw.Src)
	draw.Draw(img, image.Rect(1*blockw,  0,(1+1)*blockw, 40), &image.Uniform{C: color.BlendLab(clr2,0.00015455)}, image.ZP, draw.Src)
	draw.Draw(img, image.Rect(2*blockw,  0,(2+1)*blockw, 80), &image.Uniform{C: color.BlendLab(clr2,0.145223319530487)}, image.ZP, draw.Src)
	draw.Draw(img, image.Rect(3*blockw,  0,(3+1)*blockw, 120), &image.Uniform{C: color.BlendLab(clr2,0.121651209890842)}, image.ZP, draw.Src)
	draw.Draw(img, image.Rect(4*blockw,  0,(4+1)*blockw, 200), &image.Uniform{C: color.BlendLab(clr2,0.593493841588497)}, image.ZP, draw.Src)

	fmt.Println("xxx")
	fmt.Println(color.BlendLab(clr,0.0891487151384354).Hsv())
	fmt.Println(color.BlendLab(clr,0.00015455).Hsv())
	fmt.Println(color.BlendLab(clr,0.145223319530487).Hsv())
	fmt.Println(color.BlendLab(clr,0.121651209890842).Hsv())
	fmt.Println(color.BlendLab(clr,0.593493841588497).Hsv())

*/





	/*toimg, err := os.Create("test.png")
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	defer toimg.Close()

	png.Encode(toimg, img)*/





	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	r := e.Group("/")
	controller.ThemeController(r, "api")
	e.Logger.Fatal(e.Start(":5002"))

}