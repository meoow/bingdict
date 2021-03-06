package main

import "github.com/meoow/nodefinder"
import "fmt"
import "net/http"
import "log"
import "strings"

const (
	bing_uri    string = "http://cn.bing.com/dict/search?q=%s&go=&qs=bs&form=CM&mkt=zh-CN&setlang=ZH"
	userAgent   string = "Mozilla/5.0 (iPod; CPU iPhone OS 6_1_6 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10B500 Safari/8536.25"
	language    string = "en-us,en;q=0.5"
	charset     string = "ISO-8859-1,utf-8;q=0.7,*;q=0.7"
	format      string = "text/xml,application/xml,application/xhtml+xml,text/html;q=0.9,text/plain;q=0.8,image/png,*/*;q=0.5"
	content_enc string = "gzip, deflate"
)

var (
	qdefPath = nodefinder.NewPath("div.qdef")
	prPath   = nodefinder.NewPath("div.hd_pr")
	prusPath = nodefinder.NewPath("div.hd_prUS")
	listPath = nodefinder.NewPath("ul/li")
)

func main() {
	word := getText()

	reqst, _ := http.NewRequest("GET", fmt.Sprintf(bing_uri, word), nil)
	reqst.Header.Add("User-Agent", userAgent)
	reqst.Header.Add("Accept-Language", language)
	reqst.Header.Add("Accept-Charset", charset)
	reqst.Header.Add("Content-Encoding", content_enc)

	client := &http.Client{}

	resp, err := client.Do(reqst)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	defNode, err := nodefinder.Find(qdefPath, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if len(defNode) == 0 {
		return
	}

	pr1Nodes := nodefinder.FindByNode(prusPath, defNode[0])
	if len(pr1Nodes) > 0 && strings.HasSuffix(pr1Nodes[0].FirstChild.Data, "] ") {
		fmt.Println(pr1Nodes[0].FirstChild.Data)
	}

	pr2Nodes := nodefinder.FindByNode(prPath, defNode[0])
	if len(pr2Nodes) > 0 && strings.HasSuffix(pr2Nodes[0].FirstChild.Data, "] ") {
		fmt.Println(pr2Nodes[0].FirstChild.Data)
	}

	fmt.Println("")
	defsNodes := nodefinder.FindByNode(listPath, defNode[0])
	for _, dn := range defsNodes {
		fmt.Printf("[%s] %s\n",
			dn.FirstChild.FirstChild.Data,
			dn.LastChild.FirstChild.FirstChild.Data)
	}

}
