package main

//NetworkInfoResponse for Network Info
type NetworkInfoResponse struct {
	Status string `json:"Status"`
	Data NetworkInfoData `json:"data"`
}

//NetworkInfoData Information From Request
type NetworkInfoData struct {
	DHCP                     string `json:"DHCP"`
	DHCPServer               string	`json:"DHCPServer"`
	DHCPBootServerUserOption string `json:"DHCPBootServerUserOption"`
	DHCPBootServerOption     string `json:"DHCPBootServerOption"`
	VLANID string `json:"VLANID"`
	DefaultGateway string `json:"DefaultGateway"`
	IPV4Address string `json:"IPV4Address"`
	LANSpeed string `json:"LANSpeed"`
	DHCPBootServerOptionType string `json:"DHCPBootServerOptionType"`
	DNSDomain string `json:"DNSDomain"`
	CDPCompability string `json:"CDPCompability"`
	DHCPOption60Format string `json:"DHCPOption60Format"`
	IPV6Address string `json:"IPV6Address"`
	SNTPAddress string `json:"SNTPAddress"`
	VLANDiscoveryMode string `json:"VLANDiscoveryMode"`
	SubnetMask string `json:"SubnetMask"`
	LANPortStatus string `json:"LANPortStatus"`
	ProvServerType string `json:"ProvServerType"`
	DNSServer string `json:"DNSServer"`
	UpgradeServer string `json:"UpgradeServer"`
	AlternateDNSServer string `json:"AlternateDNSServer"`
	VLANIDOption string `json:"VLANIDOption"`
	LLDP string `json:"LLDP"`
	ProvServerAddress string `json:"ProvServerAddress"`
	ProvServerUser string `json:"ProvServerUser"`
	ZTPStatus string `json:"ZTPStatus"`
}
