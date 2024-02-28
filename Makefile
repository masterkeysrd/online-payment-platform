
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

.PHONY: build
simulator_start:
	bash ./scripts/simulator_start.sh

.PHONY: app_start
app_start:
	bash ./scripts/app_start.sh

.PHONY: app_stop
app_stop:
	bash ./scripts/app_stop.sh

.PHONY: app_clean
app_clean:
	bash ./scripts/app_clean.sh

.PHONY: services_start
services_start:
	bash ./scripts/services_start.sh

.PHONY: services_stop
services_stop:
	bash ./scripts/services_stop.sh

.PHONY: services_clean
services_clean:
	bash ./scripts/services_clean.sh