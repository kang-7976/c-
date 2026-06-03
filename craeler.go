package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func main() {
	var scanner *bufio.Scanner
	scanner = bufio.NewScanner(os.Stdin)
	var prompt string
	prompt = "输入关键词："
	fmt.Print(prompt)

	var scanRes bool
	scanRes = scanner.Scan()
	if scanRes == false {
		return
	}
	var keyword string
	keyword = scanner.Text()
	if keyword == "" {
		fmt.Println("关键词不能为空")
		return
	}

	var encodeKeyword string
	encodeKeyword = url.QueryEscape(keyword)
	var rawBaidu string
	rawBaidu = "https://www.baidu.com/s?wd=" + encodeKeyword
	var rawGithub string
	rawGithub = "https://github.com/search?q=" + encodeKeyword
	var ua string
	ua = "Mozilla/5.0 (Android; Termux)"
	var client http.Client
	var req *http.Request
	var err error
	var resp *http.Response
	var body []byte

	// 百度校验+爬取
	fmt.Println("\n---校验百度链接---")
	req, err = http.NewRequest(http.MethodGet, rawBaidu, nil)
	if err == nil {
		req.Header.Set("User-Agent", ua)
		resp, err = client.Do(req)
		if err == nil {
			defer resp.Body.Close()
			// 状态码200才算页面真实存在
			if resp.StatusCode == http.StatusOK {
				body, err = io.ReadAll(resp.Body)
				if err == nil {
					fmt.Println("✅ 链接有效：", rawBaidu)
					fmt.Println("页面片段：")
					var cut int
					cut = 700
					if len(body) < cut {
						cut = len(body)
					}
					fmt.Println(string(body[:cut]))
				}
			} else {
				fmt.Println("❌ 关键词无匹配内容，链接无效，不展示地址")
			}
		} else {
			fmt.Println("❌ 百度网络请求失败")
		}
	}

	// GitHub校验+爬取
	fmt.Println("\n---校验Github链接---")
	req, err = http.NewRequest(http.MethodGet, rawGithub, nil)
	if err == nil {
		req.Header.Set("User-Agent", ua)
		resp, err = client.Do(req)
		if err == nil {
			defer resp.Body.Close()
			if resp.StatusCode == http.StatusOK {
				body, err = io.ReadAll(resp.Body)
				if err == nil {
					fmt.Println("✅ 链接有效：", rawGithub)
					fmt.Println("页面片段：")
					var cut int
					cut = 700
					if len(body) < cut {
						cut = len(body)
					}
					fmt.Println(string(body[:cut]))
				}
			} else {
				fmt.Println("❌ 关键词无匹配内容，链接无效，不展示地址")
			}
		} else {
			fmt.Println("❌ Github网络请求失败")
		}
	}
}

