export DEFAULT_FILE=config/default.yml
export SETTING_FILE=plugins/MiraiAPIHTTP/exportting.yml
export FACE_MAP_FILE=config/face-map.yml

export WIN_AMD64=bin/main-windows-amd64.exe
export WIN_386=bin/main-windows-386.exe
export DARWIN_AMD64=bin/main-darwin-amd64.bin
export DARWIN_386=bin/main-darwin-386.bin
export LINUX_AMD64=bin/main-linux-amd64.bin
export LINUX_386=bin/main-linux-386.bin
export LINUX_ARM=bin/main-linux-arm.bin

if [ -d $DARWIN_AMD64 ];then
go run src/main/main.go $1 $DEFAULT_FILE
else
./$DARWIN_AMD64 $1
fi