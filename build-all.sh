# build all including web.
# web needs to already by pnpm i'd

set -exu
HERE=$(dirname $(realpath $BASH_SOURCE))
cd $HERE

go build -o erbuilds.exe bin/er-builds/er_builds.go
go build -o builds-downloader.exe bin/builds-downloader/builds_downloader.go

cd er-builds-web
rm -rf build
pnpm build