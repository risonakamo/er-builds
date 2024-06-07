# build langfile download tool. give "run" to run

set -exu
HERE=$(dirname $(realpath $BASH_SOURCE))
cd $HERE/..

go build -o lang-downloader.exe bin/lang-file-download/lang_file_download.go

set +u
if [[ "$1" == "run" ]]; then
    ./lang-downloader.exe
fi