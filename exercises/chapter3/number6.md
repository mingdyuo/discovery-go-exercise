## 입출력 함수 공부하기

### 인터페이스 `Read`, `ReadFull`의 차이

[여기](https://mingrammer.com/translation-go-walkthrough-io-package/) 를 정리했다.

### 스트림에서 바이트를 읽기 위한 기본 구조체
`Reader` 인터페이스
```go
type Reader interface {
    Read(p []byte) (n int, err error)    
}
```

