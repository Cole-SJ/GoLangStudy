# Chapter 9, 10 흐름 제어문

---

`chapter 9, 10`에서는 Go 언어에서 조건을 통해서 어떻게 프로그램의 흐름을 제어하는지 설명한다. 책에서는 두 챕터로 나눠 설명하는데, 어차피 흐름을 제어하는 것은 동일한 목적을 가지고 있기 때문에, 한 글에 몰아서 작성한다. 해당 두 챕터에서는 아래의 내용을 주로 설명한다.

1. **if 문 , if…else 조건문**
2. **조건문 중첩**
3. **if 초기문**
4. **switch 문**
5. **switch 초기문**
6. **break와 fallthrough**
    
    ---
    

## if 문, if…else 조건문

- 프로그램에서 특정한 조건일 때, 혹은 특정한 조건이 아닐 때, 각각 다른 로직을 실행시켜야하는 경우가 있다. 이 때, `if` 키워드를 통해 분기문을 작성할 수 있다.
- 형식은 아래와 같으며 `조건문`의 결과가 `true`일 때,  코드 블록 안의 코드를 실행한다. 무수히 많은 `else if` 문을 추가할 수 있고, `else` 와 `else if`는 필요하지 않다면, 생략 가능하다.
    
    ```go
    if 조건문 1 {
    	statement 1
    } else if 조건문 2 {
    	statement 2
    } else {
      statement 3
    }
    
    //예시
    package main
    
    import "fmt"
    
    func main(){
    
    	temp := 33
    
    	if temp > 28 {
    		fmt.Println("에어컨을 켠다.")
    	} else if temp <= 3 {
    		fmt.Println("히터를 켠다.")	
    	} else {
    		fmt.Println("대기 한다.")
    	}
    }
    ```
    
- `&&(AND)`, `||(OR)` 연산자를 통해 여러 조건을 한번에 검사할 수 있고, `short-circuit`을 사용하여, 조건의 연산을 더욱 빠르게 처리가 가능하다.
    - `short-circuit`
        - `condition 1 && condition 2` 에서 `condition 1`이 `false`라면 `condtion 2`를 검사하지 않는다.
        - `condition 1 || condition 2`에서 `condition 1`이 `true`라면 `condition 2`를 검사하지 않는다.

> ***하지만, 우리는 코드의 가독성도 프로그램의 품질에 많은 영향을 미치기 때문에, 최대한 여러 조건을 검사하는 상황을 피하고, 최대한 괄호를 통해 가독성을 높이는 노력이 필요하다.***
> 

---

## 조건문 중첩

- 말 그래도 `if`문 안에 다른 `if`문을 사용하는 것이다. 특히, 사람의 언어 중에서는 단일 `if`으로 처리하기 어려운 상황들이 존재한다.
- 그러한 경우에는 조건문을 중첩하여 사용할 수 있다.
    
    ```go
    if( a.IsEmpty() && b.IsNotEmpty()){
    
    	if(c.IsEmpty() || d.IsNotEmpty()){
    	 ...
    	}
    
    } else{
    
    	if(c.IsEmpty() || d.IsNotEmpty()){
    		 ...
    		}
    }
    ```
    

> ***다만 중첩이 심할 경우에는 코드의 가독성을 크게 저하시키는 요인이 될 수 있으므로, 3중첩 이상은 피하자. 제발.***
> 

---

## if 초기문

- 다른 언어에서 지원하지 않는 신기한 기능이다. 처음봐서 매우 신기하다.
- 특정 함수의 결과를 통해 조건문을 실행시키고 싶은 경우가 있다. 아래와 같이, `if`문 조건을 검사하기 전에 초기문을 넣을 수 있다. 초기문은 검사에 사용할 변수를 초기화할 때 주로 사용한다.
    
    ```go
    if 초기문; 조건문 {
    	statement
    }
    
    // 예시
    
    if filename, success := UploadFile(); success {
    	fmt.Println("Upload success", filename)
    } else {
    	fmt.Println("Failed to upload")
    }
    ```
    
- ***단, 초기문 안에서 선언한 변수는 해당 조건문 안에서만 사용할 수 있으니 주의해야한다.변하면 안되는 값에 상수를 사용하자.***

---

## switch 문

- 조건이 많고 다양할 때는 `switch` 문을 통해 각각의 조건에 맞는 프로그램 로직을 수행시켜 분기 처리할 수 있다. 이는 복잡한 `if else` 문을 정리하는 기능을 한다.
- 주로 `const` 열거값에 따라 수행되는 로직을 변경할 때, 주로 사용한다.
    
    > ***하지만 열거값이 추가, 삭제됨에 따라 switch문을 함께 수정해야하는 Strong coupling 문제가 발생한다. 이러한 문제는 산탄총 수술 문제라고 하며, 그렇기에 하나의 열거값에 연관된 switch 문은 최대한 적게 사용하는 것이 좋다.***
    > 
- `switch`문은 다른 프로그래밍 언어와 구조가 유사하지만, **차이점은 `break`문을 작성하지 않아도, 아래의 정의된 `case`에 해당하는 코드가 실행되지 않는다는 것이다. 즉, 해당하는 조건의 `case`의 코드만 실행된다는 것이다.**
    
    ```go
    switch 비굣값 {
    	case 값1:
    		statement 1
    	case 값2:
    		statement 2
    	case 값3:
    		statement 3
    	case 값4, 값5:
    		statement 4
    	default:
    		statement 5
    }
    
    //예시
    package main
    
    import "fmt"
    
    func main() {
    	a := 3
    
    	switch a {
    
    	case 1:
    		fmt.Println("1")
    	case 2:
    		fmt.Println("2")
    	case 3:
    		fmt.Println("3")
    	case 4, 5:
    		fmt.Println("4 or 5")
    	default:
    		fmt.Println("a > 5")
    	}
    }
    ```
    
- 각 `case`에 선언된 값과 비굣값을 비교하고 일치하는 `case`의 코드 만을 실행시키며, 한번에 여러 값을 `OR` 조건으로 비교할 수 있다.
- 조건에 부합하는 `case`가 없을 경우에는 `default`의 코드를 실행시킨다.
- `switch` 문을 아래와 같이도 작성할 수 있다. 또한 `switch` 키워드 옆의 `true`는 생략할 수 있다.
    
    ```go
    package main
    
    import "fmt"
    
    func main() {
    
    	temp := 18
    
    	switch true {
    
    	case temp < 10, temp > 30:
    		fmt.Println("바깥 활동 하기 좋은 날씨가 아닙니다.")
    	case temp >= 10 && temp < 20:
    		fmt.Println("약간 추울 수 있으니 가벼운 겉옷을 준비하세요")
    	case temp >= 15 && temp < 25:
    		fmt.Println("야외 활동하기 좋은 날씨입니다..")
    	default:
    		fmt.Println("따뜻합니다.")
    
    	}
    
    }
    ```
    

---

## switch 초기문

- `if` 문과 마찬가지로 아래와 같이, `switch` 문에서도 초기문을 사용할 수 있고, 초기문에서 선언된 변수는 해당 조건문 안에서만 사용 가능하다.
- 또한, `switch` 문에서도 `true`가 생략 가능했으니, `switch` 초기문에서도 필요한 경우 비굣값의 `true` 는 생략 가능한 건 당연지사다.
    
    ```go
    
    switch 초기문; 비굣값 {
    
    	case 값1:
    		...
    	case 값2:
    		.
    		.
    		.
    	default:
    	...
    
    }
    
    //예시
    package main
    
    import "fmt"
    
    func getMyAge() int {
    	return 22
    }
    
    func main() {
    	switch age := getMyAge(); age {
    	case 10:
    		fmt.Println("TeenAge")
    	case 33:
    		fmt.Println("Pair 3")
    	default:
    		fmt.Println("My age is ", age)
    	}
    
    	//fmt.Println("age, is". age)
    }
    ```
    

---

## **break와 fallthrough**

- 위에서 언급했듯이, 다른 언어와 다르게 Go에서 `switch`문의 `case`문은 `case`문 종료 시에 `break`문을 사용하지 않아도, 딱 실행된 `case`문에서 자동적으로 `switch`문을 빠져나온다.
- 역으로, 다른 언어와 같이 하나의 `case`문을 수행한 뒤, 다음 `case`문을 수행하고 싶다면, `fallthrough` 키워드를 사용하면된다.
    
    > ***하지만, `fallthrough`를 사용하는 것은 가독성에 혼란을 주기 때문에 사용하지 않는 것을 권장한다.***
    > 
    

---

흐름제어문은 다른 언어와 매우 유사한 성격을 가지고 있다. 다만 초기문은 생소하기 때문에, 한 번 코드를 직접 작성해보면서 연습해보는 것을 추천한다.
