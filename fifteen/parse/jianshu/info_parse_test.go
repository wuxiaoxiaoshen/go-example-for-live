package jianshu

import (
	"go-example-for-live/fifteen/download"
	"testing"

	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func TestInfoParse(test *testing.T) {

	docJianshu, _ := download.GetHtmlResponse("https://www.jianshu.com/u/58f0817209aa")

	tt := []struct {
		doc *goquery.Document
	}{
		{
			doc: docJianshu,
		},
	}

	for _, t := range tt {
		fmt.Println(InfoParse(t.doc))
	}

}
