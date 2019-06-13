package fostaskmapper

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"encoding/xml"
  "encoding/json"
)

const (
	ivrequestType   = "requestType"
	ivrequestXML      = "requestXML"
	ovresultJSON    = "responseJSON"
)

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}


type SlaType struct {
	TypicalDuration int `xml:"typicalDruation"`
	MaximumDuration int `xml:"maximumDuration"`
}
type OrderLineType struct {
	OrderLineNumber string `xml:"orderLineNumber"`
	ProductID       string `xml:"productID"`
	Action          string `xml:"Action"`
	Quantity        string `xml:"quantity"`
	UoM             string `xml:"uom"`
	EoL             string `xml:"eol"`
}

type PlanItemType struct {
	PlanItemID                 string        `xml:"planItemID"`
	Description                string        `xml:"description"`
	ProcessComponentID         string        `xml:"processComponentID"`
	ProcessComponentName       string        `xml:"processComponentName"`
	ProcessComponentVersion    string        `xml:"processComponentVersion"`
	ProcessComponentRecordType string        `xml:"processComponentRecordType"`
	OrderLine                  OrderLineType `xml:"orderLine"`
	Action                     string        `xml:"action"`
}

type PlanItemExecuteRequest struct {
	OrderID  string       `xml:"orderID"`
	OrderRef string       `xml:"orderRef"`
	PlanID   string       `xml:"planID"`
	PlanItem PlanItemType `xml:"planItem"`
	Sla      SlaType      `xml:"sla"`
}


// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	requestType := context.GetInput(ivrequestType).(string)
  requestXML := context.GetInput(ivrequestXML).(string)
	switch requestType {

	case "PlanItemExecuteRequest":

    req := PlanItemExecuteRequest{}

    err := xml.Unmarshal([]byte(requestXML), &req)

		if err != nil {
			return false, err
		}else {
      b, err := json.Marshal(req)
		  if err != nil {
			  return false, err
		  }
      //output := string(b)
		 //context.SetOutput(ovresultJSON, output)
		 context.SetOutput(ovresultJSON, req)
     if string(b) != "" {
     }

		}

	}

	return true, nil
}
