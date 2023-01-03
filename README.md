### To rebuild vue files

```bash
cd pkg/http/web/app
yarn run build
```

Files will be updated in /pkg/http/web/app/dist

### To run Vue app only (without Go)
port 3000
```bash
yarn serve
```
### To run Go app
port 8000

```bash
cd pkg/http/web/app
yarn run build # build vue files
go run main.go
```
