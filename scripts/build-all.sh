# build all including web.
# web needs to already by pnpm i'd

set -exu
HERE=$(dirname $(realpath $BASH_SOURCE))
cd $HERE

bash downloader-dev.sh
bash erbuilds-dev.sh
bash nica-downloader-dev.sh
bash langfile-download-dev.sh

cd ..
cd er-builds-web
rm -rf build
pnpm build