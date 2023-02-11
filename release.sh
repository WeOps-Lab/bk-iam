RELEASE_PATH=/opt/release/bkiam
rm -Rf $RELEASE_PATH

mkdir -p $RELEASE_PATH
make build

cp ./VERSION $RELEASE_PATH
cp -Rf ./build/* $RELEASE_PATH
cp -Rf ./release.md $RELEASE_PATH
mkdir $RELEASE_PATH/bin
cp -Rf ./iam $RELEASE_PATH/bin
