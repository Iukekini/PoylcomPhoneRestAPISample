# Sample Application for Poylcom Phone API

The included code is a sample of how to connect to the Phone REST API on a Polycom VVX Phone. Currently the sample only gets the network info from the device. You can see the full range of functions from the Polycom API Docs. 

https://community.polycom.com/t5/VoIP-SIP-Phones/FAQ-REST-API/td-p/98914

## The Prerequisites

Here's what you'll need to get started:

 * Polycom UC Software 5.8.0 for any of the following VVX models: (101, 150, 201, 250, 300, 301, 310, 311, 350, 400, 401, 410, 411, 450, 500, 501, 600, 601 and 1500)
 * Rest Api Enabled in Preferences
 * Password changed from the default 456
 
 
## Build and Run Sample


```bash
git clone https://github.com/Iukekini/PoylcomPhoneRestAPISample.git
cd PoylcomPhoneRestAPISample
go build
./PoylcomPhoneRestAPISample -ip=192.168.1.100 -username=Polycom -password=***
```
## Sample Result

```bash
./PoylcomPhoneRestAPISample -ip=10.0.1.100 -username=Polycom -password=***
INFO[07-31|23:13:30] Attempting to connect to phone via:      https://10.0.1.100/api/v1/mgmt/network/info=nil
INFO[07-31|23:13:31] Successfully called phone API
{
  "Status": "2000",
  "data": {
    "DHCP": "enabled",
    "DHCPServer": "10.0.1.1",
    "DHCPBootServerUserOption": "",
    "DHCPBootServerOption": "160",
    "VLANID": "",
    "DefaultGateway": "10.0.1.1",
    "IPV4Address": "10.0.1.100",
    "LANSpeed": "1000MB",
    "DHCPBootServerOptionType": "String",
    "DNSDomain": "",
    "CDPCompability": "enabled",
    "DHCPOption60Format": "ASCII String",
    "IPV6Address": "::",
    "SNTPAddress": "north-america.pool.ntp.org",
    "VLANDiscoveryMode": "Fixed",
    "SubnetMask": "255.255.255.0",
    "LANPortStatus": "active",
    "ProvServerType": "FTP",
    "DNSServer": "10.0.1.1",
    "UpgradeServer": "",
    "AlternateDNSServer": "0.0.0.0",
    "VLANIDOption": "129",
    "LLDP": "enabled",
    "ProvServerAddress": "boot.internalphone.org",
    "ProvServerUser": "pcsysmet",
    "ZTPStatus": "disabled"
  }
```
## Things to Note

 * Per the instrutions on the Polycom API Documention The sample disables checking any certificates. If your phones are loaded with a valid cert this won't be an issue. 
