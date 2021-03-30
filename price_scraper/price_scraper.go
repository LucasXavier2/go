package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func fetchUrl(url string) string {
	fmt.Println("Loading url...")
	resp, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	return string(body)
}

func getAmazonPriceFromHTML(productPageHTML string) string {
	var expression string = `id="priceblock_ourprice".*>(.*?)(</span>)`
	re, err := regexp.Compile(expression)
	var prices []string

	if err != nil {
		fmt.Println("There was an error compiling your regex expression")
		log.Fatalln(err)
	}

	prices = re.FindStringSubmatch(productPageHTML)

	if len(prices) > 1 {
		return prices[1]
	}
	return "No price found"
}

func main() {

	var amazonProductUrl string = "https://www.amazon.com.br/480Gb-SanDisk-Armazenamento-Interno-Preto/dp/B01F9G46Q8/ref=sr_1_6?__mk_pt_BR=%C3%85M%C3%85%C5%BD%C3%95%C3%91&dchild=1&keywords=ssd+480gb&qid=1617123616&s=electronics&sr=1-6"
	var responseText string = fetchUrl(amazonProductUrl)

	amazonPrice := getAmazonPriceFromHTML(responseText)
	fmt.Printf("Product price: %s", amazonPrice)

}
