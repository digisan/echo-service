# echo-service

template for golang echo micro-service via HTTP/HTTP2

1. `cd ./server`
2. `./cert.sh` if set HTTP2
3. `./swagger/swag init` if not existing, download from 'https://github.com/swaggo/swag/releases'
4. `go build`
5. check api paths in www/js in line with reg_api.go. 
