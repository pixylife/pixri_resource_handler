package generator

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/lucasb-eyer/go-colorful"
	"pixri_resource_handler/pkg"
	"pixri_resource_handler/pkg/model"
	"pixri_resource_handler/resources"
)

const baseUrl = "http://127.0.0.1:5000/"

type Test struct {
	Success int    `json:"success"`
	AppName string `json:"app_name"`
}

type ColorScore struct {
	Success  int    `json:"success"`
	AppName  string `json:"app_name"`
	KeyWords struct {
		Power          float64 `json:"power"`
		Importance     float64 `json:"importance"`
		Youth          float64 `json:"youth"`
		Excitement     float64 `json:"excitement"`
		Bold           float64 `json:"bold"`
		Passion        float64 `json:"passion"`
		Friendliness   float64 `json:"friendliness"`
		Energy         float64 `json:"energy"`
		Uniqueness     float64 `json:"uniqueness"`
		Confidence     float64 `json:"confidence"`
		Cheerful       float64 `json:"cheerful"`
		Happiness      float64 `json:"happiness"`
		Enthusiasm     float64 `json:"enthusiasm"`
		Antiquity      float64 `json:"antiquity"`
		Warmth         float64 `json:"warmth"`
		Optimism       float64 `json:"optimism"`
		Clarity        float64 `json:"clarity"`
		Growth         float64 `json:"growth"`
		Stability      float64 `json:"stability"`
		Financial      float64 `json:"financial"`
		Environmental  float64 `json:"environmental"`
		Health         float64 `json:"health"`
		Peaceful       float64 `json:"peaceful"`
		Calm           float64 `json:"calm"`
		Safety         float64 `json:"safety"`
		Openness       float64 `json:"openness"`
		Reliability    float64 `json:"reliability"`
		Trust          float64 `json:"trust"`
		Strength       float64 `json:"strength"`
		Dependable     float64 `json:"dependable"`
		Luxury         float64 `json:"luxury"`
		Romance        float64 `json:"romance"`
		Mystery        float64 `json:"mystery"`
		Creative       float64 `json:"creative"`
		Wise           float64 `json:"wise"`
		Imaginative    float64 `json:"imaginative"`
		Edginess       float64 `json:"edginess"`
		Sophistication float64 `json:"sophistication"`
		Cleanliness    float64 `json:"cleanliness"`
		Virtue         float64 `json:"virtue"`
		Simplicity     float64 `json:"simplicity"`
		Neutrality     float64 `json:"neutrality"`
		Formality      float64 `json:"formality"`
		Melancholy     float64 `json:"melancholy"`
		Balance        float64 `json:"balance"`
	} `json:"key_words"`
}




func getColorScore(app_name string, app_desc string) ColorScore{
//	defer wg.Done()

	client := resty.New()
	 resp,_ := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody("{\"app_name\":\""+app_name+"\",\"app_desc\":\""+app_desc+"\"}").
		Post(baseUrl)


	if resp.StatusCode() == 200 {
		data := ColorScore{}
		_ = json.Unmarshal([]byte(resp.String()), &data)
		return data

	} else {
		return ColorScore{}
	}
}

func mapColorLogic(app_name string, app_desc string) pkg.KeywordList {
	var keywords = pkg.KeywordList{}

	col := 	getColorScore(app_name, app_desc)


	keywords = append(keywords, pkg.Keyword{Key: "Power", Value: col.KeyWords.Power})
	keywords = append(keywords, pkg.Keyword{Key: "Importance", Value: col.KeyWords.Importance})
	keywords = append(keywords, pkg.Keyword{Key: "Youth", Value: col.KeyWords.Youth})
	keywords = append(keywords, pkg.Keyword{Key: "Excitement", Value: col.KeyWords.Excitement})
	keywords = append(keywords, pkg.Keyword{Key: "Bold", Value: col.KeyWords.Bold})
	keywords = append(keywords, pkg.Keyword{Key: "Passion", Value: col.KeyWords.Passion})
	keywords = append(keywords, pkg.Keyword{Key: "Friendliness", Value: col.KeyWords.Friendliness})
	keywords = append(keywords, pkg.Keyword{Key: "Energy", Value: col.KeyWords.Energy})
	keywords = append(keywords, pkg.Keyword{Key: "Uniqueness", Value: col.KeyWords.Uniqueness})
	keywords = append(keywords, pkg.Keyword{Key: "Confidence", Value: col.KeyWords.Confidence})
	keywords = append(keywords, pkg.Keyword{Key: "Cheerful", Value: col.KeyWords.Cheerful})
	keywords = append(keywords, pkg.Keyword{Key: "Happiness", Value: col.KeyWords.Happiness})
	keywords = append(keywords, pkg.Keyword{Key: "Enthusiasm", Value: col.KeyWords.Enthusiasm})
	keywords = append(keywords, pkg.Keyword{Key: "Antiquity", Value: col.KeyWords.Antiquity})
	keywords = append(keywords, pkg.Keyword{Key: "Warmth", Value: col.KeyWords.Warmth})
	keywords = append(keywords, pkg.Keyword{Key: "Optimism", Value: col.KeyWords.Optimism})
	keywords = append(keywords, pkg.Keyword{Key: "Clarity", Value: col.KeyWords.Clarity})
	keywords = append(keywords, pkg.Keyword{Key: "Growth", Value: col.KeyWords.Growth})
	keywords = append(keywords, pkg.Keyword{Key: "Stability", Value: col.KeyWords.Stability})
	keywords = append(keywords, pkg.Keyword{Key: "Financial", Value: col.KeyWords.Financial})
	keywords = append(keywords, pkg.Keyword{Key: "Environmental", Value: col.KeyWords.Environmental})
	keywords = append(keywords, pkg.Keyword{Key: "Health", Value: col.KeyWords.Health})
	keywords = append(keywords, pkg.Keyword{Key: "Peaceful", Value: col.KeyWords.Peaceful})
	keywords = append(keywords, pkg.Keyword{Key: "Calm", Value: col.KeyWords.Calm})
	keywords = append(keywords, pkg.Keyword{Key: "Safety", Value: col.KeyWords.Safety})
	keywords = append(keywords, pkg.Keyword{Key: "Openness", Value: col.KeyWords.Openness})
	keywords = append(keywords, pkg.Keyword{Key: "Reliability", Value: col.KeyWords.Reliability})
	keywords = append(keywords, pkg.Keyword{Key: "Trust", Value: col.KeyWords.Trust})
	keywords = append(keywords, pkg.Keyword{Key: "Strength", Value: col.KeyWords.Strength})
	keywords = append(keywords, pkg.Keyword{Key: "Dependable", Value: col.KeyWords.Dependable})
	keywords = append(keywords, pkg.Keyword{Key: "Luxury", Value: col.KeyWords.Luxury})
	keywords = append(keywords, pkg.Keyword{Key: "Romance", Value: col.KeyWords.Romance})
	keywords = append(keywords, pkg.Keyword{Key: "Mystery", Value: col.KeyWords.Mystery})
	keywords = append(keywords, pkg.Keyword{Key: "Creative", Value: col.KeyWords.Creative})
	keywords = append(keywords, pkg.Keyword{Key: "Wise", Value: col.KeyWords.Wise})
	keywords = append(keywords, pkg.Keyword{Key: "Imaginative", Value: col.KeyWords.Imaginative})
	keywords = append(keywords, pkg.Keyword{Key: "Edginess", Value: col.KeyWords.Edginess})
	keywords = append(keywords, pkg.Keyword{Key: "Sophistication", Value: col.KeyWords.Sophistication})
	keywords = append(keywords, pkg.Keyword{Key: "Cleanliness", Value: col.KeyWords.Cleanliness})
	keywords = append(keywords, pkg.Keyword{Key: "Virtue", Value: col.KeyWords.Virtue})
	keywords = append(keywords, pkg.Keyword{Key: "Simplicity", Value: col.KeyWords.Simplicity})
	keywords = append(keywords, pkg.Keyword{Key: "Neutrality", Value: col.KeyWords.Neutrality})
	keywords = append(keywords, pkg.Keyword{Key: "Formality", Value: col.KeyWords.Formality})
	keywords = append(keywords, pkg.Keyword{Key: "Formality", Value: col.KeyWords.Formality})
	keywords = append(keywords, pkg.Keyword{Key: "Melancholy", Value: col.KeyWords.Melancholy})
	keywords = append(keywords, pkg.Keyword{Key: "Balance", Value: col.KeyWords.Balance})

	return keywords
}

func sortScore(app_name string, app_desc string) pkg.KeywordList {
	keyList := mapColorLogic(app_name, app_desc)

	pkg.GetMaxKeywords(&keyList)
	return keyList
}

func GenerateTheme(app_name string, app_desc string, app_id int) []model.Theme {
	var themes = []model.Theme{}
	keyList := sortScore(app_name, app_desc)

	fmt.Println(keyList)

	for i := 0; i < 3; i++ {
		keyword := keyList[i]

		primaryColor := resources.ColorList[keyword.Key]
		fmt.Println(primaryColor)
		color, _ := colorful.Hex(primaryColor)
		blendColor := pkg.GetColorBlend(color, keyword.Value)
		fmt.Println(blendColor.Hex())

		for j := 0; j < 3; j++ {
			var theme = model.Theme{}
			theme.ApplicationID = app_id
			if j==0 {
				c1,c2,c3 := pkg.GetMonochromaticValueColor(blendColor)
				theme.PrimaryColor = c1.Hex()
				textColor := pkg.GetTextColor(c1)
				theme.TextColorAppBar = textColor.Hex()
				theme.SecondaryColor = c2.Hex()
				theme.BodyColor = "#FFFFFF"
				color, _ = colorful.Hex(theme.BodyColor)
				bodTextColor := pkg.GetTextColor(color)
				theme.TextColorBody = bodTextColor.Hex()
				theme.PrimaryDarkColor = c3.Hex()
				themes = append(themes,theme)
			}else if j==1{
				c1,c2,c3 := pkg.GetMonochromaticSaturationColor(blendColor)
				theme.PrimaryColor = c1.Hex()
				textColor := pkg.GetTextColor(c1)
				theme.TextColorAppBar = textColor.Hex()
				theme.SecondaryColor = c2.Hex()
				theme.BodyColor = "#FFFFFF"
				color, _ = colorful.Hex(theme.BodyColor)
				bodTextColor := pkg.GetTextColor(color)
				theme.TextColorBody = bodTextColor.Hex()
				theme.PrimaryDarkColor = c3.Hex()
				themes = append(themes,theme)

			}else if j==2 {
				c1,c2,c3 := pkg.GetTriadColor(blendColor)
				theme.PrimaryColor = c1.Hex()
				textColor := pkg.GetTextColor(c1)
				theme.TextColorAppBar = textColor.Hex()
				theme.SecondaryColor = c2.Hex()
				theme.BodyColor = "#FFFFFF"
				color, _ = colorful.Hex(theme.BodyColor)
				bodTextColor := pkg.GetTextColor(color)
				theme.TextColorBody = bodTextColor.Hex()
				theme.PrimaryDarkColor = c3.Hex()
				themes = append(themes,theme)
			}

		}
	}
	return themes
}
