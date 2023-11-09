module benchmarks/base64/base64-go

go 1.21

require (
	benchmarks/common v0.0.0-00010101000000-000000000000
	github.com/chenzhuoyu/base64x v0.0.0-20230717121745-296ad89f973d
)

require (
	github.com/bytedance/sonic v1.10.2 // indirect
	github.com/chenzhuoyu/iasm v0.9.1 // indirect
	github.com/klauspost/cpuid/v2 v2.2.6 // indirect
	golang.org/x/arch v0.6.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
)

replace benchmarks/common => ../../common/go
