package kurs

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/wendylau87/xfers2021/entities"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func(d *domain) callBCA()([]entities.CreateKurs, error){
	results := []entities.CreateKurs{}
	resp, err := http.Get("https://www.bca.co.id/id/informasi/kurs")
	if err != nil {
		d.logger.LogError("error crawl BCA because %s", err)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		d.logger.LogError("error read body because %s", err)
	}

	doc.Find(".m-table-kurs").Children().Each(func(i int, sel *goquery.Selection) {
		sel.Find("tbody tr").Each(func (k int, tr *goquery.Selection){
			result := entities.CreateKurs{
				ValidDate: time.Now().Format("2006-01-02"),
			}
			tr.Find("td").Each(func (j int, td *goquery.Selection){
				valString := strings.ReplaceAll(strings.ReplaceAll(td.Find("p").Text(), ".",""),",",".")
				if valString != ""{
					valFloat, err := strconv.ParseFloat(valString,64)
					if err != nil{
						result.Name = valString
					}else{
						rateType, _ := td.Find("p").Attr("rate-type")
						if rateType == "ERate-buy"{
							result.ERate.Buy = valFloat
						}else if rateType == "ERate-sell"{
							result.ERate.Sell = valFloat
						}else if rateType == "TT-buy"{
							result.TTCounter.Buy = valFloat
						}else if rateType == "TT-sell"{
							result.TTCounter.Sell = valFloat
						}else if rateType == "BN-buy"{
							result.BankNote.Buy = valFloat
						}else if rateType == "BN-sell"{
							result.BankNote.Sell = valFloat
						}
					}
				}
			})
			results = append(results, result)
		})
	})

	return results, nil
}
