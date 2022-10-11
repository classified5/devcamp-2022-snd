# Backend Shipper Service for Tokopedia Devcamp 2022 Server and Database

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

Insert Shipper
```shell
curl --request POST \
  --url http://localhost:9000/shipper \
  --header 'Content-Type: application/json' \
  --data '{
	"shipper_name": "shipper1",
	"shipper_image": "shipper1.jpg",
	"shipper_description": "shipper1 is fast",
	"max_weight": 100
}'
```

Update Shipper By ID
```shell
curl --request PUT \
  --url http://localhost:9000/shipper/1 \
  --header 'Content-Type: application/json' \
  --data '{
	"shipper_name": "shipper1",
	"shipper_image": "shipper1.jpg",
	"shipper_description": "shipper1 is fast and efficient",
	"max_weight": 10
}'
```

Delete Shipper By ID
```shell
curl --request DELETE \
  --url http://localhost:9000/shipper/1
```

Get Shipper By ID
```shell
curl --request GET \
  --url http://localhost:9000/shipper/1
```

Get All Shipper
```shell
curl --request GET \
  --url http://localhost:9000/shippers
```

## Code Structure

Code structure for shipper service.

```
service
 ├── database             # Database Initialization Configuration
 ├── model                # Entity Definition
 ├── server               # Server Initialization Configuration
        ├── handlers      # HTTP Handler       
 ├── shippermodule        # Business Logic
 ├── main.go              # Service Initalization
 └── Dockerfile           # Dockerfile to build the image
```
