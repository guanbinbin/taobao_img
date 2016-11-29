package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/hunterhug/go_tool/spider"
	"github.com/hunterhug/go_tool/spider/query"
	"github.com/hunterhug/go_tool/util"
	"strings"
)

func main() {
	c, e := util.ReadfromFile("./taobao.txt")
	if e != nil {
		fmt.Println("打开taobao.txt出错")
	} else {
		urls := strings.Split(string(c), "\n")
		for _, url := range urls {
			url := strings.Replace(strings.TrimSpace(url), "\r", "", -1)
			downlod(url)
		}

	}
}

func md55(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	rs := hex.EncodeToString(h.Sum(nil))
	return rs
}

func downlod(urlmany string) {
	temp := strings.Split(urlmany, ",")
	url := temp[0]
	filename := util.TodayString(3)
	if len(temp) >= 2 {
		filename = temp[1]
	}
	dir := "./" + filename
	util.MakeDir(dir)
	s, e := spider.NewSpider(nil)
	if e != nil {

	} else {
		s.Url = url
		s.NewHeader("Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.99 Safari/537.36", "detail.tmall.com", nil)
		content, err := s.Get()
		if err != nil {

		} else {
			//fmt.Println(string(content))
			docm, err := query.QueryBytes(content)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				docm.Find("img").Each(func(num int, node *goquery.Selection) {
					util.Sleep(2)
					fmt.Println("暂停两秒")
					img, e := node.Attr("src")
					if e && img != "" {
						temp := strings.Replace(img, "60x60", "720x720", -1)
						temp = strings.Replace(temp, "430x430", "720x720", -1)
						fmt.Println("下载" + temp)
						s.Url = "http:" + temp
						imgsrc, e := s.Get()
						if e != nil {
							fmt.Println("下载出错" + temp + ":" + e.Error())
						}
						filename := md55(temp)
						if util.FileExist(dir + "/" + filename + ".jpg") {
							fmt.Println("文件存在：" + dir + "/" + filename)
						} else {
							e = util.SaveToFile(dir+"/"+filename+".jpg", imgsrc)
							if e == nil {
								fmt.Println("成功保存在" + dir + "/" + filename)
							}
						}
					}
				})

			}

		}
	}

}
