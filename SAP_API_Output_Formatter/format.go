package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-bill-of-material-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
)

func ConvertToItem(raw []byte, l *logger.Logger) (*Item, error) {
	pm := &responses.Item{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Item. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 1 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &Item{
		BillOfMaterial:               data.BillOfMaterial,
		BillOfMaterialVariant:        data.BillOfMaterialVariant,
		BillOfMaterialCategory:       data.BillOfMaterialCategory,
		BillOfMaterialVersion:        data.BillOfMaterialVersion,
		BillOfMaterialItemNodeNumber: data.BillOfMaterialItemNodeNumber,
		HeaderChangeDocument:         data.HeaderChangeDocument,
		Material:                     data.Material,
		Plant:                        data.Plant,
		ValidityStartDate:            data.ValidityStartDate,
		ValidityEndDate:              data.ValidityEndDate,
		BillOfMaterialComponent:      data.BillOfMaterialComponent,
		ComponentDescription:         data.ComponentDescription,
		BillOfMaterialItemQuantity:   data.BillOfMaterialItemQuantity,
		ComponentScrapInPercent:      data.ComponentScrapInPercent,
		IsDeleted:                    data.IsDeleted,
	}, nil
}
