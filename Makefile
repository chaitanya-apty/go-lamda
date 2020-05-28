linux:
	GOOS=linux go build -o build/main main.go
package:
	zip function.zip build/main