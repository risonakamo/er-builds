# clean all data and build products

set -exu
HERE=$(dirname $(realpath $BASH_SOURCE))
cd $HERE

cd ..
rm -rf data/*
touch data/keep.txt
rm *.exe