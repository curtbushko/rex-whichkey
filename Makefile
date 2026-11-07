# Task remains the single source of build and quality commands.
# PATH may provide a make wrapper that execs Task instead of reading this file.
.DEFAULT_GOAL := default
TASK ?= task
ARGS ?=

.PHONY: default build run clean install lint test tidy

default build clean install lint test tidy:
	$(TASK) $@

# GNU make: make run ARGS='width=0.8 height=0.3 ascii=true'
# The PATH Task wrapper instead accepts: make run -- width=0.8 ascii=true
run:
	$(TASK) run -- $(ARGS)

# Preserve access to named Task quality gates, e.g. make lint:vet.
%:
	$(TASK) $@
