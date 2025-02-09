package main

import (
	"fmt"
	"sync-and-async-process-in-go/generator"
	"sync-and-async-process-in-go/struct_response"
	"sync-and-async-process-in-go/util"
)

func main() {
	util.SafeExecute(func() { // metodmain dijalanakna dalam method safeExecution untuk menghindari error yang tidak di inginlkan (Selain error bisnins)
		var success = struct_response.SucccessStruct{}
		success.Message = "Successfully to Get Data"
		success.Code = 200
		success.Data = generator.GeneratorSynchronous(15, 1000)
		fmt.Println(util.StructToJson(success))

		success.Message = "Successfully to Get Data"
		success.Code = 200
		success.Data = generator.GeneratorAsynchronous(15, 1000)
		fmt.Println(util.StructToJson(success))
	})

}
