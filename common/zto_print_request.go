package common

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"net/url"
)

// doPrint 云打印-打印接口 request
type ZTOPrintRequest struct {
	PartnerCode         string               `json:"partnerCode"`
	PrintChannel        string               `json:"printChannel"`
	DeviceID            string               `json:"deviceId,omitempty"`
	QRCodeID            string               `json:"qrcodeId,omitempty"`
	PrinterID           string               `json:"printerId,omitempty"`
	PrintType           string               `json:"printType"`
	Repetition          bool                 `json:"repetition,omitempty"`
	PrintParam          *PrintParam          `json:"printParam"`
	Sender              *Sender              `json:"sender"`
	Receiver            *Receiver            `json:"receiver"`
	AppreciationService *AppreciationService `json:"appreciationService,omitempty"`
	PayType             string               `json:"payType,omitempty"`
	SiteName            string               `json:"siteName,omitempty"`
	Remark              string               `json:"remark,omitempty"`
}

type PrintParam struct {
	ParamType    string `json:"paramType"`
	MailNo       string `json:"mailNo,omitempty"`
	ElecAccount  string `json:"elecAccount,omitempty"`
	ElecPwd      string `json:"elecPwd,omitempty"`
	PrintMark    string `json:"printMark,omitempty"`
	PrintBagaddr string `json:"printBagaddr,omitempty"`
}

type AppreciationService struct {
	AppreciationType string `json:"appreciationType,omitempty"`
	CollectMoneytype int    `json:"collectMoneytype,omitempty"`
	GoodsName        string `json:"goodsName,omitempty"`
}

func (r *ZTOPrintRequest) Sign(key []byte) (string, error) {
	var buf bytes.Buffer
	body, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	_, err = buf.Write([]byte("request="))
	_, err = buf.Write(body)
	_, err = buf.Write(key)
	if err != nil {
		return "", err
	}
	md5 := md5.Sum(buf.Bytes())
	return base64.StdEncoding.EncodeToString(md5[:]), nil
}

func (r *ZTOPrintRequest) EncodeBody() (string, error) {
	jsonBody, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	values := url.Values{"request": []string{string(jsonBody)}}
	return values.Encode(), nil
}
