BAZEL=bzl 

.PHONY: build
build:
	$(BAZEL) build ...

.PHONY: test
test:
	$(BAZEL) test ... --runs_per_test=30

.PHONY: tidy
tidy:
	go mod tidy
	$(BAZEL) run update_go_repositories
	$(BAZEL) run gazelle

