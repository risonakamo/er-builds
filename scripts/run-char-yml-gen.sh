set -exu
HERE=$(dirname $(realpath $BASH_SOURCE))
cd $HERE

cd ../bin/char-yaml-gen
go run char_yaml_gen.go