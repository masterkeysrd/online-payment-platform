
.PHONY: init
init: deps

.PHONY: deps
deps:
	bash ./scripts/deps.sh

.PHONY: serve
serve:
	bash ./scripts/serve.sh