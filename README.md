# sap-api-integrations-bill-of-material-reads  
sap-api-integrations-bill-of-material-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で部品表（BOM）データを取得するマイクロサービスです。  
sap-api-integrations-bill-of-material-reads には、サンプルのAPI Json フォーマットが含まれています。  
sap-api-integrations-bill-of-material-reads は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。  
https://api.sap.com/api/OP_API_BILL_OF_MATERIAL_SRV_0002/overview  

## 動作環境
sap-api-integrations-bill-of-material-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。   
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。   
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須） 

## クラウド環境での利用  
sap-api-integrations-bill-of-material-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-bill-of-material-reads が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/OP_API_BILL_OF_MATERIAL_SRV_0002/overview  
* APIサービス名(=baseURL): API_BILL_OF_MATERIAL_SRV;v=0002

## 本レポジトリ に 含まれる API名
sap-api-integrations-bill-of-material-reads には、次の API をコールするためのリソースが含まれています。  

* MaterialBOMItem（部品表 - 明細）

## API への 値入力条件 の 初期値
sap-api-integrations-bill-of-material-reads において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.BillOfMaterial.Material（品目）
* inoutSDC.BillOfMaterial.Plant（プラント）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"Product", "Plant", "Accounting" が指定されています。    
  
```
"api_schema": "sap.s4.beh.billofmaterial.v1.BillOfMaterial.Created.v1",
"accepter": ["Item"],
"material_code": "",
"plant": "1010",
"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
"api_schema": "sap.s4.beh.billofmaterial.v1.BillOfMaterial.Created.v1",
"accepter": ["All"],
"material_code": "",
"plant": "1010",
"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncGetBillOfMaterial(material, plant string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
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
```

## Output  
本マイクロサービスでは、[golang-logging-library](https://github.com/latonaio/golang-logging-library) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP 部品表  の 明細 が取得された結果の JSON の例です。  
以下の項目のうち、"BaseUnit" ～ "WeightUnit" は、/SAP_API_Output_Formatter/type.go 内 の Type Item {} による出力結果です。"cursor" ～ "time"は、golang-logging-library による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-bill-of-material-reads/SAP_API_Caller/caller.go#L46",
	"function": "sap-api-integrations-bill-of-material-reads/SAP_API_Caller.(*SAPAPICaller).Item",
	"level": "INFO",
	"message": "&{BillOfMaterial:00000001 BillOfMaterialVariant:1 BillOfMaterialCategory:M BillOfMaterialVersion: BillOfMaterialItemNodeNumber:1 HeaderChangeDocument: Material:SG23 Plant:1010 ValidityStartDate:/Date(1136073600000)/ ValidityEndDate:/Date(253402214400000)/ BillOfMaterialComponent:RM13 ComponentDescription:RAW13,PD,Subcontracting BillOfMaterialItemQuantity:100 ComponentScrapInPercent:0.00 IsDeleted:false}",
	"time": "2021-12-08T21:13:42.001292+09:00"
}
```