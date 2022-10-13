# Backend Product Service for Tokopedia Devcamp 2022 Server and Database

You can run the service using docker detached option

```shell
docker-compose up -d
```

If you want to rebuild the service, run this command

```shell
docker-compose build
```

# HTTP REST API

Root Handler
```shell
curl --request GET \
  --url http://localhost:9000/
```

Insert Product
```shell
curl --request POST \
  --url http://localhost:9000/product \
  --header 'Content-Type: application/json' \
  --data '{
	"product_name": "product1",
	"product_image": "product1.jpg",
	"product_description": "product1 is fast",
	"max_weight": 100
}'
```

Update Product By ID
```shell
curl --request PUT \
  --url http://localhost:9000/product/1 \
  --header 'Content-Type: application/json' \
  --data '{
	"product_name": "product1",
	"product_image": "product1.jpg",
	"product_description": "product1 is fast and efficient",
	"max_weight": 10
}'
```

Get Product By ID
```shell
curl --request GET \
  --url http://localhost:9000/product/1
```

Get All Product
```shell
curl --request GET \
  --url http://localhost:9000/products
```

## Code Structure

Code structure for product service.

```
service
 ├── database             # Database Initialization Configuration
 ├── model                # Entity Definition
 ├── server               # Server Initialization Configuration
        ├── handlers      # HTTP Handler       
 ├── productmodule        # Business Logic
 ├── main.go              # Service Initalization
 └── Dockerfile           # Dockerfile to build the image
```
