package responses

type BillOfMaterial struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			Material              string `json:"Material"`
			Plant                 string `json:"Plant"`
			BillOfMaterial        string `json:"BillOfMaterial"`
			BillOfMaterialVariant string `json:"BillOfMaterialVariant"`
			ValidityStartDate     string `json:"ValidityStartDate"`
			ValidityEndDate       string `json:"ValidityEndDate"`
			HeaderIsDeleted       bool   `json:"HeaderIsDeleted"`
		} `json:"results"`
	} `json:"d"`
}
