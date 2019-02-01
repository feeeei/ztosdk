package common

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
)

type ZTOUpdateRequest struct {
	CompanyID string              `json:"company_id"`
	MsgType   string              `json:"msg_type"`
	Data      *[]ZTOUpdateContent `json:"data"`
}

type ZTOUpdateContent struct {
	OrderCode  string `json:"orderCode"`
	FieldName  string `json:"fieldName"`
	FieldValue string `json:"fieldValue"`
	Reason     string `json:"reason,omitempty"`
}

func (r *ZTOUpdateRequest) Sign(key *[]byte) (string, error) {
	var buf bytes.Buffer
	body, err := json.Marshal(r.Data)
	if err != nil {
		return "", err
	}
	raw := fmt.Sprintf("company_id=%s&data=%s&msg_type=%s", r.CompanyID, body, r.MsgType)
	_, err = buf.Write([]byte(raw))
	_, err = buf.Write(*key)
	if err != nil {
		return "", err
	}
	md5 := md5.Sum(buf.Bytes())
	return base64.StdEncoding.EncodeToString(md5[:]), nil
}

func (r *ZTOUpdateRequest) EncodeBody() (string, error) {
	jsonBody, err := json.Marshal(r.Data)
	if err != nil {
		return "", err
	}

	values := url.Values{
		"company_id": []string{r.CompanyID},
		"msg_type":   []string{r.MsgType},
		"data":       []string{string(jsonBody)},
	}
	return values.Encode(), nil
}
