//go:build tinygo.wasm

// Code generated by protoc-gen-go-plugin. DO NOT EDIT.
// versions:
// 	protoc-gen-go-plugin v0.1.0
// 	protoc               v3.21.5
// source: tests/fields/proto/fields.proto

package proto

import (
	context "context"
	wasm "github.com/knqyf263/go-plugin/wasm"
)

const FieldTestPluginAPIVersion = 1

//export field_test_api_version
func _field_test_api_version() uint64 {
	return FieldTestPluginAPIVersion
}

var fieldTest FieldTest

func RegisterFieldTest(p FieldTest) {
	fieldTest = p
}

//export field_test_test
func _field_test_test(ptr, size uint32) uint64 {
	b := wasm.PtrToByte(ptr, size)
	var req Request
	if err := req.UnmarshalVT(b); err != nil {
		return 0
	}
	response, err := fieldTest.Test(context.Background(), req)
	if err != nil {
		return 0
	}

	b, err = response.MarshalVT()
	if err != nil {
		return 0
	}
	ptr, size = wasm.ByteToPtr(b)
	return (uint64(ptr) << uint64(32)) | uint64(size)
}