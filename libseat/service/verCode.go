// //接口类型：互亿无线触发短信接口，支持发送验证码短信、订单通知短信等。
// // 账户注册：请通过该地址开通账户http://user.ihuyi.com/register.html
// // 注意事项：
// //（1）调试期间，请使用用系统默认的短信内容：您的验证码是：【变量】。请不要把验证码泄露给其他人。
// //（2）请使用 APIID 及 APIKEY来调用接口，可在会员中心获取；
// //（3）该代码仅供接入互亿无线短信接口参考使用，客户可根据实际需要自行编写；

// // APIID：C71073387

// // APIKEY：c0862d55874ad83c1545411093728871

package service

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/logs"
)

func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
func VerCode(codeNumber string, phonenumber string) {
	v := url.Values{}
	_now := strconv.FormatInt(time.Now().Unix(), 10)
	//fmt.Printf(_now)
	_account := "XXXXXXXXXXXXX"  //查看用户名 登录用户中心->验证码通知短信>产品总览->API接口信息->APIID
	_password := "XXXXXXXXXXXXX" //查看密码 登录用户中心->验证码通知短信>产品总览->API接口信息->APIKEY
	_mobile := phonenumber
	_content := fmt.Sprintf("您的验证码是:%s。", codeNumber)
	v.Set("account", _account)
	v.Set("password", GetMd5String(_account+_password+_mobile+_content+_now))
	v.Set("mobile", _mobile)
	v.Set("content", _content)
	v.Set("time", _now)
	body := strings.NewReader(v.Encode()) //把form数据编下码
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "http://106.ihuyi.com/webservice/sms.php?method=Submit&format=json", body)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	//fmt.Printf("%+v\n", req) //看下发送的结构

	resp, err := client.Do(req) //发送
	defer resp.Body.Close()     //一定要关闭resp.Body
	data, _ := ioutil.ReadAll(resp.Body)
	logs.Info(string(data), err)
}

// 验证码要在redis中保存一分钟,具体来说是80s,扣除一般网络延迟
func GetCode() (codeNumber int) {
	rand.Seed(time.Now().UnixNano())
	codeNumber = rand.Intn(8999) + 1000
	fmt.Println(fmt.Sprintf("您的验证码是:%d。", codeNumber))
	// // 调用发送短信
	// verCode(codeNumber, phonenumber)
	return
}
