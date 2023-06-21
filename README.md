
# Price Fetcher In GRPC

This microservice is Created for Fetch Price from Database using GRPC.

```
var priceMocker = map[string]float64{
	"BTC": 20_000.0,
	"ETH": 200.0,
	"MSR": 100_00.0,
}

```
Use this Command

Step 1:-

```
make

```

Step 2:

```
make run
```

Open Brower and Type this

```
http://localhost:3000/?ticker=MSR

```

