# util 常用工具整理

## 安装
`go get github.com/darcyjoven/util`

## exel

## logger
```go
import "github.com/darcyjoven/util/logger"
...
    // 路径，日志文件名，文件后缀
    logger.Execute("./", "gomail", "log")
...
```
## mail
```go
import "github.com/darcyjoven/util/mail"

func main(){
    m := mail.NewMail(&EmailPara{
        // 主题
		Subject:    "无附件",
        // 正文
		Content:    []byte("无附件邮件"),
        // 不写也可以
		Types:      "text/plain;charset=utf-8",
        // 第一个为文件名，第二个是文件路径
		Attachment: map[string]string{"mail_test.go": "mail_test.go"},
        // 收件人，抄送，秘宋
		To:         []string{"darcy_joven@live.com"},
		Cc:         []string{"darcy_joven@live.com"},
		Bcc:        []string{"darcy_joven@live.com"},
        // 邮件服务器配置
        // 用户
		User:       "darcy_joven@live.com",
        // 邮件服务器地址
		Host:       "smtp.qiye.aliyun.com",
        // 端口
		Port:       "80",
        // 密码
		Password:   "aliyunxxx",
	})
	err := m.Send()
	if err != nil {
		log.Println(err)
		return
	}
}
```
## util
```go
import "github.com/darcyjoven/util"
...
// markdown 转为html，目前样式是默认的
util.MarkTotHtml(byte[]("# 111"))
...
// 数字转为字母 1->A,26->Z,52->BZ...
util.NumToLetter(12)
...
// 短日期  20220710
util.GetShortDate()
...
// 长日期20220710150405000
util.GetLongDate()
...
```