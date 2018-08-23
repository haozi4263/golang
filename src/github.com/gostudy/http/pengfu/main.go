package main

import (
	"fmt"
	"regexp"  //正则表达式
	"strconv"
	"net/http"
	"strings"
	"os"
)

//把内容写入到文件
func StoreJoyToFile(i int, fileTitle []string, fileContent []string)  {
	//新建文件
	f, err := os.Create(strconv.Itoa(i)+".txt")
	if err != nil {
		fmt.Println("os.Create = err", err)
		return
	}
	defer f.Close()

	//写内容
	n := len(fileTitle)
	for i := 0; i < n; i++ {
		//写标题
		f.WriteString(fileTitle[i]+"\n")
		//写内容
		f.WriteString(fileContent[i]+"\n")

		f.WriteString("\n================================================================================================================================================================\n")
	}
}



func HttpGet(url string) (result string, err error)  {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()

	//读取网页内容
	buf := make([]byte, 4*1024)
	for {
		n ,_ := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		result += string(buf[:n])	//累加读取的内容
	}
	return
}

func SpiderOneJoy(url string) (title, content string, err error)  {
	//开始爬取页面内容
	result, err1 := HttpGet(url)
	if err1 != nil {
		err = err1
		return
	}

	//取关键信息
	//取标题  <h1>  标题  </h1>  只取1个
	re1 := regexp.MustCompile(`<h1>(?s:(.*?))</h1>`)
	if re1 == nil {
		//fmt.Println("regexp.MustCompile err")
		err = fmt.Errorf("regexp.MustCompile err")
		return
	}

	//取内容
	tmpTitle := re1.FindAllStringSubmatch(result, 1) //最后一个参数为1，只过滤第一个
	for _, data := range tmpTitle {
		title = data[1]
		title = strings.Replace(title, "\r", "", -1)//旧的字符串替换成新的字符串, 将\n换成空字符
		title = strings.Replace(title, "\n", "", -1)
		title = strings.Replace(title, " ", "", -1)  //空格替换空字符
		title = strings.Replace(title, "\t", "", -1)  //替换tab健
		break
	}

	//取内容 <div class="content-txt pt10"> 段子内容 <a id="prev" href=">
	re2 := regexp.MustCompile(`<div class="content-txt pt10">(?s:(.*?))<a id="prev" href="`)
	if re2 == nil {
		//fmt.Println("regexp.MustCompile err")
		err = fmt.Errorf("regexp.MustCompile err")
		return
	}

	//取内容
	tmpContent := re2.FindAllStringSubmatch(result, -1)
	for _, data := range tmpContent {
		content = data[1]
		content = strings.Replace(content, "\t", "", -1)
		content = strings.Replace(content, "\n", "", -1)
		content = strings.Replace(content, "\r", "", -1)
		content = strings.Replace(content, "\t", "", -1)
		content = strings.Replace(content, "<br />", "", -1)
		content = strings.Replace(content, "<br>", "", -1)
		content = strings.Replace(content, "&nbsp", "", -1)
		content = strings.Replace(content, "&ldquo;", "", -1)
		content = strings.Replace(content, "&rdquo;", "", -1)
		content = strings.Replace(content, "</p>", "", -1)
		content = strings.Replace(content, "<p>", "", -1)
		content = strings.Replace(content, "&hellip;", "", -1)
		content = strings.Replace(content, "; ; ;", "", -1)


		break
	}
	return
}

func SpiderPage(i int)  {
	//明确爬取的网址
	url := "https://www.pengfu.com/xiaohua_" + strconv.Itoa(i) + ".html"
	fmt.Printf("正在爬取第%d个网页： %s\n", i, url)

	//开始爬取页面内容
	result, err := HttpGet(url)
	if err != nil {
		fmt.Println("HttpGet err =",  err)
		return
	}
	//fmt.Println("r = ",result) 调试
	//取内容 <h1 class="dp-b"><a herf=" 一个段子url连接   "
	//解析表达式
	re := regexp.MustCompile(`<h1 class="dp-b"><a href="(?s:(.*?))"`)
	if re == nil {
		fmt.Println("regexp.MustCompile err")
		return
	}
	// 取关键信息
	joyUrls := re.FindAllStringSubmatch(result, -1) //-1取所有内容
	//fmt.Println("joyUrls = ", joyUrls)  //调试打印网址


	fileTite := make([]string, 0)
	fileContent := make([]string, 0)

	//取第二网址,第一个返回下标，第二个返回内容
	for _, data := range joyUrls {
		fmt.Println("url = ", data[1])//爬取第二个有用网址
		//爬取每一个笑话里面的段子
		title, content, err := SpiderOneJoy(data[1])  //返回标题，内容，错误
		if err != nil{
			fmt.Println("SpiderOneJoy err = ", err)
		}
		//fmt.Printf("title = #%v#", title)
		//fmt.Printf("contend = #%v#", content)

		fileTite = append(fileTite, title)  //追加内容
		fileContent = append(fileContent, content)//追加内容
	}
	fmt.Println("fileTile= ", fileTite)  //打印内存
	fmt.Println("fileContent= ", fileContent) //打印内容

	//把内容写入到文件
	StoreJoyToFile(i, fileTite, fileContent)

}


func DoWork(start, end int)  {
	fmt.Printf("准备爬取第%d页到%d页的网址\n", start, end)

	for i := start; i <= end; i++ {
		//定义一个函数，爬主页面
		SpiderPage(i)
	}
}


func main()  {
	var start, end int
	fmt.Printf("请输入起始页( >= 1) :")
	fmt.Scan(&start)
	fmt.Printf("请输入终止页( >= 起始页) :")
	fmt.Scan(&end)

	DoWork(start, end)  //工作函数
}
