PROJECT_DIR := $(CURDIR)/services

init:
	cd $(PROJECT_DIR)/portal && go mod init portal
	cd $(PROJECT_DIR)/service1 && go mod init service1

setup_portal:
	cd $(PROJECT_DIR)/portal && go mod tidy

build_portal:
	cd $(PROJECT_DIR)/portal && go build

run_portal:
	cd $(PROJECT_DIR)/portal && go run main.go > portal.log 2>&1 & echo $$! > portal.pid

setup_s1:
	cd $(PROJECT_DIR)/service1 && go mod tidy

build_s1:
	cd $(PROJECT_DIR)/service1 && go build

run_s1:
	cd $(PROJECT_DIR)/service1 && go run main.go > service1.log 2>&1 & echo $$! > service1.pid

portal: setup_portal build_portal run_portal

s1: setup_s1 build_s1 run_s1

all: portal s1

stop:
	@kill `cat portal.pid` || true
	@kill `cat service1.pid` || true
	@rm -f portal.pid service1.pid

.PHONY: setup build run all stop
