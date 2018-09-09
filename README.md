# localweb
A very simple web server for static files.

Use this program to serve HTML, JavaScript, or CSS files on your local machine for test purposes. I use this when I need to quickly test a particular HTML/JavaScript sample without having to worry about setting up Apache or Nginx.

This program will serve via HTTP any files under same folder as the executable.

## Running the executable
Download the executable for your operating system from the [Releases tab](https://github.com/hectorcorrea/localweb/releases), run it, and point your browser to http://localhost:9001

You can change the listening port via the `-port` parameter, for example:

```
./localweb -port 8080
```

## Running the code
```
git clone ...
cd localweb
go build
./localweb
```
