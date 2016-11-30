# http://www.cnblogs.com/nima/p/6114613.html

taobao.go

```
package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/hunterhug/go_tool/spider"
	"github.com/hunterhug/go_tool/spider/query"
	"github.com/hunterhug/go_tool/util"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(`��ӭʹ���Ա���èͼƬ����С���ߣ���ͬ��Ŀ¼д�����ӽ�taobao.txt������EXE����`)
	fmt.Println("�����磺tmall.com/item.htm?id=523350171126&skuId=3120562159704,tmall")
	fmt.Println("---------------��������ҳ��ͼƬ�ᱣ����tmallĿ¼-----------------------")
	c, e := util.ReadfromFile("./taobao.txt")
	if e != nil {
		fmt.Println("��taobao.txt����")
	} else {
		urls := strings.Split(string(c), "\n")
		for _, url := range urls {
			url := strings.Replace(strings.TrimSpace(url), "\r", "", -1)
			downlod(url)
		}

	}
	fmt.Println("���ֶ��ر�ѡ��...")
	util.Sleep(100)
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
		dudu := "detail.tmall.com"
		if strings.Contains(url, "item.taobao.com") {
			dudu = "item.taobao.com"
		}
		s.NewHeader("Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.99 Safari/537.36", dudu, nil)
		content, err := s.Get()
		if err != nil {

		} else {
			//fmt.Println(string(content))
			docm, err := query.QueryBytes(content)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				//fmt.Println(string(content))
				docm.Find("img").Each(func(num int, node *goquery.Selection) {
					img, e := node.Attr("src")
					if e == false {
						img, e = node.Attr("data-src")
					}
					if e && img != "" {
						if strings.Contains(img, ".gif") {
							return
						}
						fmt.Println("ԭʼ�ļ���" + img)
						r, _ := regexp.Compile(`([\d]{1,4}x[\d]{1,4})`)
						imgdudu := r.FindStringSubmatch(img)
						sizes := "720*720"
						if len(imgdudu) == 2 {
							sizes = imgdudu[1]
						}
						temp := strings.Replace(img, sizes, "720x720", -1)
						filename := md55(temp)
						if util.FileExist(dir + "/" + filename + ".jpg") {
							fmt.Println("�ļ����ڣ�" + dir + "/" + filename)
						} else {
							fmt.Println("����:" + temp)
							s.Url = "http:" + temp
							imgsrc, e := s.Get()
							if e != nil {
								fmt.Println("���س���" + temp + ":" + e.Error())
								return
							}
							e = util.SaveToFile(dir+"/"+filename+".jpg", imgsrc)
							if e == nil {
								fmt.Println("�ɹ�������" + dir + "/" + filename)
							}
							util.Sleep(2)
							fmt.Println("��ͣ����")
						}
					}
				})

			}

		}
	}

}


```


��Դ��ͬ��Ŀ¼д��taobao.txt��


```
https://detail.tmall.com/item.htm?id=523350171126&skuId=3120562159704,myword
```

ͼƬ���ᱣ����myword����


���Ȱ�װ��

```
go get -v github.com/PuerkitoBio/goquery
go get -v github.com/hunterhug/go_tool
```

Ȼ���ܣ�

```
go run taobao.go
```

������鷳

�뵽�������ش��exeִ���ļ���

http://pan.baidu.com/s/1jHKUGZG


����goĿ¼������taobao.rar

Դ���ڣ�


```
https://github.com/hunterhug/taobao_img
```


��ͼ���£�

![](http://images2015.cnblogs.com/blog/672593/201611/672593-20161130115120865-468057454.png)
![](http://images2015.cnblogs.com/blog/672593/201611/672593-20161130115133177-1022436365.png)



����150�ֵ���ʲ�����������ҳ��ѡ��
����150�ֵ���ʲ�����������ҳ��ѡ��
����150�ֵ���ʲ�����������ҳ��ѡ��
����150�ֵ���ʲ�����������ҳ��ѡ��
����150�ֵ���ʲ�����������ҳ��ѡ��
����150�ֵ���ʲ�����������ҳ��ѡ��
����150�ֵ���ʲ�����������ҳ��ѡ��
����150�ֵ���ʲ�����������ҳ��ѡ��
����150�ֵ���ʲ�����������ҳ��ѡ��
����150�ֵ���ʲ�����������ҳ��ѡ��
����150�ֵ���ʲ�����������ҳ��ѡ��
����150�ֵ���ʲ�����������ҳ��ѡ��
����150�ֵ���ʲ�����������ҳ��ѡ��
����150�ֵ���ʲ�����������ҳ��ѡ��