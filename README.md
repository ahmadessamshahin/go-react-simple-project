# To Run project

    docker-compose up -d --build

# Request

### You have to create phone to have something to list as there is no seed

```
curl -X POST \             
  -H "Content-type: application/json" \
  -H "Accept: application/json" \
  -d '{"number":"(237) 236811111"}' \
  "http://localhost:4000/phone-number" 
```

```
curl -X GET \                               
  -H "Content-type: application/json" \
  -H "Accept: application/json" \
  "localhost:4000/phone-numbers?count=3&start=2"
```