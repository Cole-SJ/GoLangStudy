# Chapter 13 구조체

---

`chapter 13`에서는 Go 언어에서의 구조체(`struct`)에 대해서 설명한다.

1. **구조체의 선언과 변수 사용**
2. **구조체 특징**

---

## 구조체의 선언과 변수 사용

### 구조체이란?

- `구조체(structure)`은 여러 필드를 묶어서 하나의 객체처럼 사용할 수 있게 해준다. 다른 타입들을 변수 하나로 묶어준다.

### 구조체의 선언

- `구조체`는 type 키워드를 명시하여, 새로운 사용자 정의 타입을 정의하는 것을 알린 뒤, 타입명을 명세한다.
- 타입명의 첫번째 글자가 대문자면 외부 패키지로 공개되는 타입이다.
- 일반 변수 선언과 같이 사용자가 정의한 타입을 통해 변수를 선언하고, 내부 필드들을 아래와 같이 `.`을 통해 접근/할당 할 수 있다.

```go
/*
* type [타입명(구조체 이름)] struct {
*   [필드명] [타입]
*    ...
*   [필드명] [타입]
* }
*
*/

type Student struct {
  Name string
  Class int
  No int
  Score float64
}

var a Student
a.Name = "Jermaine"
a.Class = 2
a.No = 13
a.Score = 16.4
```

### 구조체의 초기화 방법

```go
type House struct {
    Address string
    Size int
    Price float64
    Type string
}
```

- 구조체 변수에 대한 초기화 방법에는 여러가지 방법이 있다.(ex `House` 타입)

1. 초깃값 생략
- 모든 필드가 기본값으로 초기화 된다.(`Address`: "" , `Size`: 0, `Price`: 0.0, `Type`: "")

```go
var house House 
```

1. 모든 필드 초기화
- 모든 필드가 선언 순서대로 내부 필드에 할당되다.

```go
var house House = House {"Address Field", 28, 9.80, "주택" }

//or

var hose House = House {
    "Address Field",
    28,
    9.80,
    "주택", //마지막에 쉼표를 추가해야함.
}
```

1. 일부 필드 초기화
- 일부 필드만 초기화하는데 `"필드명: 필드값"` 형태로 대괄호 안에 추가하면되며, 초기화 대상이 아닌 필드는 기본값으로 초기화된다.

```go

var house House = House { Size: 28, Price: 9.80 }

//or

var hose House = House {
    Size: 28,
    Price: 9.80, //마지막에 쉼표를 추가해야함.
}
```

## 구조체의 특징

### 구조체를 포함하는 구조체

```go
type User struct {
    Name string
    ID string
    Age int
}

type VIPUser struct {
    UserInfo User
    VipLevel int
    Price int

```

- 구조체는 필드로서, 다른 구조체를 포함할 수 있다.
- 포함하는 방식에는 두 가지가 있다.
1. 내장타입처럼 포함하기
    - 그냥 단순 변수 처럼 사용하면 된다.
    
    ```go
    func main() {
            user := User { "김철수", "KR kim", 32 }
            vip := VIPUser {
                user,
                2,
                2500,
            }
    
          fmt.Printf("%s", vip.UserInfo.Name)
    }
    ```
    
2. 포함된 필드 방식(`Embedded Field`)
    - 1 번 방식에서는 User에 접근하기 위해서는 `.`을 두 번 사용해서 접근해야한다.
        - java에서 상속과 같은 느낌으로 한번에 해당 필드에 접근하기 위해서 `포함된 필드 방식`을 사용할 수 있다.
        - 필드 이름이 겹칠 경우에는 현재 접근중인 제일 바깥 타입의 필드에 접근하고, 중첩된 필드에 접근하기 위해서는 해당 구조체 필드에 별도 접근이 필요하다.(예시참고)

```go
type User struct {
    Name string
    ID string
    Age int
    Level int
}

type VIPUser struct {
    User //별도 필드가 아닌 구조체 자체를 포함시킨다.
    Level int
    Price int
}
user := User { "김철수", "KR kim", 32, 12}
vip := VIPUser {
    user,
    2,
    2500,
}

fmt.Printf("%s", vip.Name)
fmt.Printf("%d, %d", vip.Level, vip.User.Level)
```

### 구조체의 크기

- 구조체는 해당 필드를 담을 수 있는 메모리 공간을 할당받는다.
- 구조체의 크기는 두가지 값을 고려하여 결정된다.
    1. 필드들의 타입별 메모리 크기의 합
    2. 필드 선언 순서에 따른 패딩: cpu 레지스터가 4바이트씩(32비트 cpu), 8바이트씩(64비트 cpu) 읽는지에 따라서 메모리 정렬(`Memory Alignment`)가 발생하기 때문에 메모리 패딩이 발생 가능
