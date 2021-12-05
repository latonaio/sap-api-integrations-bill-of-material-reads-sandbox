package main

import (
	sap_api_caller "sap-api-integrations-bill-of-material-reads/SAP_API_Caller"
	"sap-api-integrations-bill-of-material-reads/sap_api_input_reader"

	"github.com/latonaio/golang-logging-library/logger"
)

func main() {
	l := logger.NewLogger()
	fr := sap_api_input_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs//SDC_Bill_Of_Material_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/", l,
	)

    caller.AsyncGetBillOfMaterial(
        inoutSDC.BillOfMaterial.Material,
        inoutSDC.BillOfMaterial.Plant,
        inoutSDC.BillOfMaterial.ValidityEndDate,
    )
}
