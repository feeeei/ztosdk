package ztosdk

import (
	"time"

	"github.com/feeeei/ztosdk/common"

	"github.com/feeeei/ztosdk/common/base"
)

type ZTOClient struct {
	Host      string
	CompanyID string
	Key       []byte
	Partner   string
	Debug     bool
}

func NewZTOClient(host, companyID, key, partner string) *ZTOClient {
	k := []byte(key)
	debug := getEnvironment(host)

	return &ZTOClient{
		Host:      host,
		CompanyID: companyID,
		Partner:   partner,
		Key:       k,
		Debug:     debug,
	}
}

func getEnvironment(host string) bool {
	return host != "http://japi.zto.cn/"
}

// SubmitOrderCode 创建电子运单
// 文档地址：https://zop.zto.com/apiDoc/  电子面单 -> 获取运单号（有密钥）
func (client *ZTOClient) SubmitOrderCode(content *common.ZTOContent, verify string) (*common.ZTOOrderResponse, error) {
	if client.Debug {
		content.ID = "xfs101100111011"
	}
	now := base.Time(time.Now())
	request := &common.ZTOSubmitEncryptRequest{
		Partner:  client.Partner,
		Datetime: &now,
		Verify:   verify,
		Content:  content,
	}
	sign, err := request.Sign(client.Key)
	if err != nil {
		return nil, err
	}
	return client.postOrderRequest("submitOrderCode", sign, request)
}

// PartnerInsertSubmitagent 创建电子运单
// 文档地址：https://zop.zto.com/apiDoc/  电子面单 -> 获取单号（无秘钥）
func (client *ZTOClient) PartnerInsertSubmitagent(content *common.ZTOContent) (*common.ZTOOrderResponse, error) {
	if content.Partner == "" {
		content.Partner = client.Partner
	}
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
	return client.postOrderRequest("partnerInsertSubmitagent", sign, request)
}

// UpdateOrders 预约寄件-批量更新订单（可用与批量取消订单）
// 文档地址：https://zop.zto.com/apiDoc/  订单服务 -> 预约寄件-订单取消
func (client *ZTOClient) UpdateOrders(orders []common.ZTOUpdateContent) ([]common.ZTOUpdateResponse, error) {
	request := &common.ZTOUpdateRequest{
		CompanyID: client.CompanyID,
		MsgType:   "UPDATE",
		Data:      orders,
	}
	sign, err := request.Sign(client.Key)
	if err != nil {
		return nil, err
	}
	return client.updateOrderRequest("commonOrderUpdate", sign, request)
}

// UpdateOrder 预约寄件-更新订单（可用于取消订单）
// 文档地址：https://zop.zto.com/apiDoc/  订单服务 -> 预约寄件-订单取消
func (client *ZTOClient) UpdateOrder(order *common.ZTOUpdateContent) (*common.ZTOUpdateResponse, error) {
	responses, err := client.UpdateOrders([]common.ZTOUpdateContent{*order})
	if err != nil {
		return nil, err
	}
	return &responses[0], err
}

// CancelOrder 预约寄件-取消订单
// 文档地址：https://zop.zto.com/apiDoc/  订单服务 -> 预约寄件-订单取消
func (client *ZTOClient) CancelOrder(orderCode, reason string) (*common.ZTOUpdateResponse, error) {
	return client.UpdateOrder(&common.ZTOUpdateContent{
		OrderCode:  orderCode,
		FieldName:  "status",
		FieldValue: "cancel",
		Reason:     reason,
	})
}

// DoPrint 云打印-打印接口
// 文档地址：https://zop.zto.com/apiDoc/  电子面单 -> 云打印-打印接口
func (client *ZTOClient) DoPrint(request *common.ZTOPrintRequest) (*common.ZTOPrintResponse, error) {
	sign, err := request.Sign(client.Key)
	if err != nil {
		return nil, err
	}
	return client.postPrintRequest("doPrint", sign, request)
}

// TraceInterfaceNewTraces 快件轨迹-获取快件轨迹信息
// 文档地址：https://zop.zto.com/apiDoc/  快件轨迹 ->获取快件轨迹信息
func (client *ZTOClient) TraceInterfaceNewTraces(billCodes []string) ([]common.ZTOTraceResponse, error) {
	request := &common.ZTOTracesRequest{
		CompanyID: client.CompanyID,
		MsgType:   "NEW_TRACES",
		Data:      billCodes,
	}
	sign, err := request.Sign(client.Key)
	if err != nil {
		return nil, err
	}
	return client.postTraceInterfaceNewTraces("traceInterfaceLatest", sign, request)
}

func (client *ZTOClient) TraceInterfaceLatest(billCodes []string) ([]common.ZTOLastTraceResponse, error) {
	request := &common.ZTOTracesRequest{
		CompanyID: client.CompanyID,
		MsgType:   "LATEST",
		Data:      billCodes,
	}
	sign, err := request.Sign(client.Key)
	if err != nil {
		return nil, err
	}
	return client.postTraceInterfaceLatest("traceInterfaceLatest", sign, request)
}
