.PHONY: build
build:
	bazel build ...

.PHONY: test
test:
	bazel test ... --runs_per_test=30

.PHONY: tidy
tidy:
	go mod tidy
	bazel run update_go_repositories
	bazel run gazelle

