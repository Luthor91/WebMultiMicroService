PROJECT_DIR := $(CURDIR)/services

init:
	cd $(PROJECT_DIR)/portal && go mod init portal
	cd $(PROJECT_DIR)/service1 && go mod init service1

portal:
	cd $(PROJECT_DIR)/portal && go mod tidy
	cd $(PROJECT_DIR)/portal && go build
	cd $(PROJECT_DIR)/portal && go run main.go&

s1:
	cd $(PROJECT_DIR)/service1 && go mod tidy
	cd $(PROJECT_DIR)/service1 && go build
	cd $(PROJECT_DIR)/service1 && go run main.go&

all: portal s1


.PHONY: setup build run all stop
