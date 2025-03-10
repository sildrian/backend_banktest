# backend_banktest
Backend test using golang

# Step to run application
1) makesure already install go and postgresql
2) Import postgresql database above
3) run, go mod tidy (to install package)
4) run, go run main.go

# To run unit test
- after do preperation above, to run unit test do as below
- dont forget to enable database config in customer.service.go (for unit test running well)
- run, go test .\app\customers\controllers\customers.controller_test.go
