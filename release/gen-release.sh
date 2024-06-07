# must have already done pnpm i
# creates release and places into output dir

set -exu
HERE=$(dirname $(realpath $BASH_SOURCE))
cd $HERE

# --- config
releaseName=er_builds-1.1.0
# --- end config


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

cp $topdir/builds-downloader .
cp $topdir/erbuilds.exe .
cp $topdir/nica-downloader.exe .
cp -r $topdir/doc/for-release/* .
cp $topdir/version.md .

mkdir -p config
cp -r $topdir/config/chars.yml config
cp -r $topdir/config/saved-langfile.txt config

mkdir -p data

mkdir -p er-builds-web
cp -r $topdir/er-builds-web/build er-builds-web/

rm -rf $outputdir/$releaseName
mv $workdir $outputdir/$releaseName

echo "done"