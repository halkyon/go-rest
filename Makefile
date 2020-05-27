BUILD_PATH ?= ./build
PLATFORMS ?= linux/amd64 darwin/amd64 windows/amd64
OS = $(word 1, $(subst /, ,$@))
ARCH = $(word 2, $(subst /, ,$@))
BUILD_FILE = app-$(OS)-$(ARCH)

build: $(PLATFORMS)
$(PLATFORMS):
	@echo "Building $(BUILD_FILE)"
	@mkdir -p $(BUILD_PATH)
	@CGO_ENABLED=0 GOOS=$(OS) GOARCH=$(ARCH) go build -trimpath -ldflags="-w -s" -o $(BUILD_PATH)/$(BUILD_FILE)
	# optional: use upx to make resulting binaries even smaller!
	#@upx -qqq -9 $(BUILD_PATH)/$(BUILD_FILE)
	@cd $(BUILD_PATH) && md5sum $(BUILD_FILE) >> ./md5sums

clean:
	@rm -rf $(BUILD_PATH)/*

.PHONY: build $(PLATFORMS)