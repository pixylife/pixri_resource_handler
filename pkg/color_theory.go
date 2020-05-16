package pkg

import (
	"fmt"
	"sort"
)

type Keyword struct {
	Key string
	Value float64
}
type KeywordList []Keyword

func GetMaxKeywords(m *KeywordList) *KeywordList{
	sort.Sort(sort.Reverse(m))
	fmt.Println(m)
	return m
}


func (p KeywordList) Len() int { return len(p) }
func (p KeywordList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p KeywordList) Swap(i, j int){ p[i], p[j] = p[j], p[i] }

