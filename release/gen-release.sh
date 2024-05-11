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

cd $topdir/scripts

bash build-all.sh


# construct output dir
cd $workdir
cp -r $topdir/config .
mkdir -p data
mkdir -p er-builds-web
cp -r $topdir/er-builds-web/build er-builds-web/
cp $topdir/builds-downloader .
cp $topdir/erbuilds.exe .
cp -r $topdir/doc/for-release/* .

rm -rf $outputdir/$releaseName
mv $workdir $outputdir/$releaseName

echo "done"