# build and run data download tool

set -exu
HERE=$(dirname $(realpath $BASH_SOURCE))
cd $HERE

go build -o data-downloader.exe bin/data-downloader/data_download.go
./data-downloader.exe