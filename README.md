#Configuring DB2 in WIndows with Docker

1.	cd $GOPATH\src\github.com\prathaponnet\db2check

2.	Create Vendor libraries


	Command 1:	>go mod init
	Command 2:	>go mod vendor
		
	Check a new folder vendor created under $GOPATH\src\github.com\prathaponnet\db2check

3. delete the files go.mod	, go.sum before building

4. docker build .