package common

import (
	"github.com/feeeei/ztosdk/common/base"
)

type ZTOOrderSearchRequest struct {
	SendID    string     `json:"sendId"`
	StartTime *base.Time `json:"starttime,omitempty"`
	EndTime   *base.Time `json:"endtime,omitempty"`
	Status    int        `json:"status,omitempty"`
	SendName  string     `json:"sendName,omitempty"`
	SendPhone string     `json:"sendPhone,omitempty"`
	RecName   string     `json:"recName,omitempty"`
	RecPhone  string     `json:"recPhone,omitempty"`
	PageSize  int        `json:"pageSize,omitempty"`
	PageIndex int        `json:"pageIndex,omitempty"`
}
