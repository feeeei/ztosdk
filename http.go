package ztosdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/feeeei/ztosdk/common"
)

var httpClient *http.Client

func init() {
	httpClient = &http.Client{}
}

func (client *ZTOClient) postOrderRequest(path, sign string, r common.ZTORequest) (*common.ZTOOrderResponse, error) {
	resp, err := client.postRequest(path, sign, r)
	if err != nil {
		return nil, err
	}
	var response common.ZTOOrderResponse
	if err := json.Unmarshal(*(*resp)["data"], &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (client *ZTOClient) updateOrderRequest(path, sign string, r common.ZTORequest) (*[]common.ZTOUpdateResponse, error) {
	resp, err := client.postRequest(path, sign, r)
	if err != nil {
		return nil, err
	}
	var responses []common.ZTOUpdateResponse
	if err := json.Unmarshal(*(*resp)["data"], &responses); err != nil {
		return nil, err
	}

	return &responses, nil
}

func (client *ZTOClient) postPrintRequest(path, sign string, r common.ZTORequest) (*common.ZTOPrintResponse, error) {
	resp, err := client.postRequest(path, sign, r)
	if err != nil {
		return nil, err
	}
	var response common.ZTOPrintResponse
	if err := json.Unmarshal(*(*resp)["result"], &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (client *ZTOClient) postRequest(path, sign string, r common.ZTORequest) (*map[string]*json.RawMessage, error) {
	url := fmt.Sprintf("%s%s", client.Host, path)
	requestBody, err := r.EncodeBody()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, strings.NewReader(requestBody))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	req.Header.Add("X-Companyid", client.CompanyID)
	req.Header.Add("X-Datadigest", sign)

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var respbody map[string]*json.RawMessage
	if err := json.Unmarshal(body, &respbody); err != nil {
		return nil, err
	}
	if respbody["status"] != nil && string(*respbody["status"]) == "false" {
		message := strings.Replace(string(*respbody["message"]), "\"", "", -1)
		statusCode := strings.Replace(string(*respbody["statusCode"]), "\"", "", -1)
		return nil, fmt.Errorf("%s %s", statusCode, message)
	}
	return &respbody, nil
}
