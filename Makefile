build:
	go build -o check *.go

try:
	cat ./diff-example.txt | ./check