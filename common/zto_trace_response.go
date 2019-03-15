package common

import (
	"github.com/feeeei/ztosdk/common/base"
)

type ZTOTraceResponse struct {
	BillCode string          `json:"billCode"`
	Traces   *[]ZTOTraceInfo `json:"traces"`
}

type ZTOLastTraceResponse struct {
	BillCode string       `json:"billCode"`
	Traces   ZTOTraceInfo `json:"traces"`
}

type ZTOTraceInfo struct {
	Desc              string     `json:"desc"`
	DispOrRecMan      string     `json:"dispOrRecMan,omitempty"`
	DispOrRecManCode  string     `json:"dispOrRecManCode,omitempty"`
	DispOrRecManPhone string     `json:"dispOrRecManPhone,omitempty"`
	IsCenter          string     `json:"isCenter"`
	PreOrNextCity     string     `json:"preOrNextCity,omitempty"`
	PreOrNextProv     string     `json:"preOrNextProv,omitempty"`
	PreOrNextSite     string     `json:"preOrNextSite,omitempty"`
	PreOrNextSiteCode string     `json:"preOrNextSiteCode,omitempty"`
	Remark            string     `json:"remark,omitempty"`
	ScanCity          string     `json:"scanCity"`
	ScanDate          *base.Time `json:"scanDate"`
	ScanProv          string     `json:"scanProv"`
	ScanSite          string     `json:"scanSite"`
	ScanSiteCode      string     `json:"scanSiteCode"`
	ScanSitePhone     string     `json:"scanSitePhone"`
	ScanType          string     `json:"scanType"`
	SignMan           string     `json:"signMan"`
}
