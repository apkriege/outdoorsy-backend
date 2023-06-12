# outdoorsy-backend

Simple GO API that returns rentals based on an optional query string.

### Endpoints
`/rentals/:id`
`/rentals/near=33.64,-117.93&price_min=9000&price_max=75000&limit=3&offset=6&sort=price`
  
### Running Tests
`go test test/rentals_test.go`
  
### Starting the App
Starting and seeding the postgres database
`docker-compose up`

Running the app
`go run main.go`
