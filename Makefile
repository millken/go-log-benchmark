GOTEST_FLAGS=-cpu=1,2,4 -benchmem -benchtime=5s

TEXT_PKGS=Gokit Logrus Zerolog GoLog Zap
JSON_PKGS=Gokit Logrus Zerolog GoLog Zap

TEXT_PKG_TARGETS=$(addprefix test-text-,$(TEXT_PKGS))
JSON_PKG_TARGETS=$(addprefix test-json-,$(JSON_PKGS))

.PHONY: all deps test test-text test-json benchmark $(TEXT_PKG_TARGETS) $(JSON_PKG_TARGETS)

all: deps test 

deps:
	go mod download
	go install golang.org/x/perf/cmd/benchstat@latest

test: test-text test-json

test-text: $(TEXT_PKG_TARGETS)

$(TEXT_PKG_TARGETS): test-text-%:
	go test $(GOTEST_FLAGS) -bench "$*.*Text"

test-json: $(JSON_PKG_TARGETS)

$(JSON_PKG_TARGETS): test-json-%:
	go test $(GOTEST_FLAGS) -bench "$*.*JSON"

bench:
	go test -v -benchmem -run=^$$ -bench=. ./... > benchmark.txt

benchmark: 
	go test -bench=. | tee ./graphic/out.dat ; 
	awk '/Benchmark/{count ++; gsub(/BenchmarkTest/,""); printf("%d,%s,%s,%s\n",count,$$1,$$2,$$3)}' ./graphic/out.dat > ./graphic/final.dat ; \
	gnuplot -e "file_path='./graphic/final.dat'" -e "graphic_file_name='./graphic/operations.png'" -e "y_label='number of operations'" -e "y_range_min='000000000''" -e "y_range_max='100000000'" -e "column_1=1" -e "column_2=3" ./graphic/performance.gp ; \
	gnuplot -e "file_path='./graphic/final.dat'" -e "graphic_file_name='./graphic/time_operations.png'" -e "y_label='each operation in nanoseconds'" -e "y_range_min='000''" -e "y_range_max='5000'" -e "column_1=1" -e "column_2=4" ./graphic/performance.gp ; \
	rm -f ./graphic/out.dat ./graphic/final.dat ; \
	echo "'graphic/operations.png' and 'graphic/time_operations.png' graphics were generated."

benchx:
	benchx-cli run Benchmark

update-benchmark:
	go test -bench . -benchmem | go run cmd/benchstat/main.go
	cat readme_template.md benchmark_result.md > README.md