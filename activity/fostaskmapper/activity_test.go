package fostaskmapper

import (
	"io/ioutil"
	"testing"
  "encoding/json"
	"fmt"
	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func getRequestXML() string {

  var requestwithns = `
<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<ns21:PlanItemExecuteRequest xmlns:ns21="http://www.tibco.com/AFF/V4.0.0/services/orchestrator/execution" xmlns="http://www.tibco.com/AFF/classes/resultstatus" xmlns:ns10="http://www.tibco.com/AFF/V4.0.0/applications/orchestrator/message" xmlns:ns11="http://www.tibco.com/AFF/V4.0.0/applications/orchestrator/order" xmlns:ns12="http://www.tibco.com/AFF/V4.0.0/applications/orchestrator/plan" xmlns:ns13="http://www.tibco.com/AFF/V4.0.0/classes/orchestrator/processComponentModel" xmlns:ns14="http://www.tibco.com/AFF/V4.0.0/classes/order" xmlns:ns15="http://www.tibco.com/AFF/V4.0.0/types/commontypes" xmlns:ns16="http://www.tibco.com/AFF/V4.0.0/services/orchestrator/provider" xmlns:ns17="http://www.tibco.com/AFF/V4.0.0/classes/resultstatus" xmlns:ns18="http://www.tibco.com/AFF/V4.0.0/classes/planFragment" xmlns:ns19="http://www.tibco.com/AFF/V4.0.0/services/orchestrator/data" xmlns:ns2="http://www.staffware.com/frameworks/gen/valueobjects" xmlns:ns20="http://www.tibco.com/AFF/V4.0.0/services/transientDataStore" xmlns:ns22="http://www.tibco.com/AFF/V4.0.0/services/orchestrator/model" xmlns:ns23="http://www.tibco.com/AFF/V4.0.0/services/orchestrator/order" xmlns:ns24="http://www.tibco.com/AFF/V4.0.0/services/orchestrator/plan" xmlns:ns25="http://www.tibco.com/aff/commontypes" xmlns:ns26="http://www.tibco.com/aff/plan" xmlns:ns27="http://www.tibco.com/aff/planfragments" xmlns:ns28="http://www.tibco.com/AFF/classes/segment" xmlns:ns29="http://www.tibco.com/AFF/classes/productmodel" xmlns:ns3="http://www.tibco.com/AFF/classes/inventory" xmlns:ns30="http://www.tibco.com/AFF/classes/role" xmlns:ns31="http://www.tibco.com/AFF/classes/action" xmlns:ns32="http://www.tibco.com/AFF/classes/customermodel" xmlns:ns33="http://www.tibco.com/AFF/classes/customer" xmlns:ns34="http://www.tibco.com/AFF/classes/discountmodel" xmlns:ns35="http://www.tibco.com/AFF/classes/pricemodel" xmlns:ns36="http://www.tibco.com/aff/eca/model/service" xmlns:ns37="http://www.tibco.com/aff/eca/model" xmlns:ns38="http://www.tibco.com/fom/customAuditTrail" xmlns:ns39="http://www.tibco.com/AFF/classes/rulemodel" xmlns:ns4="http://www.tibco.com/AFF/OCV/services" xmlns:ns40="http://www.tibco.com/aff/internalerrorhandler" xmlns:ns41="http://www.tibco.com/aff/orderservice" xmlns:ns5="www.tibco.com/be/ontology/OfferConfigurationAndValidation/Events/Amendments/Events/ExecutionPlanAmendRequest" xmlns:ns6="http://www.tibco.com/AFF/executionPlan/ExecutionPlanAmendResponse" xmlns:ns7="www.tibco.com/be/ontology/OfferConfigurationAndValidation/Events/AOPD/Events/PlanDevelopment/ExecutionPlanNewRequest" xmlns:ns8="http://www.tibco.com/AFF/V4.0.0/applications/orchestrator/aopd" xmlns:ns9="http://www.tibco.com/AFF/V4.0.0/applications/orchestrator/notification" businessTransactionID="97a37a90fdcb4680b32fc56fa46c0632" correlationID="2ba3843d-6743-4c94-bf34-caa3f485850d">
  <ns21:orderID>128</ns21:orderID>
  <ns21:orderRef>OCS-fa2bf996-ba67-4170-86d8-b9f33a22d9c8</ns21:orderRef>
  <ns21:planID>126</ns21:planID>
  <ns21:planItem>
    <ns21:planItemID>15</ns21:planItemID>
    <ns21:description>Mailserver Activate</ns21:description>
    <ns21:processComponentID>MAILSERVER_ACTIVATE</ns21:processComponentID>
    <ns21:processComponentName>MAILSERVER_ACTIVATE</ns21:processComponentName>
    <ns21:processComponentVersion>1</ns21:processComponentVersion>
    <ns21:processComponentRecordType>Process</ns21:processComponentRecordType>
    <ns21:orderLine>
      <ns21:orderLineNumber>1</ns21:orderLineNumber>
      <ns21:productID>PO_EMAIL_BOX</ns21:productID>
      <ns21:action>PROVIDE</ns21:action>
      <ns21:quantity>2</ns21:quantity>
      <ns21:uom>NA</ns21:uom>
      <ns21:eol>false</ns21:eol>
    </ns21:orderLine>
    <ns21:action>PROVIDE</ns21:action>
  </ns21:planItem>
  <ns21:sla>
    <ns21:typicalDuration>2000</ns21:typicalDuration>
    <ns21:maximumDuration>5000</ns21:maximumDuration>
  </ns21:sla>
</ns21:PlanItemExecuteRequest>
`
  return requestwithns;
}

func TestPlanItemExecuteRequest(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())
  //fmt.Println("RequestXML#########")
  //fmt.Println(getRequestXML())
	//setup attrs
	tc.SetInput(ivrequestType, "PlanItemExecuteRequest")
  tc.SetInput(ivrequestXML, getRequestXML())
	_, err := act.Eval(tc)
	if err != nil {
		t.Error("unable to set env value", err)
		t.Fail()
	}
  output := tc.GetOutput(ovresultJSON)
  result, err := json.Marshal(output)
  if err != nil {
    fmt.Println(" Error = ",err)
    return
  }
  fmt.Println("Output JSON=" , string(result))
  fmt.Println("Output =", output)
	//fmt.Println("Output =",tc.GetOutput(ovresultJSON).(string))
  

}

