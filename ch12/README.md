# Chapter 12 배열

---

`chapter 12`에서는 Go 언어에서 배열의 개념과 사용법을 설명한다.

1. **************************배열의 선언**************************
2. ******************************배열값 접근******************************
3. **배열 순회**

---

## 배열의 선언

### 배열이란?

- `배열(Array)` 은 **같은 타입**의 데이터들이 **메모리 공간에 연속적으로 존재**하는 데이터 타입이다.
    - ***같은 타입이 메모리 공간에 연속적으로 존재한다는 것에 집중해야한다.***
    - `배열`을 이루는 같은 타입의 각각의 값들을 `원소(Element)`라고 한다.
    - 각 `원소`가 위치한 순서를 `인덱스(Index)`라고 하며, 순서는 메모리 공간에서의 맨 처음을 기준으로 한다.
        - **맨 처음 `인덱스`를 0으로 기준 잡는다.**
- `배열`은 같은 타입의 여러 데이터들을 하나의 변수로 효과적으로 사용할 수 있게 해준다.
- `배열`은 연속된 메모리공간에 같은 타입의 데이터들을 저장한다고 했다.
    - 만약 길이가 10인 `int32`의 `배열`이 있다고 한다면, 총 `10 * 4바이트` = `40바이트`의 메모리 공간을 차지한다. 또한, `3번째 인덱스`의 `원소`의 주솟값은 `배열`의 `시작 주소(0번째 원소의 주소)` + `4(0~3) * 4바이트` = `배열의 시작 주소` + `16바이트의 위치`에 해당 `원소`가 존재한다.

### 배열의 선언 및 초기화 방법

- `배열` 변수 선언의 기본적인 형태는 아래와 같이 한다.
    
    ```go
    /*
    * var 변수명 [요소 개수]타입 = [요소 개수]타입{ 원소들 }
    *
    * or
    *
    * 변수명 := [요소 개수]타입 { 원소들 }
    */
    
    var t [5]float64 {1.2, 2.3, 3.4, 4.5, 5.6}
    
    //or
    
    t := [5]float64 {1.2, 2.3, 3.4, 4.5, 5.6}
    ```
    
    - 이는  `t`라는 변수에 5개 짜리 `float64` 타입을 가진 배열의 원소를 `1.2`, `2.3`, `3.4`, `4.5`, `5.6` 로 선언 및 초기화를 하였다는 것을 의미한다.
    - **`원소`의 개수는 무조건 상수여야만 한다, 즉, 변수로 `배열`의 개수를 선언해서는 안된다.**
- `원소`의 초기화 없이 아래와 같이 선언할 수 있다.
    
    ```go
    /*
    * var 변수명 [요소 개수]타입 
    *
    * or
    *
    * 변수명 := [요소 개수]타입
    */
    
    var t [5]float64 
    
    //or
    
    t := [5]float64 
    ```
    
    - 이렇게 선언을 한다면, `t`의 각 5개의 `원소`에는 변수의 기본값인 `0.0`으로 초기화가 된다.
        - `배열`의 `원소` 타입에 따라 초기화가 되는 기본값이 존재한다.
- `{ }`를 활용하여 개발자가 원하는 `인덱스` 만을 초기화 할 수 있다.
    
    ```go
    /*
    * var 변수명 [요소 개수]타입 { 인덱스:값, ... , 인덱스:값}
    *
    * or
    *
    * 변수명 := [요소 개수]타입 { 인덱스:값, ... , 인덱스:값 }
    */
    
    var t [5]float64 { 1:3.15, 3: 4.13 }
    
    //or
    
    t := [5]float64 { 1:3.15, 3: 4.13 }
    ```
    
    - `1번째 인덱스` 에 위치한 `원소`가 `3.15`로 초기화되고, `3번째 인덱스`에 위치한 `원소`가 `4.13`로 초기화 된다.
- `[ ... ]`를 활용하여 선언 시, `배열`의 `원소` 개수를 생략할 수 있다.
    
    ```go
    /*
    * var 변수명 [...]타입 { 원소들 }
    *
    * or
    *
    * 변수명 := [...]타입 { 원소들 }
    */
    
    var t [...]float64 { 2, 3, 4 }
    
    //or
    
    t := [...]float64 { 2, 3, 4 }
    ```
    
    - 자동적으로 3개의 길이를 가진 `배열`이 선언 및 초기화 된다.

### 배열의 복사

- 대입 연산자를 사용하여 `배열` 대 `배열`을 복사할 수 있다.
- 대입 연산자는 우변의 값을 좌변의 메모리 공간에 복사하는 연산임을 기억하자.
    - 즉, `배열`의 시작 `원소`의 주소를 복사하는 것과 같다.
    - 하지만 대입 시, 서로의 데이터 타입이 맞지 많으면, 다른 변수의 대입 연산과 마찬가지로 오류가 발생한다.
    
    ```go
    var a[5] int
    var b[5] int
    
    b = a
    ```
    

---

## 배열값 접근

- 기본적으로 `배열`에 대한 접근은 `인덱스`를 통해 이루어진다. `대괄호 [ ]` 를 작성하고 접근하고자 하는 `원소`의 `인덱스`를 `대괄호` 안에 적는다.
- **배열값 읽기**
    
    ```go
    var a [3]int64 = [3]int64 { 1, 2, 3 }
    
    fmt.Println(a[2])
    // 3
    ```
    

- **배열값 쓰기**
    
    ```go
    var a [3]int64 = [3]int64 { 1, 2, 3 }
    
    a[2] = 300
    fmt.Println(a[2])
    // 300
    ```
    

---

## 배열 순회

- `반복문`을 통해 `배열` 순회할 수 있다. **(연속적인 메모리 공간에 저장된다는 점을 상기하자.)**
    - `len()` 함수를 통해 `배열` 순회
        
        ```go
        for i := 0; i < len(nums); i++ {
        	fmt.Println(nums[i])
        }
        ```
        
        - `len` 함수는 `배열`의 길이를 반환하는 함수다.
    
    - `range` 키워드를 통해 `배열` 순회
        
        ```go
        for i, v := range nums {
        	fmt.Println(i, v)
        }
        ```
        
        - `range` 키워드는 `iterable`한 변수의 `원소`에 하나씩 접근하며, `인덱스`와 `원소`의 값을 반환한다. 즉 위의 코드에서는 `i`에는 `인덱스`를 `v`에는 `원소`의 값을 반환하며 `반복문`을 동작시킨다.
        

---

## Appendix. 다중 배열

- `다중 배열`은 아래와 같이 중첩된 `배열`을 의미한다. **선형대수에서 나오는 행렬을 생각하면 편하다.**
    - 주로 `이중 배열`은 `x, y 좌표계`와 같이 평면 위치 데이터들을 표현하기 편하며, `삼중 배열`은  `x, y, z 좌표계`와 같이 공간 위치 데이터들을 표현하기에 편리한 경향이 있다.
    - `이중, 삼중 배열`보다 더욱 `배열`을 중첩한 `다중 배열` 또한 가능한데, 이는 `배열`의 `원소`로 다른 `배열`이 존재하며, 이러한 구조가 반복될 수 있음을 뜻한다.
    - 단, 아무리 `다중 배열`이라도 `배열`의 정의가 변하는 것은 아니다.
        - `**배열`은 메모리공간에 같은 타입의 데이터가 연속적으로 저장된 형태를 의미하며, 프로그래머가 보기편한 형태로 `다중 배열`을 구성할 뿐, 모든 `원소`는 메모리 공간에 연속적인 주소 값을 가지며 존재하게 된다.**
    
    ```go
    
    //이중배열의 선언 및 초기화
    var a [2][3]int = [2][3]int { 
    	{1, 2, 3},
    	{4, 5, 6},
    }
    
    //삼중배열의 선언 및 초기화
    var b [2][3][4] int = [2][3][4]int {
    	{
    		{1, 2, 3, 4}, 
    		{4, 5, 6, 7}, 
    		{7, 8, 9, 10},
    	{
    		{10, 11, 12, 13},
    		{13, 14, 15, 16},
    		{16, 17, 18, 19},
    	},
    } 
    ```