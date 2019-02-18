package common

import (
	"fmt"

	"github.com/feeeei/ztosdk/common/base"
)

type ZTOContent struct {
	Partner          string    `json:"partner,omitempty"`
	ID               string    `json:"id"`
	TypeID           string    `json:"typeid"`
	TradeID          string    `json:"tradeid,omitempty"`
	BranchID         string    `json:"branch_id,omitempty"`
	Seller           string    `json:"seller,omitempty"`
	Buyer            string    `json:"buyer,omitempty"`
	Sender           *Sender   `json:"sender"`
	Receiver         *Receiver `json:"receiver"`
	Weight           float32   `json:"weight,omitempty"`
	Size             string    `json:"size,omitempty"`
	Quantity         string    `json:"quantity,omitempty"`
	Price            float32   `json:"price,omitempty"`
	Freight          float32   `json:"freight,omitempty"`
	Premium          float32   `json:"premium,omitempty"`
	PackCharges      float32   `json:"packCharges,omitempty"`
	OtherCharges     float32   `json:"otherCharges,omitempty"`
	OrderSum         float32   `json:"orderSum,omitempty"`
	CollectMoneyType string    `json:"collectMoneytype,omitempty"`
	CollectSum       float32   `json:"collectSum,omitempty"`
	Remark           string    `json:"remark,omitempty"`
	OrderType        string    `json:"order_type,omitempty"`
}

type Sender = worker
type Receiver = worker

// Prov City County 省市区三级关系，会根据不同request，构造为对应request要求格式
type worker struct {
	ID        string     `json:"id,omitempty"`
	Name      string     `json:"name"`
	Company   string     `json:"company,omitempty"`
	Mobile    string     `json:"mobile,omitempty"`
	Phone     string     `json:"Phone,omitempty"`
	Area      int        `json:"area,omitempty"`
	Prov      string     `json:"prov,omitempty"`
	City      string     `json:"city"`
	County    string     `json:"County,omitempty"`
	Address   string     `json:"address,omitempty"`
	ZIPCode   string     `json:"zipcode,omitempty"`
	Email     string     `json:"email,omitempty"`
	Im        string     `json:"im,omitempty"`
	StartTime *base.Time `json:"starttime,omitempty"`
	EndTime   *base.Time `json:"endtime,omitempty"`
}

func transportToSimpleAddress(worker *worker) {
	worker.City = fmt.Sprintf("%s,%s,%s", worker.Prov, worker.City, worker.County)
	worker.Prov = ""
	worker.County = ""
}
