docker run --name postgre12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine


https://github.com/lib/pq driver used for connection to talk to driver



docker exec -it postgres12 psql -U root -d simple_bank

https://github.com/begmaroman/go-micro-boilerplate