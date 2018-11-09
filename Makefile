
all: clean test build-dot

test: 
	go test

build-dot:
	dot -o test.png -T png test.dot

clean:
	rm -rf *.dot *.png
