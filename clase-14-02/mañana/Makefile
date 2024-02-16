# execute the tests
.PHONY tests:
tests: 
	go test -v ./...

# generate the coverage report
.PHONY coverage:
coverage: 
	go test -coverprofile=coverage.out ./...

# get coverage report in html format
.PHONY coverage_html:
coverage_html:
	go tool cover -html=coverage.out

# get total coverage
.PHONY total_coverage:
total_coverage:
	go tool cover -func="coverage.out"

# run the linter
.PHONY linter:
linter:
	staticcheck ./...