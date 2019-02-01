## 中通开放平台SDK for Go

官方 API 文档地址：[https://zop.zto.com/apiDoc/](https://zop.zto.com/apiDoc/)

## 实现接口列表
 - 获取运单号（有秘钥）submitOrderCode
 - 获取单号（无密钥） partnerInsertSubmitagent
 - 云打印-打印接口 doPrint
 - 预约寄件-订单取消 commonOrderUpdate

## 使用示例
1. 引入

   ```go
   go get github.com/feeeei/ztosdk
   ```
2. New Client
   ```go
   // 测试环境参数
   // 不同业务接口有不同的测试参数，请注意区分
   // 如果某些业务仅需其中几个字段，其它字段请传递""
   
   // 下述参数为 获取运单号（有秘钥） submitOrderCode 测试参数
   host      := "http://58.40.16.120:9001/"
   companyID := "kfpttestCode"
   key       := "kfpttestkey=="
   partner   := "test"
   verify    := "ZTO123"
 
   client := ztosdk.NewZTOClient(host, companyID, key, partner, verify)
   ```
1. 实现业务
   ##### 电子面单 -> 获取单号(有密钥)
   ```go
   ztoContent := &common.ZTOContent{
	  ID:      "id",
	  TypeID:  "1",
	  TradeID: "tradeid",
	  Sender: &common.Sender{
	    Name:    "发件人姓名",
	    Company: "发件人公司（可不填）",
	    Mobile:  "18xxxxxxxxx",
	    City:    "北京市,北京市,xx区",
	    Address: "详细地址",
	    ZIPCode: "邮编（可不填）",
	    Email:   "发件人邮箱（可不填）",
	  },
	  Receiver: &common.Receiver{
	    Name:    "收件人姓名",
	    Mobile:  "15xxxxxxxxx",
	    City:    "北京市,北京市,xx区",
	    Address: "详细地址",
    },
   }
   
   resp, err := client.SubmitOrderCode(ztoContent)
   // TODO ......
   ```
   
   ##### 电子面单 -> 获取单号（无密钥）
   ```go
   ztoContent := &common.ZTOContent{
   // TODO ...
   }
   
   resp, err := client.PartnerInsertSubmitagent(ztoContent)
   // TODO ......
   ```
   
   ##### 电子面单 -> 云打印-打印接口
   ```go
   request := &common.ZTOPrintRequest{
   // TODO ...
   }
   resp, err := client.DoPrint(request)
   // TODO ......
   ```
   
   ##### 订单服务 -> 预约寄件-订单取消
   ```go
   resp, err := client.CancelOrder("订单号", "取消原因")
   // TODO ......
   ```
   
## License
  feeeei/ztosdk is released under the [MIT License](https://opensource.org/licenses/MIT).