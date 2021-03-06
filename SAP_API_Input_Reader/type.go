package sap_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema      string `json:"api_schema"`
	MaterialCode   string `json:"material_code"`
	Plant_Supplier string `json:"plant/supplier"`
	Stock          string `json:"stock"`
	DocumentType   string `json:"document_type"`
	DocumentNo     string `json:"document_no"`
	PlannedDate    string `json:"planned_date"`
	ValidatedDate  string `json:"validated_date"`
	Deleted        bool   `json:"deleted"`
}

type SDC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Defect        struct {
		DefectInternalID            string      `json:"DefectInternalID"`
		Defect                      string      `json:"Defect"`
		DefectCategory              string      `json:"DefectCategory"`
		CreationDate                string      `json:"CreationDate"`
		LastChangeDate              string      `json:"LastChangeDate"`
		DefectText                  string      `json:"DefectText"`
		DefectCodeCatalog           string      `json:"DefectCodeCatalog"`
		DefectCodeGroup             string      `json:"DefectCodeGroup"`
		DefectCode                  string      `json:"DefectCode"`
		DefectCodeVersion           string      `json:"DefectCodeVersion"`
		DefectObjectCodeCatalog     string      `json:"DefectObjectCodeCatalog"`
		DefectObjectCodeGroup       string      `json:"DefectObjectCodeGroup"`
		DefectObjectCode            string      `json:"DefectObjectCode"`
		DefectiveQuantity           string      `json:"DefectiveQuantity"`
		DefectiveQuantityUnit       string      `json:"DefectiveQuantityUnit"`
		ManufacturingOrder          string      `json:"ManufacturingOrder"`
		OrderInternalID             string      `json:"OrderInternalID"`
		ManufacturingOrderOperation string      `json:"ManufacturingOrderOperation"`
		ManufacturingOrderSequence  string      `json:"ManufacturingOrderSequence"`
		CreationTime                string      `json:"CreationTime"`
		LastChangeTime              string      `json:"LastChangeTime"`
		DefectClass                 string      `json:"DefectClass"`
		NumberOfDefects             int         `json:"NumberOfDefects"`
		InspPlanOperationInternalID string      `json:"InspPlanOperationInternalID"`
		InspectionCharacteristic    string      `json:"InspectionCharacteristic"`
		InspectionSubsetInternalID  string      `json:"InspectionSubsetInternalID"`
		MaterialSample              string      `json:"MaterialSample"`
		WorkCenterTypeCode          string      `json:"WorkCenterTypeCode"`
		MainWorkCenterInternalID    string      `json:"MainWorkCenterInternalID"`
		MainWorkCenterPlant         string      `json:"MainWorkCenterPlant"`
		MainWorkCenter              string      `json:"MainWorkCenter"`
		Equipment                   string      `json:"Equipment"`
		FunctionalLocation          string      `json:"FunctionalLocation"`
		IsDeleted                   string      `json:"IsDeleted"`
		DefectOrigin                string      `json:"DefectOrigin"`
		Material                    string      `json:"Material"`
		Plant                       string      `json:"Plant"`
		InspectionLot               string      `json:"InspectionLot"`
		CatalogProfile              string      `json:"CatalogProfile"`
		ChangedDateTime             string      `json:"ChangedDateTime"`
	} `json:"Defect"`
	APISchema  string   `json:"api_schema"`
	Accepter   []string `json:"accepter"`
	DefectCode string   `json:"defect_code"`
	Deleted    bool     `json:"deleted"`
}
