# 기본 변수 설정
BINARY_NAME=image_converter
BUILD_DIR=build

# 기본 환경 변수
CGO_ENABLED=1
GO=go

# 기본 타겟 (현재 시스템에 맞는 빌드)
.PHONY: build-current
build-current:
	CGO_ENABLED=$(CGO_ENABLED) $(GO) build -o $(BUILD_DIR)/$(BINARY_NAME)

# Mac M1/M2/M3 (arm64)
.PHONY: build-mac-arm
build-mac-arm:
	CGO_ENABLED=$(CGO_ENABLED) GOOS=darwin GOARCH=arm64 $(GO) build -o $(BUILD_DIR)/$(BINARY_NAME)_mac_arm64

# Mac Intel (amd64)
.PHONY: build-mac-intel
build-mac-intel:
	CGO_ENABLED=$(CGO_ENABLED) GOOS=darwin GOARCH=amd64 $(GO) build -o $(BUILD_DIR)/$(BINARY_NAME)_mac_amd64

# Windows (amd64)
.PHONY: build-windows
build-windows:
	CGO_ENABLED=$(CGO_ENABLED) GOOS=windows GOARCH=amd64 $(GO) build -o $(BUILD_DIR)/$(BINARY_NAME)_windows.exe

# 빌드 디렉토리 생성
$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

# 정리
.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)

# 현재 시스템 정보 출력
.PHONY: info
info:
	@echo "GOOS: $$(go env GOOS)"
	@echo "GOARCH: $$(go env GOARCH)"
	@echo "CGO_ENABLED: $(CGO_ENABLED)"

# 도움말
.PHONY: help
help:
	@echo "사용 가능한 명령어:"
	@echo "  make build-current		- 현재 시스템에 맞는 버전 빌드"
	@echo "  make build-mac-arm  	- Mac M1/M2/M3용 빌드"
	@echo "  make build-mac-intel 	- Mac Intel용 빌드"
	@echo "  make build-windows  	- Windows용 빌드"
	@echo "  make clean         	- 빌드 파일 정리"
	@echo "  make build         	- 빌드 파일 생성"
	@echo "  make info          	- 현재 시스템 정보 출력"