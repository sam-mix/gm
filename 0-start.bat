@ECHO OFF

cd server
start go run main.go

sleep 2

cd ../web
start npm run serve
