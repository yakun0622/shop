package models

type SunGroupInformation struct {
	InformationGroupId           uint   `orm:"column(information_group_id)"`
	InformationLicenseNumber     string `orm:"column(information_licenseNumber);size(45);null"`
	InformationLicenseImage      string `orm:"column(information_licenseImage);size(255);null"`
	InformationOrganizationCode  string `orm:"column(information_organizationCode);size(45);null"`
	InformationOrganizationImage string `orm:"column(information_organizationImage);size(255);null"`
	InformationTaxId             string `orm:"column(information_taxId);size(45);null"`
	InformationTaxImage          string `orm:"column(information_taxImage);size(45);null"`
	InformationCreateDate        int    `orm:"column(information_createDate);null"`
	InformationLicenseStart      int    `orm:"column(information_licenseStart);null"`
	InformationLicenseEnd        int    `orm:"column(information_licenseEnd);null"`
	InformationManageScope       string `orm:"column(information_manageScope);size(1000);null"`
	InformationRegisterCapital   int    `orm:"column(information_registerCapital);null"`
	InformationBrandAgent        string `orm:"column(information_brandAgent);size(128);null"`
	InformationBrandAgentNode    string `orm:"column(information_brandAgentNode);size(255);null"`
	InformationAfterSaleAdress   string `orm:"column(information_afterSaleAdress);size(45);null"`
	InformationFax               string `orm:"column(information_fax);size(45);null"`
}
