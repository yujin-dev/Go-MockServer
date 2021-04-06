# GO-server 

## HTTP REST API
- configuration.json 읽어서 HTTP 로 전송
- Server 2대 connection status 전송

## Socket 
- server 
- client
hello 코드 출처: https://kamang-it.tistory.com/entry/golanggotcpgo%EC%96%B8%EC%96%B4%EC%97%90%EC%84%9C-tcp%EC%86%8C%EC%BC%93%EC%9C%BC%EB%A1%9C-%ED%86%B5%EC%8B%A0%ED%95%98%EA%B8%B0

## Usage
```go
func main() {
	runType("socket")
}
```
`socket`으로 server를 실행함( `rest`로 설정 가능)

### 실행
```$ go run main.go ```
