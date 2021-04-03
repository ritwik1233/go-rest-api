(sudo kill -9 `sudo lsof -t -i:9001` && sudo mongod --fork --logpath /var/log/mongod.log --port 9001 --dbpath /home/ritwik/data/db) || sudo mongod --fork --logpath /var/log/mongod.log --port 9001 --dbpath /home/ritwik/data/db
cd models && go build *.go
cd ../
cd handlers && go build *.go
cd ../
cd handlermethods && go build *.go
cd ../
cd dev && go build *.go
cd ../

ENV="DEV" go run main.go 
# & cd client && npm start
