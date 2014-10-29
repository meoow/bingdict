package main

import "github.com/meoow/nodefinder"
import "fmt"
import "net/http"
import "log"

const (
	bing_uri  string = "http://cn.bing.com/dict/search?q=%s&go=&qs=bs&form=CM&mkt=zh-CN&setlang=ZH"
	userAgent string = "Mozilla/5.0 (iPod; CPU iPhone OS 6_1_6 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10B500 Safari/8536.25"
	language  string = "en-us,en;q=0.5"
	charset   string = "ISO-8859-1,utf-8;q=0.7,*;q=0.7"
	format    string = "text/xml,application/xml,application/xhtml+xml,text/html;q=0.9,text/plain;q=0.8,image/png,*/*;q=0.5"
)

func main() {
	word := getText()

	reqst, _ := http.NewRequest("GET", fmt.Sprintf(bing_uri, word), nil)
	reqst.Header.Add("User-Agent", userAgent)
	reqst.Header.Add("Accept-Language", language)
	reqst.Header.Add("Accept-Charset", charset)

	client := &http.Client{}

	resp, err := client.Do(reqst)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	defNode, err := nodefinder.Find(
		nodefinder.NewPath("div.qdef"), resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if len(defNode) == 0 {
		return
	}

	pr1Nodes := nodefinder.FindByNode(
		nodefinder.NewPath("div.hd_prUS"), defNode[0])
	pr2Nodes := nodefinder.FindByNode(
		nodefinder.NewPath("div.hd_pr"), defNode[0])
	for _, pn := range pr1Nodes {
		fmt.Println(pn.FirstChild.Data)
	}
	for _, pn := range pr2Nodes {
		fmt.Println(pn.FirstChild.Data)
	}

	fmt.Println("")
	defsNodes := nodefinder.FindByNode(nodefinder.NewPath("ul/li"), defNode[0])
	for _, dn := range defsNodes {
		fmt.Printf("[%s] %s\n",
			dn.FirstChild.FirstChild.Data,
			dn.LastChild.FirstChild.FirstChild.Data)
	}

}
