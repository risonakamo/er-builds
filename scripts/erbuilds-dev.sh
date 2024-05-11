# build er builds app. give "run" command to run

set -exu
HERE=$(dirname $(realpath $BASH_SOURCE))
cd $HERE/..

go build -o erbuilds.exe bin/er-builds/er_builds.go

set +u
if [[ "$1" == "run" ]]; then
    ./erbuilds.exe
fi