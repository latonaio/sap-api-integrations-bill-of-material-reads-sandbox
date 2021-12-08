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

type Item struct {
		BillOfMaterial               string `json:"BillOfMaterial"`
		BillOfMaterialVariant        string `json:"BillOfMaterialVariant"`
		BillOfMaterialCategory       string `json:"BillOfMaterialCategory"`
		BillOfMaterialVersion        string `json:"BillOfMaterialVersion"`
		BillOfMaterialItemNodeNumber string `json:"BillOfMaterialItemNodeNumber"`
		HeaderChangeDocument         string `json:"HeaderChangeDocument"`
		Material                     string `json:"Material"`
		Plant                        string `json:"Plant"`
		ValidityStartDate            string `json:"ValidityStartDate"`
		ValidityEndDate              string `json:"ValidityEndDate"`
		BillOfMaterialComponent      string `json:"BillOfMaterialComponent"`
		ComponentDescription         string `json:"ComponentDescription"`
		BillOfMaterialItemQuantity   string `json:"BillOfMaterialItemQuantity"`
		ComponentScrapInPercent      string `json:"ComponentScrapInPercent"`
		IsDeleted                    bool   `json:"IsDeleted"`
}
