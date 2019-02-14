if [ ! -d "dist_linux" ]; then
    mkdir dist_linux
else
    rm -rf dist_linux
    mkdir dist_linux
fi

echo "go build now . . ."
go build


cp server dist_linux
rm server

cp -R Config dist_linux

rm dist_linux/Config/.gitkeep

echo "build complate!"