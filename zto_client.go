package ztosdk

import (
	"time"

	"github.com/feeeei/ztosdk/common"

	"github.com/feeeei/ztosdk/common/base"
)

type ZTOClient struct {
	Host      string
	CompanyID string
	Key       *[]byte
	Partner   string
	Verify    string
	Debug     bool
}

func NewZTOClient(host, companyID, key, partner, verify string) *ZTOClient {
	k := []byte(key)
	debug := getEnvironment(host)

	return &ZTOClient{
		Host:      host,
		CompanyID: companyID,
		Partner:   partner,
		Verify:    verify,
		Key:       &k,
		Debug:     debug,
	}
}

func getEnvironment(host string) bool {
	return host != "http://japi.zto.cn/"
}

// SubmitOrderCode 创建电子运单
// 文档地址：https://zop.zto.com/apiDoc/  电子面单 -> 获取运单号（有密钥）
func (client *ZTOClient) SubmitOrderCode(content *common.ZTOContent) (*common.ZTOResponse, error) {
	if client.Debug {
		content.ID = "xfs101100111011"
	}
	now := base.Time(time.Now())
	request := &common.ZTOSubmitEncryptRequest{
		Partner:  client.Partner,
		Datetime: &now,
		Verify:   client.Verify,
		Content:  content,
	}
	sign, err := request.Sign(client.Key)
	if err != nil {
		return nil, err
	}
	return client.postRequest("submitOrderCode", sign, request)
}

// PartnerInsertSubmitagent 创建电子运单
// 文档地址：https://zop.zto.com/apiDoc/  电子面单 -> 获取单号（无秘钥）
func (client *ZTOClient) PartnerInsertSubmitagent(content *common.ZTOContent) (*common.ZTOResponse, error) {
	if client.Debug {
		content.Partner = "test"
		content.ID = "xfs101100111011"
	}
	request := &common.ZTOSubmitAgentRequest{
		CompanyID: client.CompanyID,
		MsgType:   "submitAgent",
		Data:      content,
	}
	sign, err := request.Sign(client.Key)
	if err != nil {
		return nil, err
	}
	return client.postRequest("partnerInsertSubmitagent", sign, request)
}
