package sap_api_output_formatter

type BillOfMaterialReads struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	APISchema     string `json:"api_schema"`
	MaterialCode  string `json:"material_code"`
	Plant         string `json:"plant"`
	Deleted       bool   `json:"deleted"`
}

type BillOfMaterial struct {
	Material              string `json:"Material"`
	Plant                 string `json:"Plant"`
	BillOfMaterial        string `json:"BillOfMaterial"`
	BillOfMaterialVariant string `json:"BillOfMaterialVariant"`
	ValidityStartDate     string `json:"ValidityStartDate"`
	ValidityEndDate       string `json:"ValidityEndDate"`
	HeaderIsDeleted       bool   `json:"HeaderIsDeleted"`
}

type BillOfMaterialItemNodeNumber struct {
	Material                     string `json:"Material"`
	Plant                        string `json:"Plant"`
	BillOfMaterial               string `json:"BillOfMaterial"`
	BillOfMaterialVariant        string `json:"BillOfMaterialVariant"`
	BillOfMaterialItemNodeNumber string `json:"BillOfMaterialItemNodeNumber"`
	BillOfMaterialComponent      string `json:"BillOfMaterialComponent"`
	BillOfMaterialItemQuantity   string `json:"BillOfMaterialItemQuantity"`
	ComponentScrapInPercent      string `json:"ComponentScrapInPercent"`
	ItemIsDeleted                bool   `json:"ItemIsDeleted"`
}
