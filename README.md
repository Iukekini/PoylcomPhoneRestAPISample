# Sample Application for Poylcom Phone API

The included code is a sample of how to connect to the Phone REST API on a Polycom VVX Phone. 

## The Prerequisites

Here's what you'll need to get started:

 * Polycom UC Softare 5.8.0 for any of the following VVX models: (101, 150, 201, 250, 300, 301, 310, 311, 350, 400, 401, 410, 411, 450, 500, 501, 600, 601 and 1500)
 * Rest Api Enabled in Preferences
 * Password changed from the default 456
 
 
## Build and Run Sample

Replace the IPAddress with the IPAddress of the Phone you want to test with. 

```bash
git clone https://github.com/Iukekini/PoylcomPhoneRestAPISample.git
cd PoylcomPhoneRestAPISample
go build
./PoylcomPhoneRestAPISample -ip=192.168.1.100 -username=Polycom -password=***
```
