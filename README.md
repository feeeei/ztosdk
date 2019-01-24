## 中通开放平台SDK for Go

官方 API 文档地址：https://zop.zto.com/apiDoc/

## 实现接口列表
 - 获取运单号（有秘钥） submitOrderCode

## 使用示例
1. 引入

   ```go
   go get github.com/feeeei/ztosdk
   ```
2. New Client
   ```go
   // 测试环境参数
   host      := "http://58.40.16.120:9001/"
   companyID := "kfpttestCode"
   key       := "kfpttestkey=="
   partner   := "test"
   verify    := "ZTO123"
 
   client := ztosdk.NewZTOClient(host, companyID, key, partner, verify)
   ```
3. 实现业务
 
   ```go
 
   // 获取运单号（有秘钥） submitOrderCode
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
	})
   resp, err := client.SubmitOrderCode(ztoContent)
   //TODO ......
   ```
## License
  feeeei/ztosdk is released under the [MIT License](https://opensource.org/licenses/MIT).