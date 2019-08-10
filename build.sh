#! /bin/bash
mkdir dist
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags '-w -s'
tar -zcvf cmdp.windows.v3.0.0.tar.gz cmdp.exe
rm -f cmdp.exe
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags '-w -s'
tar -zcvf cmdp.mac.v3.0.0.tar.gz cmdp
rm -f cmdp
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w -s'
tar -zcvf cmdp.linux.v3.0.0.tar.gz cmdp
rm -f cmdp
CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -ldflags '-w -s'
tar -zcvf cmdp.linux32.v3.0.0.tar.gz cmdp
rm -f cmdp
mv cmdp.windows.v3.0.0.tar.gz dist/cmdp.windows.v3.0.0.tar.gz
mv cmdp.mac.v3.0.0.tar.gz dist/cmdp.mac.v3.0.0.tar.gz
mv cmdp.linux.v3.0.0.tar.gz dist/cmdp.linux.v3.0.0.tar.gz
mv cmdp.linux32.v3.0.0.tar.gz dist/cmdp.linux32.v3.0.0.tar.gz
echo 'success'