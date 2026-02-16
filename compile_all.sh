# Compile the executable for Mac, Linux,and Windows
GOOS=darwin GOARCH=arm64 go build -o localweb_arm64
GOOS=darwin GOARCH=amd64 go build -o localweb_amd64
GOOS=linux go build -o localweb_linux
GOOS=windows GOARCH=386 go build -o localweb.exe

# Create the universal Mac executable
lipo -create -output localweb localweb_arm64 localweb_amd64