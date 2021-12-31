package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-bill-of-material-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
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

func (c *SAPAPICaller) AsyncGetBillOfMaterial(material, plant string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(material, plant)
				wg.Done()
			}()
		case "Item":
			func() {
				c.Item(material, plant)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) Header(material, plant string) {
	headerData, err := c.callBillOfMaterialSrvAPIRequirementHeader("MaterialBOM", material, plant)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(headerData)

	itemData, err := c.callToItem(headerData[0].ToItem)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(itemData)

}

func (c *SAPAPICaller) callBillOfMaterialSrvAPIRequirementHeader(api, material, plant string) ([]sap_api_output_formatter.Header, error) {
	url := strings.Join([]string{c.baseURL, "API_BILL_OF_MATERIAL_SRV;v=0002", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithHeader(req, material, plant)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToItem(url string) ([]sap_api_output_formatter.ToItem, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToItem(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Item(material, plant string) {
	data, err := c.callBillOfMaterialSrvAPIRequirementItem("MaterialBOMItem", material, plant)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callBillOfMaterialSrvAPIRequirementItem(api, material, plant string) ([]sap_api_output_formatter.Item, error) {
	url := strings.Join([]string{c.baseURL, "API_BILL_OF_MATERIAL_SRV;v=0002", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithItem(req, material, plant)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToItem(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithHeader(req *http.Request, material, plant string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("Material eq '%s' and Plant eq '%s'", material, plant))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithItem(req *http.Request, material, plant string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("Material eq '%s' and Plant eq '%s'", material, plant))
	req.URL.RawQuery = params.Encode()
}
