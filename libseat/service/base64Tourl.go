package service

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/astaxie/beego/logs"
	uuid "github.com/satori/go.uuid"
)

type urlMsg struct {
	Url         string
	AccessToken string `json:"access_token"`
	Owner       string
	Repo        string
	Path        string
	Content     string `json:"content"`
	Message     string `json:"message"`
}

func Base64Tourl(image string) (url string) {
	//初始化
	var urlmsg = &urlMsg{
		AccessToken: "71c71d8dab23658fa61ee3a95bbe44db",

		Owner: "itanrong",

		Repo: "picstore",

		Message: "add a picture",

		Content: base64.StdEncoding.EncodeToString([]byte(image)),
	}
	getpath(urlmsg)
	urlmsg.Url = fmt.Sprintf("https://gitee.com/api/v5/repos/%s/%s/contents/%s", urlmsg.Owner, urlmsg.Repo, urlmsg.Path)

	_, err := giteeUpload(urlmsg)
	if err != nil {
		logs.Error("gitee post request err")
		return ""
	}
	// fmt.Println(string(body))
	return getDownloadUrl(urlmsg)
}

// 获得path
func getpath(u *urlMsg) {
	uUID := uuid.NewV4()
	t := time.Now().Format("2006010215") //year month day hour
	u.Path = fmt.Sprintf("%v/%v%s", t, uUID, "output.png")
}

//post请求
func giteeUpload(u *urlMsg) (bodyText []byte, err error) {
	client := &http.Client{}
	header := urlMsg{
		AccessToken: u.AccessToken,
		Content:     u.Content,
		Message:     u.Message,
	}
	jsons, err := json.Marshal(header) //转换成JSON返回的是byte[]
	if err != nil {
		log.Fatal(err)
	}
	var data = strings.NewReader(string(jsons))
	req, err := http.NewRequest("POST", u.Url, data)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err = ioutil.ReadAll(resp.Body) //byte[]
	defer resp.Body.Close()
	if err != nil {
		return bodyText, err
	}
	// fmt.Println(bodyText)
	return bodyText, nil
}

func getDownloadUrl(u *urlMsg) (downloadeurl string) {
	return fmt.Sprintf("https://gitee.com/%s/%s/raw/master/%s", u.Owner, u.Repo, u.Path)
}
