
.PHONY: init
init: deps

.PHONY: deps
deps:
	bash ./scripts/deps.sh

.PHONY: fmt
fmt:
	bash ./scripts/fmt.sh

.PHONY: serve
serve:
	bash ./scripts/serve.sh