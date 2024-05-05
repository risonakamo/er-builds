# build all including web.
# web needs to already by pnpm i'd

set -exu
HERE=$(dirname $(realpath $BASH_SOURCE))
cd $HERE

go build -o server.exe bin/server/server.go
go build -o data-downloader.exe bin/data-downloader/data_download.go

cd er-builds-web
pnpm build