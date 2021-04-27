# CSV to JSON in Go

Takes a CSV file and outputs JSON either to sdtout or a specified output file.
Optionally use `-d` flga for alternative delimeter (tab or pipe)

## Usage

Data to sdtout
`go run main.go [input_file]`

Data to output file
`go run main.go [input_file] [output_file]`

With optional delimeter
`go run main.go -d [tab | pipe] [input_file] [output_file]`
