package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library/logger"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
	}
}

func (c *SAPAPICaller) AsyncGetBillOfMaterial(Material, Plant, ValidityEndDate string) {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		c.BillOfMaterial(Material, Plant, ValidityEndDate)
		wg.Done()
	}()
	wg.Wait()
}

func (c *SAPAPICaller) BillOfMaterial(Material, Plant, ValidityEndDate string) {
	res, err := c.callBillOfMaterialSrvAPIRequirement("MaterialBOMItem", Material, Plant, ValidityEndDate)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)

}

func (c *SAPAPICaller) callBillOfMaterialSrvAPIRequirement(api, Material, Plant, ValidityEndDate string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "API_BILL_OF_MATERIAL_SRV;v=0002", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "Material, Plant, ValidityEndDate")
	params.Add("$filter", fmt.Sprintf("Material eq '%s' and Plant eq '%s' and ValidityEndDate eq '%s'", Material, Plant, ValidityEndDate))
	req.URL.RawQuery = params.Encode()

	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, nil
}