# build and run data download tool

set -exu
HERE=$(dirname $(realpath $BASH_SOURCE))
cd $HERE/..

go build -o builds-downloader.exe bin/builds-downloader/builds_downloader.go
./builds-downloader.exe