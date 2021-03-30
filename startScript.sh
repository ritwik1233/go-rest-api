(sudo kill -9 `sudo lsof -t -i:9001` && sudo mongod --fork --logpath /var/log/mongod.log --port 9001 --dbpath /home/ritwik/data/db) || sudo mongod --fork --logpath /var/log/mongod.log --port 9001 --dbpath /home/ritwik/data/db
go build main.go
./main