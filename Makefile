USER = techjacker
REPO = systemdlogger
COMMIT_ID = 94c52865e3b449ca594d09995b99efc28a24c53d

RULES_DIR = test/fixtures/rules
RULES_URL = https://raw.githubusercontent.com/michenriksen/gitrob/master/signatures.json

SINGLE_TEST = "TestSplitDiffs"
# SINGLE_TEST = "Test_ScanDiffs"

install:
	@go install -race ./cmd/diffence

build:
	@go build -race ./cmd/diffence

# https://github.com/dominikh/go-tools/
lint:
	@golint  -set_exit_status ./...
	@go vet ./...
	@interfacer $(go list ./... | grep -v /vendor/)
	@gosimple ./...
	@staticcheck ./...
	@unused ./...

rules:
	@curl -s $(RULES_URL) > $(RULES_DIR)/gitrob.json

diff:
	@curl -s https://api.github.com/repos/$(USER)/$(REPO)/commits/$(COMMIT_ID) \
		-H "Accept: application/vnd.github.VERSION.diff"

test:
	@go test

test-single:
	@go test -test.run $(SINGLE_TEST)

test-cover:
	@go test -cover

test-race:
	@go test -race


.PHONY: test* run diff rules lint install build
