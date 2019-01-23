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
// 文档地址：https://zop.zto.com/apiDoc/  电子面单 -> 获取运单号
func (client *ZTOClient) SubmitOrderCode(r *common.ZTOContent) (*common.ZTOResponse, error) {
	if client.Debug {
		r.ID = "xfs101100111011"
	}
	now := base.Time(time.Now())
	request := common.ZTORequest{
		Partner:  client.Partner,
		Datetime: &now,
		Verify:   client.Verify,
		Content:  r,
	}
	if err := request.Sign(client.Key); err != nil {
		return nil, err
	}
	resp, err := client.postRequest("submitOrderCode", &request)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
