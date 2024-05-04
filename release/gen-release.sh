# must have already done pnpm i
# creates release and places into output dir

set -exu
HERE=$(dirname $(realpath $BASH_SOURCE))
cd $HERE

# --- config
releaseName=er_builds-1.0.0


workdir=$HERE/workdir
outputdir=$HERE/output
topdir=$HERE/..

rm -rf $workdir
mkdir -p $workdir
mkdir -p $outputdir

cd ..

# build programs
go build -o server.exe bin/server/server.go
go build -o data-downloader.exe bin/data-downloader/data_download.go

# build web
cd er-builds-web
pnpm build


# construct output dir
cd $workdir
cp -r $topdir/config .
mkdir -p data
mkdir -p er-builds-web
cp -r $topdir/er-builds-web/build er-builds-web/
cp $topdir/data-downloader.exe .
cp $topdir/server.exe .
cp -r $topdir/doc/for-release/* .

rm -rf $outputdir/$releaseName
mv $workdir $outputdir/$releaseName

echo "done"