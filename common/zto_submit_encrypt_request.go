package common

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"net/url"

	"github.com/feeeei/ztosdk/common/base"
)

// ZTOSubmitRequest 获取运单号(有密钥) request
type ZTOSubmitEncryptRequest struct {
	Partner  string      `json:"partner,omitempty"`
	Datetime *base.Time  `json:"datetime,omitempty"`
	Verify   string      `json:"verify,omitempty"`
	Content  *ZTOContent `json:"content,omitempty"`
}

func (r *ZTOSubmitEncryptRequest) Sign(key *[]byte) (string, error) {
	var buf bytes.Buffer
	body, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	_, err = buf.Write([]byte("data="))
	_, err = buf.Write(body)
	_, err = buf.Write(*key)
	if err != nil {
		return "", err
	}
	md5 := md5.Sum(buf.Bytes())
	return base64.StdEncoding.EncodeToString(md5[:]), nil
}

func (r *ZTOSubmitEncryptRequest) EncodeBody() (string, error) {
	transportToSimpleAddress(r.Content.Receiver)
	transportToSimpleAddress(r.Content.Sender)
	jsonBody, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	values := url.Values{"data": []string{string(jsonBody)}}
	return values.Encode(), nil
}
