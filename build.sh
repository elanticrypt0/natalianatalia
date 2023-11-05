GOOS=windows GOARCH=amd64 go build -ldflags "-w -s"

GOOS=linux GOARCH=amd64 go build -ldflags "-w -s"

# move to the build folder

rm -rf build

mkdir build
mkdir build/tangas
mkdir build/scripts

cp natalianatalia build
cp natalianatalia.exe build
cp appconfig.toml build
cp _db -R build
cp config -R build/config
cp public -R build/public
cp seeds -R build/seeds
