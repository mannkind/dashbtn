go test -v ./...

go test -coverprofile=handlers/.coverprofile github.com/mannkind/dashbtn/handlers
gover . .coverprofile
go tool cover -html=.coverprofile 
find . -name ".coverprofile" -exec rm {} \;
