# sap-api-integrations-defect-reads  
sap-api-integrations-defect-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で欠陥データを取得するマイクロサービスです。  
sap-api-integrations-defect-reads には、サンプルのAPI Json フォーマットが含まれています。  
sap-api-integrations-defect-reads は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。  
https://api.sap.com/api/OP_API_DEFECT_SRV_0001/overview   

## 動作環境
sap-api-integrations-defect-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。   
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。   
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須） 

## クラウド環境での利用  
sap-api-integrations-defect-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-defect-reads が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/OP_API_DEFECT_SRV_0001/overview    
* APIサービス名(=baseURL): API_DEFECT_SRV  

## 本レポジトリ に 含まれる API名
sap-api-integrations-defect-reads には、次の API をコールするためのリソースが含まれています。  

* A_Defect（欠陥 - ヘッダ）

## API への 値入力条件 の 初期値
sap-api-integrations-defect-reads において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.Defect.Defect（欠陥）
* inoutSDC.Defect.DefectText（欠陥テキスト）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"Header" が指定されています。    
  
```
	"api_schema": "sap.s4.beh.defect.v1.Defect.Created.v1",
	"accepter": ["Header"],
	"defect_code": "4",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "sap.s4.beh.defect.v1.Defect.Created.v1",
	"accepter": ["All"],
	"defect_code": "4",
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncGetDefect(defect, defectText string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(defect)
				wg.Done()
			}()
		case "DefectText":
			func() {
				c.DefectText(defectText)
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
本マイクロサービスでは、[golang-logging-library-for-sap](https://github.com/latonaio/golang-logging-library-for-sap) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP 欠陥 の ヘッダ が取得された結果の JSON の例です。  
以下の項目のうち、"DefectInternalID" ～ "ChangedDateTime" は、/SAP_API_Output_Formatter/type.go 内 の Type Header {} による出力結果です。"cursor" ～ "time"は、golang-logging-library-for-sap による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-defect-reads/SAP_API_Caller/caller.go#L58",
	"function": "sap-api-integrations-defect-reads/SAP_API_Caller.(*SAPAPICaller).Header",
	"level": "INFO",
	"message": [
		{
			"DefectInternalID": "$00000000004",
			"Defect": "4",
			"DefectCategory": "06",
			"CreationDate": "2019-10-28T09:00:00+09:00",
			"LastChangeDate": "2020-04-07T09:00:00+09:00",
			"DefectText": "Incomplete Packing",
			"DefectCodeCatalog": "9",
			"DefectCodeGroup": "DEF-PACK",
			"DefectCode": "0001",
			"DefectCodeVersion": "1",
			"DefectObjectCodeCatalog": "E",
			"DefectObjectCodeGroup": "",
			"DefectObjectCode": "",
			"DefectiveQuantity": "0.000",
			"DefectiveQuantityUnit": "",
			"ManufacturingOrder": "",
			"OrderInternalID": "0",
			"ManufacturingOrderOperation": "",
			"ManufacturingOrderSequence": "",
			"CreationTime": "PT12H32M29S",
			"LastChangeTime": "PT05H43M09S",
			"DefectClass": "92",
			"NumberOfDefects": 1,
			"InspPlanOperationInternalID": "0",
			"InspectionCharacteristic": "0",
			"InspectionSubsetInternalID": "0",
			"MaterialSample": "",
			"WorkCenterTypeCode": "",
			"MainWorkCenterInternalID": "0",
			"MainWorkCenterPlant": "",
			"MainWorkCenter": "",
			"Equipment": "",
			"FunctionalLocation": "",
			"IsDeleted": "",
			"DefectOrigin": "",
			"Material": "",
			"Plant": "",
			"InspectionLot": "0",
			"CatalogProfile": "",
			"ChangedDateTime": ""
		}
	],
	"time": "2022-01-28T17:49:49+09:00"
}
```
