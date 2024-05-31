set -exu
HERE=$(dirname $(realpath $BASH_SOURCE))
cd $HERE/..

go build -o nica-downloader.exe bin/nica-downloader/nica_downloader.go

set +u
if [[ "$1" == "run" ]]; then
    ./nica-downloader.exe
fi