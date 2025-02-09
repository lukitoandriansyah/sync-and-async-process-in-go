package util

import (
	"fmt"
	"runtime/debug"
	"sync-and-async-process-in-go/struct_response"
)

// Wrapper function untuk menangkap panic
func SafeExecute(fn func()) {
	defer func() {
		if r := recover(); r != nil {
			var error = struct_response.ErrorStruct{}
			error.Message = fmt.Sprintf("Terjadi error: %v", r)
			error.Code = 500
			error.Data = string(debug.Stack())
			fmt.Println(StructToJson(error))
			return
		}
	}()
	fn()
}
