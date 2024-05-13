# GO Ride Fare Estimation

A fare estimation that will use as input a list of tuples of the form (id_ride, lat, lng, timestamp) representing the position of the taxi-cab during a ride, filter the invalid entries and export a text file following the format: (id_ride fare_estimate)


## Getting Set Up

Before running the application, you will need to ensure that you have a few requirements installed;
You will need Go.

### Go

[Go](https://golang.org/) is an open source programming language that makes it easy to build simple, reliable, and efficient software.

## Project Structure

### `/cmd`

Main application for this project.

### `/internal`

Internal application logic

### `/test`

Additional e2e test and test data.

### `/output`

Default folder where the result file will be created.

## Running

    go run .\cmd\ride\main.go -fp {FILE_PATH} -rfp {RESULT_FILE_PATH}

or

    go build -o app -i ./cmd/ride/main.go

    ./app -fp {FILE_PATH} -rfp {RESULT_FILE_PATH}


Usage of go-ride-fare-estimation:

    -fp string
        file path of the rides positions. (default "test/testdata/paths.csv")
    -rfp string
        file path of the fare estimation results. (default "output/result.csv")


[@maxalencar](https://github.com/maxalencar)