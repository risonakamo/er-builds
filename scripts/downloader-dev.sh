# build download tool. give "run" to also run

set -exu
HERE=$(dirname $(realpath $BASH_SOURCE))
cd $HERE/..

go build -o builds-downloader.exe bin/builds-downloader/builds_downloader.go

if [[ "$1" == "run" ]]; then
    ./builds-downloader.exe
fi