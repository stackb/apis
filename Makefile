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

.PHONY: update
update:
	$(BAZEL) run //build/stack/bzl/v1beta:v1beta1_go_compiled_sources.update
	$(BAZEL) run //build/stack/starlark/info/v1beta1:v1beta1_go_compiled_sources.update
	$(BAZEL) run //build/stack/protobuf/compiler/v1alpha1:v1alpha1_go_compiled_sources.update
	$(BAZEL) run //build/stack/protobuf/package/v1alpha1:v1alpha1_go_compiled_sources.update
	$(BAZEL) run //build/stack/protobuf/package/v1alpha2:v1alpha2_go_compiled_sources.update
	$(BAZEL) run //build/stack/syntax_highlight/v1beta1:syntax_highlight_go_compiled_sources.update
