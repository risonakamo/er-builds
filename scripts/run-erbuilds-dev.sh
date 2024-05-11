# build and run server

set -exu
HERE=$(dirname $(realpath $BASH_SOURCE))
cd $HERE/..

go build -o erbuilds.exe bin/er-builds/er_builds.go
./erbuilds.exe