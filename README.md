# GO Ride

A fare estimation that will use as input a list of tuples of the form (id_ride, lat, lng, timestamp) representing the position of the taxi-cab during a ride, filter the invalid entries and export a text file following the format: (id_ride fare_estimate)


## Getting Set Up

Before running the application, you will need to ensure that you have a few requirements installed;
You will need Go.

### Go

[Go](https://golang.org/) is an open source programming language that makes it easy to build simple, reliable, and efficient software.

## Project Structure

### `/cmd`

Main applications for this project.

### `/internal`

Internal application logic

### `/deployment`

Deployment configurations and templates such as docker-compose

### `/scripts`

Scripts used for several purposes such as building/running application, etc

### `/test`

Additional external tests and test data.

### `/output`

Default folder where the result file will be created.

## Running

    go run .\cmd\go-ride-fare-estimation\main.go -fp {FILE_PATH} -rfp {RESULT_FILE_PATH}

or

    go build -o go-ride-fare-estimation -i ./cmd/go-ride-fare-estimation/main.go

    ./go-ride-fare-estimation -fp {FILE_PATH} -rfp {RESULT_FILE_PATH}


Usage of go-ride-fare-estimation:

    -fp string
        file path of the rides positions. (default "test/testdata/paths.csv")
    -rfp string
        file path of the fare estimation results. (default "output/result.csv")


[@maxalencar](https://github.com/maxalencar)