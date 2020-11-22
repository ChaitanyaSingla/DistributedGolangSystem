# Distributed Job Executor System with Golang

### To create TCP server, run the below command
go run .\main.go --createMaster

### To create a create and connect a client to the TCP, run the below command
go run .\main.go

### Command line arguments (Optional) which can be used for creating multiple clients
go run .\main.go --createMaster --port 8000 (Creates TCP server on Port 8000) 
go run .\main.go --masterIPAddress localhost:8000 (Creates a client which connects to the provided master IP address) 
