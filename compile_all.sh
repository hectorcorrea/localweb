# Compile the executable for Mac, Linux,and Windows
GOOS=darwin go build -o localweb
GOOS=linux go build -o localweb_linux
GOOS=windows GOARCH=386 go build -o localweb.exe
