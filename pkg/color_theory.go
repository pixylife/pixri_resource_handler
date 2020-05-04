package pkg

import "sort"

type Keyword struct {
	Key string
	Value float64
}
type KeywordList []Keyword

func GetMaxKeywords(m map[string]float64) KeywordList{
	pl := make(KeywordList, len(m))
	i := 0
	for k, v := range m {
		pl[i] = Keyword{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}


func (p KeywordList) Len() int { return len(p) }
func (p KeywordList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p KeywordList) Swap(i, j int){ p[i], p[j] = p[j], p[i] }

