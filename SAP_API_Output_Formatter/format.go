package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-bill-of-material-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library/logger"
)

func ConvertToBillOfMaterial(raw []byte, l *logger.Logger) *BillOfMaterial {
	pm := &responses.BillOfMaterial{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		l.Error(err)
		return nil
	}
	if len(pm.D.Results) == 0 {
		l.Error("Result data is not exist.")
		return nil
	}
	if len(pm.D.Results) > 1 {
		l.Error("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &BillOfMaterial{
		Material                     data.Material,
		Plant                        data.Plant,
		BillOfMaterial               data.BillOfMaterial,
		BillOfMaterialVariant        data.BillOfMaterialVariant,
		ValidityStartDate            data.ValidityStartDate,
		ValidityEndDate              data.ValidityEndDate,
		HeaderIsDeleted              data.HeaderIsDeleted,
		BillOfMaterialItemNodeNumber data.BillOfMaterialItemNodeNumber,
		BillOfMaterialComponent      data.BillOfMaterialComponent,
		BillOfMaterialItemQuantity   data.BillOfMaterialItemQuantity,
		ComponentScrapInPercent      data.ComponentScrapInPercent,
		ItemIsDeleted                data.ItemIsDeleted,
	}
}
