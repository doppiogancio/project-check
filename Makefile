build:
	go build -o check *.go

try:
	cat ./example.txt | ./check