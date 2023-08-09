# Chapter 4 변수
***
`chapter 4`에서는 Go 언어의 변수가 가지는 속성과 사용법에 대해서 설명한다.
크게 아래의 주제로 설명한다.
1. 변수의 선언 방법 (Declaration for variables)
2. 변수의 속성 (Attributes of variables)
3. 타입 변환 (Type conversion)
4. 숫자 표현 (Numeric expression)

***

## Go언어에서 변수란?
변수는 말 그대로 변할 수 있는 값을 의미한다. 컴퓨터 프로그램은 값들을 메모리 위에 적재한 상태로 연산을 수행하는데, 이때 변수라는 단위로 메모리에 적재하여 연산을 수행한다.<br>
그렇기 때문에 변수는 메모리 상에 어디에 위치해있는지(주소), 해당 주소에서 얼마만큼의 크기의 메모리 공간을 할당받아 차지하는지(크기)에 대한 정보를 가지고 있다.<br><br>
* 변수의 주소는 실제 변수가 메모리 상에 어느 주소에 위치해있는지 나타내는 값으로, 프로세스가 메모리를 할당 받았을 때, Data 영역, Text 영역, Stack 영역, Heap 영역 중에서 Stack 영역에 저장된다. 
* 변수의 크기는 변수의 타입에 따라 다르다.

### 변수의 선언 방법
<br>
프로그래머가 변수를 접근하고 변경하고 사용하기 위해서는 아래와 같이 선제적으로 **선언과 할당(기본값이라도) 과정**이 필요하다.<br>
```java
public class BootSpringBootApplication {
  public static void main(String[] args) {
    System.out.println("Hello, Honeymon");
  }
}
```
