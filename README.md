# sync-and-async-process-in-go
This implementation processes both Synchronous and Asynchronous methods in Golang. The objective is to compare the total time taken to generate random strings with a length of 15 and a total count of 1,000.

# 🛠Installation & Usage
- Clone this repository:
```plaintext
git clone https://github.com/lukitoandriansyah/sync-and-async-process-in-go.git
cd sync-and-async-process-in-go
```
- Run the application:
```plaintext
go run main.go
```

# 📂Project Structure

```plaintext
sync-and-async-process-in-go/
│── generator/
│   ├── generator_string.go        # Implements synchronous and asynchronous string generation
│── struct_model/
│   ├── detail_generate_struct.go  # Defines the structure for generation details
│   ├── result_generate_struct.go  # Defines the structure for generated results
│── struct_response/
│   ├── error_struct.go            # Defines error response structures
│   ├── success_struct.go          # Defines success response structures
│── util/
│   ├── error_handling.go          # Error handling utilities
│   ├── helper.go                  # Helper functions
│── go.mod                         # Go module configuration
│── main.go                         # Entry point of the application
```
# 🤝 Contribution
If you’d like to contribute, feel free to fork this repository and submit a pull request.
