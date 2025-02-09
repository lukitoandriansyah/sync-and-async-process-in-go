package generator

import (
	"fmt"
	"sync"
	"sync-and-async-process-in-go/struct_model"
	"sync-and-async-process-in-go/util"
	"time"
)

func GeneratorSynchronous(lengthOfString int, totalIteration int) struct_model.DetailGenerateStruct {
	fmt.Println("Generator String Synchronous")
	return DoGeneratorStringSynchronous(lengthOfString, totalIteration)
}

func DoGeneratorStringSynchronous(lengthOfString int, totalIteration int) struct_model.DetailGenerateStruct {
	startTimeSynchronous := time.Now()
	var resultDetailgenerateString struct_model.DetailGenerateStruct = struct_model.DetailGenerateStruct{}
	resultgenerate, totalStringGenerate := GeneratorStringSynchronous(lengthOfString, totalIteration)
	resultDetailgenerateString.TypeGenerate = "Synchronous"
	resultDetailgenerateString.TotalGenerate = totalStringGenerate
	resultDetailgenerateString.TotalGenerateSuccess = len(resultgenerate)
	resultDetailgenerateString.ResultGenerateString = resultgenerate
	resultDetailgenerateString.TimeTotal = fmt.Sprintf("%d ms", time.Since(startTimeSynchronous).Milliseconds())
	return resultDetailgenerateString
}

func GeneratorStringSynchronous(lengthOfString int, totalIteration int) ([]struct_model.ResultGenerateStruct, int) {
	var resultGeneratorStringList []struct_model.ResultGenerateStruct = []struct_model.ResultGenerateStruct{}
	var resultGeneratorString struct_model.ResultGenerateStruct = struct_model.ResultGenerateStruct{}
	for start := 0; start < totalIteration; start++ {
		resultString, errString := util.SecureRandomString(lengthOfString)
		if errString != nil {
			continue
		}
		resultGeneratorString.ResultGenerate = resultString
		resultGeneratorString.TimeGenerate = time.Now()

		resultGeneratorStringList = append(resultGeneratorStringList, resultGeneratorString)
	}
	return resultGeneratorStringList, totalIteration
}

func GeneratorAsynchronous(lengthOfString int, totalIteration int) struct_model.DetailGenerateStruct {
	fmt.Println("Generator String Asynchronous")
	return DoGeneratorStringAsynchronous(lengthOfString, totalIteration)
}

func DoGeneratorStringAsynchronous(lengthOfString int, totalIteration int) struct_model.DetailGenerateStruct {
	startTimeAsynchronous := time.Now()

	resultDetailgenerateString := struct_model.DetailGenerateStruct{}
	resultgenerate, totalStringGenerate := GeneratorStringAsynchronous(lengthOfString, totalIteration)

	resultDetailgenerateString.TypeGenerate = "Asynchronous"
	resultDetailgenerateString.TotalGenerate = totalIteration
	resultDetailgenerateString.TotalGenerateSuccess = totalStringGenerate
	resultDetailgenerateString.ResultGenerateString = resultgenerate
	resultDetailgenerateString.TimeTotal = fmt.Sprintf("%d ms", time.Since(startTimeAsynchronous).Milliseconds())

	return resultDetailgenerateString
}

func GeneratorStringAsynchronous(lengthOfString int, totalIteration int) ([]struct_model.ResultGenerateStruct, int) {
	// runtime.GOMAXPROCS(runtime.NumCPU()) // Uncomment in ijika memang inigiin diatur manual jumlah CPU yanhg digunakan
	// runtime.GOMAXPROCS(0) // Uncomment in ijika memang inigiin diatur manual jumlah CPU yanhg digunakan

	var resultGeneratorStringList []struct_model.ResultGenerateStruct
	var wg sync.WaitGroup
	resultChannel := make(chan struct_model.ResultGenerateStruct, totalIteration)

	// Menjalankan goroutine untuk setiap iterasi
	for i := 0; i < totalIteration; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resultString, errString := util.SecureRandomString(lengthOfString)
			if errString != nil {
				return
			}
			resultChannel <- struct_model.ResultGenerateStruct{
				ResultGenerate: resultString,
				TimeGenerate:   time.Now(),
			}
		}()
	}

	// Menutup channel setelah semua goroutine selesai
	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	// Mengumpulkan hasil dari channel
	for result := range resultChannel {
		resultGeneratorStringList = append(resultGeneratorStringList, result)
	}

	return resultGeneratorStringList, len(resultGeneratorStringList)
}
