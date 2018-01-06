# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: grue android ios grue-cross swarm evm all test clean
.PHONY: grue-linux grue-linux-386 grue-linux-amd64 grue-linux-mips64 grue-linux-mips64le
.PHONY: grue-linux-arm grue-linux-arm-5 grue-linux-arm-6 grue-linux-arm-7 grue-linux-arm64
.PHONY: grue-darwin grue-darwin-386 grue-darwin-amd64
.PHONY: grue-windows grue-windows-386 grue-windows-amd64

GOBIN = $(shell pwd)/build/bin
GO ?= latest

grue:
	build/env.sh go run build/ci.go install ./cmd/grue
	@echo "Done building."
	@echo "Run \"$(GOBIN)/grue\" to launch grue."

swarm:
	build/env.sh go run build/ci.go install ./cmd/swarm
	@echo "Done building."
	@echo "Run \"$(GOBIN)/swarm\" to launch swarm."

all:
	build/env.sh go run build/ci.go install

android:
	build/env.sh go run build/ci.go aar --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/grue.aar\" to use the library."

ios:
	build/env.sh go run build/ci.go xcode --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/grue.framework\" to use the library."

test: all
	build/env.sh go run build/ci.go test

clean:
	rm -fr build/_workspace/pkg/ $(GOBIN)/*

# The devtools target installs tools required for 'go generate'.
# You need to put $GOBIN (or $GOPATH/bin) in your PATH to use 'go generate'.

devtools:
	env GOBIN= go get -u golang.org/x/tools/cmd/stringer
	env GOBIN= go get -u github.com/jteeuwen/go-bindata/go-bindata
	env GOBIN= go get -u github.com/fjl/gencodec
	env GOBIN= go install ./cmd/abigen

# Cross Compilation Targets (xgo)

grue-cross: grue-linux grue-darwin grue-windows grue-android grue-ios
	@echo "Full cross compilation done:"
	@ls -ld $(GOBIN)/grue-*

grue-linux: grue-linux-386 grue-linux-amd64 grue-linux-arm grue-linux-mips64 grue-linux-mips64le
	@echo "Linux cross compilation done:"
	@ls -ld $(GOBIN)/grue-linux-*

grue-linux-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/386 -v ./cmd/grue
	@echo "Linux 386 cross compilation done:"
	@ls -ld $(GOBIN)/grue-linux-* | grep 386

grue-linux-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/amd64 -v ./cmd/grue
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(GOBIN)/grue-linux-* | grep amd64

grue-linux-arm: grue-linux-arm-5 grue-linux-arm-6 grue-linux-arm-7 grue-linux-arm64
	@echo "Linux ARM cross compilation done:"
	@ls -ld $(GOBIN)/grue-linux-* | grep arm

grue-linux-arm-5:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-5 -v ./cmd/grue
	@echo "Linux ARMv5 cross compilation done:"
	@ls -ld $(GOBIN)/grue-linux-* | grep arm-5

grue-linux-arm-6:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-6 -v ./cmd/grue
	@echo "Linux ARMv6 cross compilation done:"
	@ls -ld $(GOBIN)/grue-linux-* | grep arm-6

grue-linux-arm-7:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-7 -v ./cmd/grue
	@echo "Linux ARMv7 cross compilation done:"
	@ls -ld $(GOBIN)/grue-linux-* | grep arm-7

grue-linux-arm64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm64 -v ./cmd/grue
	@echo "Linux ARM64 cross compilation done:"
	@ls -ld $(GOBIN)/grue-linux-* | grep arm64

grue-linux-mips:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips --ldflags '-extldflags "-static"' -v ./cmd/grue
	@echo "Linux MIPS cross compilation done:"
	@ls -ld $(GOBIN)/grue-linux-* | grep mips

grue-linux-mipsle:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mipsle --ldflags '-extldflags "-static"' -v ./cmd/grue
	@echo "Linux MIPSle cross compilation done:"
	@ls -ld $(GOBIN)/grue-linux-* | grep mipsle

grue-linux-mips64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64 --ldflags '-extldflags "-static"' -v ./cmd/grue
	@echo "Linux MIPS64 cross compilation done:"
	@ls -ld $(GOBIN)/grue-linux-* | grep mips64

grue-linux-mips64le:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64le --ldflags '-extldflags "-static"' -v ./cmd/grue
	@echo "Linux MIPS64le cross compilation done:"
	@ls -ld $(GOBIN)/grue-linux-* | grep mips64le

grue-darwin: grue-darwin-386 grue-darwin-amd64
	@echo "Darwin cross compilation done:"
	@ls -ld $(GOBIN)/grue-darwin-*

grue-darwin-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/386 -v ./cmd/grue
	@echo "Darwin 386 cross compilation done:"
	@ls -ld $(GOBIN)/grue-darwin-* | grep 386

grue-darwin-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/amd64 -v ./cmd/grue
	@echo "Darwin amd64 cross compilation done:"
	@ls -ld $(GOBIN)/grue-darwin-* | grep amd64

grue-windows: grue-windows-386 grue-windows-amd64
	@echo "Windows cross compilation done:"
	@ls -ld $(GOBIN)/grue-windows-*

grue-windows-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=windows/386 -v ./cmd/grue
	@echo "Windows 386 cross compilation done:"
	@ls -ld $(GOBIN)/grue-windows-* | grep 386

grue-windows-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=windows/amd64 -v ./cmd/grue
	@echo "Windows amd64 cross compilation done:"
	@ls -ld $(GOBIN)/grue-windows-* | grep amd64
