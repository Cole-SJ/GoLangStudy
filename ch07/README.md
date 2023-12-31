# Chapter 7 함수
***
`chapter 7`에서는 Go 언어에서 함수를 사용하는 방법에 대해서 설명한다.
1. **함수 선언**
2. **멀티 반환**
3. **재귀 함수**
***

## 함수
* `함수`는 반복적으로 자주 사용되는 코드를 묶어서 관리하는 것으로, 이를 통해 코드의 재사용성, 가동성 및 유지보수성을 높일 수 있다.
* `메서드`는 클래스 인스턴스 혹은, 구조체 인스턴스 내부에 선언된 `함수`를 의미한다.
* `함수`의 정의는 프로세스 메모리 구조 내에서, 코드영역에 저장되어 있다.
    *호출 시, 해당 함수은 스택영역의 스택프레임에 필요한 지역변수와 매개변수를 선언하여 동작한다.

### 함수 선언

함수는 아래와 같은 구성요소로 이루어져있다.
```go
func Add(a int, b int) int {
    return a + b
}
```
1. `함수 키워드`: `func`
2. `함수 명`: `Add`, 첫글자가 대문자인 함수는 패키지 외부로 공개되는 함수고, 소문자인 함수는 패키지 내부에서 사용되는 함수다. 함수의 일종인 `메서드`는 객체 지향 관점에서 객체의 행위를 의미하므로, 동사를 사용해야한다.
3. `매개변수`: `(a int, b int)`, 함수의 입력 값으로, 필요없다면 괄호를 비워둘 수 있다.
4. `반환 타입`: `int`, 함수의 반환 타입으로 비어있으면, 함수가 값을 반환하지 않는 다는 것을 의미한다.
5. `함수 코드 블록`: `{ return a + b }`, 실제 함수의 실행문을 담당하는 코드 블록이다.

### 함수의 매개변수(Parameter)와 인자(Argument)
* 둘 다 함수와 같은 프로시저의 입력값을 지칭하는 것으로, 함수 내부에서 바라본 전달 입력 값을 매개변수, 함수 외부에서 바라본 전달 값 인자라고 생각하면 된다.
    * 즉, 함수 입장에서는 인자를 매개변수로 받아서 함수의 동작을 수행한다고 생각하면 된다.
* **함수가 호출될 때는, 인자를 매개변수로 복사하여 전달한다.** 위에 선언된 `Add` 함수를 예시로 들어 생각하면, 만약에 `Add(9,3)`이 호출된다면, 함수는 `9`를 `a`에 복사하고, `3`을 `b`에 복사한 뒤, `a`와 `b`를 이용해서 함수가 동작한다고 생각하면 된다.
* 함수는 스택영역에 `스택프레임`이라는 형태의 실행단위로 묶여서 수행된다. 스택은 `LIFO`형 자료구조로, `Push`를 통해 input을 `Pop`을 통해 output을 발생시킨다. 만약에 `main` 함수의 `스택프레임`의 수행 중간에 `foo`라는 함수가 호출된다면,
`foo` 함수의 `스택프레임`이 바로 위에 `Push`되고, 실행이 마무리 되면 해당 함수의 `스택프레임`이 `Pop`이 되는 구조를 가지고 있다.

***

## 멀티 반환 함수
* Go 언어에서 사용되는 함수는 아래와 같이, 한번에 여러 개의 값을 반환할 수 있다.
     * 여기서 `Divide` 함수는 아래와 같은 방식으로 동작하는 것을 알 수 있다.

```go

package main

import (
"fmt"
)

func Divide(a, b int) (int, bool) {
    if b == 0 {
        return 0, false
    }
    return a / b, true
}

func main() {
    c, success := Divide(9, 3)
    fmt.Println("c, success =", c, success)
    d, success := Divide(9, 0)
    fmt.Println("d, success =", d, success)
}
```


    

```go
c , success := Divide(9, 3)

c, success := 3, true
```

* 또한, 함수의 반환값으로 사용될 변수를 함수 선언 부에서 아래의 코드와 같이 미리 지정하여 반환할 수 있다.
    * 주석 처리된 함수와 동일하게 동작한다.
```go
package main

import "fmt"

func Divide(a, b int) (result int, success bool) {
    if b == 0 {
        result = 0
        success = false
        return
    }

    result = a / b
    success = true
    return
}

/*
func Divide(a, b int) (result int, success bool) {

    var (
        result int
        success bool
    )

    if b == 0 {
        result = 0
        success = false
        return
    }

    result = a / b
    success = true
    return
}
*/

func main() {
    c, success := Divide(9, 3)
    fmt.Println("c, success =", c, success)
    fmt.Println("&success =", &success)

    d, success := Divide(9, 0)
    fmt.Println("d, success =", d, success)
    fmt.Println("&success =", &success)
}
```

<br>

> * 여기서 센스가 있다면 굳이 `success`의 메모리 주소를 출력한 이유에 대해서 궁금할 수 있다. 위 코드에서 `success`의 메모리 주소는 같다.
> * `:=` 연산은 분명히 `선언`과 `초기화`를 동시에 수행하는 연산인 것으로 알았는데, 이미 success 변수가 선언되었는데, 또 한번 선언을 해도 오류가 발생하지 않는다.
> * **결론적으로, `:=`로 여러 변수를 동시에 할당할 때, 만약 기 선언된 변수가 섞여 있다면, `선언` 과정은 생략되고, `할당`만 수행한다.**
>     * 그렇기 때문에 아래 코드는 오류가 발생하지 않는다.
>```go
>
>var a = 3
>f, a := 2, 4
>fmt.Println("a,f=", a, f)
>```
>

***

## 재귀 함수

```
어느 한 컴퓨터공학과 학생이 유명한 교수님을 찾아가 물었다.

"재귀함수가 뭔가요?"
"잘 들어보게. 옛날옛날 한 산 꼭대기에 이세상 모든 지식을 통달한 선인이 있었어.
마을 사람들은 모두 그 선인에게 수많은 질문을 했고, 모두 지혜롭게 대답해 주었지.
그의 답은 대부분 옳았다고 하네.
그런데 어느날,그 선인에게 한 선비가 찾아와서 물었어.

"재귀함수가 뭔가요?"
"잘 들어보게. 옛날옛날 한 산 꼭대기에 이세상 모든 지식을 통달한 선인이 있었어.
마을 사람들은 모두 그 선인에게 수많은 질문을 했고, 모두 지혜롭게 대답해 주었지.
그의 답은 대부분 옳았다고 하네. 그
런데 어느날, 그 선인에게..........
```
그냥 이거다. 함수 본인이 함수 실행문에서 스스로 본인을 호출하는 함수다. 재귀함수는 아래와 같이 함수 본인이 본인을 호출하는 함수로 **반드시 탈출 조건이 있어야한다. 그렇지 않으면 콜스택이 꽉차서 오류가 발생한다.**
```go
package main

import "fmt"

func printNo(n int) {
    if n == 0 {
        return
    }

    fmt.Println("n =", n)
    printNo(n - 1)
    fmt.Println("After", n)
}

func main() {
    printNo(3)
}

```
> 개인적으로 재귀함수를 선호하지 않는다.
> 디버깅이 어렵고, 성능이 좋지 않기 때문이다.
