package sap_api_caller

type BillOfMaterialReads struct {
	 ConnectionKey    string `json:"connection_key"`
	 Result           bool   `json:"result"`
	 RedisKey         string `json:"redis_key"`
	 Filepath         string `json:"filepath"`
	 APISchema        string `json:"api_schema"`
	 MaterialCode     string `json:"material_code"`
	 Plant            string `json:"plant"`
	 Deleted          string `json:"deleted"`
}

type BillOfMaterial struct {
     Material                     string `json:"Material"`
     Plant                        string `json:"Plant"`
     BillOfMaterial               string `json:"BillOfMaterial"`
     BillOfMaterialVariant        string `json:"BillOfMaterialVariant"`
     ValidityStartDate            string `json:"ValidityStartDate"`
     ValidityEndDate              string `json:"ValidityEndDate"`
     HeaderIsDeleted              string `json:"HeaderIsDeleted"`
}

type BillOfMaterialItemNodeNumber struct {
     Material                     string `json:"Material"`
     Plant                        string `json:"Plant"`
     BillOfMaterial               string `json:"BillOfMaterial"`
     BillOfMaterialVariant        string `json:"BillOfMaterialVariant"`
     BillOfMaterialItemNodeNumber string `json:"BillOfMaterialItemNodeNumber"`
     BillOfMaterialComponent      string `json:"BillOfMaterialComponent"`
     BillOfMaterialItemQuantity   float64 `json:"BillOfMaterialItemQuantity"`
     ComponentScrapInPercent      float64 `json:"ComponentScrapInPercent"`
     ItemIsDeleted                string `json:"ItemIsDeleted"`
}
