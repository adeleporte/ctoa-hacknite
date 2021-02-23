mkdir %APPDATA%\terraform.d\plugins\vmware.com\edu\ctoa\0.1\windows_amd64
go build -o terraform-provider-ctoa.exe
move terraform-provider-ctoa.exe %APPDATA%\terraform.d\plugins\vmware.com\edu\ctoa\0.1\windows_amd64