go clean ./...                                      
go clean -cache    
go clean -modcache

sudo rm -rf ~/go.path/pkg/mod/github.com/briancabbott
sudo rm -rf ~/go.path/pkg/mod/entgo.io

go generate -v -x ./...
