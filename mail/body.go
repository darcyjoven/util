package mail

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

func (e *email) setBody() {
	buffer := bytes.NewBuffer(nil)
	boundary := "GoBoundary"
	//  header 设置
	header := make(map[string]string)
	header["From"] = e.server.user
	header["To"] = strings.Join(e.to, ";")
	header["Cc"] = strings.Join(e.cc, ";")
	// header["Bcc"] = strings.Join(e.bcc, ";")
	// utf8 base64 编码邮件主题
	header["Subject"] = "=?UTF-8?B?" + base64.StdEncoding.EncodeToString([]byte(e.subject)) + "?="
	header["Content-Type"] = "multipart/mixed;boundary=" + boundary
	header["Mime-Version"] = "1.0"
	header["Date"] = time.Now().String()
	err := writerHeader(buffer, header)
	if err != nil {
		log.Println(err)
		return
	}

	// 正文头部内容
	body := "\r\n--" + boundary + "\r\n"
	body += "Content-Type:" + e.body.types + "\r\n"
	body += "Content-Transfer-Encoding:base64\r\n"
	body += "\r\n" + base64.StdEncoding.EncodeToString(e.body.content) + "\r\n"
	// body += "\r\n" + string(e.body.content) + "\r\n"
	buffer.WriteString(body)

	// 附件
	for k, v := range e.attachment {
		attachment := "\r\n--" + boundary + "\r\n"
		attachment += "Content-Transfer-Encoding:base64\r\n"
		attachment += "Content-Disposition:attachment\r\n"
		// 附件名称也用base64编码，防止中文乱码

		attachment += "Content-Type:application/octet-stream;name=\"=?UTF-8?B?" + base64.StdEncoding.EncodeToString([]byte(k)) + "?=\"\r\n"
		buffer.WriteString(attachment)

		defer func() {
			if err := recover(); err != nil {
				log.Fatalln(err)
			}
		}()
		writeFile(buffer, v)
	}

	buffer.WriteString("\r\n--" + boundary + "--")
	e.body.content = buffer.Bytes()
}

// 正文头部信息
func writerHeader(buffer *bytes.Buffer, header map[string]string) error {
	h := ""
	for k, v := range header {
		h += k + ":" + v + "\r\n"
	}
	h += "\r\n"
	_, err := buffer.WriteString(h)
	return err
}

// 附件内容写入
func writeFile(buffer *bytes.Buffer, path string) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println(err)
		return
	}
	payload := make([]byte, base64.StdEncoding.EncodedLen(len(file)))
	base64.StdEncoding.Encode(payload, file)
	buffer.WriteString("\r\n")

	for index, line := 0, len(payload); index < line; index++ {
		buffer.WriteByte(payload[index])
		if (index+1)%76 == 0 {
			buffer.WriteString("\r\n")
		}
	}
}
