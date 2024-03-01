# Midas Case Study API - File Operations

This case study API demonstrates how to make various file operations with following operations:
- Shredding File
- Filling File with binary data
- Dumping File data into local sqlite database

## How to Run Locally

Run `go run ./cmd/case-study-api` in the project directory.

## Endpoints 

### HealthCheck `GET /api/v1/health`
Simple endpoint to make healthcheck of the app.

### Shred File `DELETE /api/v1/file/shred/{filePath}`
Shreds the given file with overwriting the contents of the file. 

#### Validations 
- File path should be always **.txt** otherwise API will return **400**

#### Example 

**`DELETE /api/v1/file/shred/test.txt`**

### Fill File with Binary Data `POST /api/v1/file/fill`
Fills the given '.txt' file with the binary data. 
Binary data should follow the format mentioned below.

#### Validations 
- File path should be always **.txt** otherwise API will return **400**
- Each binary data sequence should be 8 bits long to match ASCII character schema
- Each binary data sequence should be separated by **','**
- Binary data should be at least 10 lines
- New lines should be added as adding **'\n'** at the end of the binary data sequence

#### Example

**`POST /api/v1/file/fill`**

**Body:**
```
{
    "FilePath": "test.txt",
    "FileData": "01001000,01100101,01101100,01101100\n00110011,00110011\n00110011,00110001\n00110011,00110001\n01001000,01100101,01101100,01101100\n01001000,01100101,01101100,01101100\n01001000,01100101,01101100,01101100\n01001000,01100101,01101100,01101100\n01001000,01100101,01101100,01101100\n01001000,01100101,01101100,01101100\n01001000,01101100,01101100,01101100,01001000,01100101,01101111"
}
```

### Dump file data into sqlite db `POST /api/v1/file/dump`
Dumps the given '.txt' binary data filled file into local sqlite database. 
Entities will be created under 'file' table with the following schema:

```
Table: File
- Column: Name (File Name)
- Column: Content (Nth line of Binary File Content)
```

Number of entities will always be equeal to number of lines that binary file has. 

#### Validations 
- File path could be sent as empty to use latest binary file data that has been modified with fill-data opreation
- File path should be either **.txt** or empty otherwise APII will return **400**

#### Example - 1

**`POST /api/v1/file/dump`**

**Body:**
```
{
    "FilePath": "test.txt",
}
```

#### Example - 2

**`POST /api/v1/file/dump`**

**Body:**
```
{
    "FilePath": "",
}
```

## Discussion

### Shred
Such operation could be risky to expose to the end users since file that shouldn't be shredded could be subject to this operation by exposing with API.
Other disadvantage is that operation could take too long to finish if the given file is too big. Shred operation is a simple O(N) operation without any concurency in place. 
Advantage for the operation is that users can easily shred the content of the given file with a simple API call and user will make sure that it's not recoverable operation.

### Dump
For this operation, it could be tricky to deal with large files because of creating a bulk insert query for each line within a single variable. 
This operation could potentially cause a memory leak issue if the given file is too big. To mitigate this issue, bulk insert could be executed with batches using the concurrency. 

### Db Schema
Ideally, schema should have 1 to N relation with 2 tables rather than having single table. 
1 table could be responsible for the file itself and other could have N(number of lines inside the file) which is a more ideal approach for storing the data. 

### Tests

Unit tests are only limited to handler functions and it could be extended to the API handlers as well.
Additionally file operations could be potentially mocked like the sqlite repository to increase the coverage and edge case scenarios within unit tests.


