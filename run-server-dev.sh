# build and run server

set -exu
HERE=$(dirname $(realpath $BASH_SOURCE))
cd $HERE

go build -o server.exe bin/server/server.go
./server.exe