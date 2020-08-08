package Expr

import (
	"fmt"
	"testing"
)

//数值的算术运算
func BenchmarkCalc1(b *testing.B) {
	rule2:="$number + 100 - 30 / ( 3 * 2 )"
	m:=map[string]string{"name":"Mary","age":"30","number":"100","item":"a","ngx":"a"}
	fmt.Println(Calc(rule2,m))
}

//关系运算
func BenchmarkCalc2(b *testing.B) {
	rule2:="$name==Mary"
	m:=map[string]string{"name":"Mary","age":"30","number":"12345678","item":"a","ngx":"a"}
	if ok :=Calc(rule2,m).(bool);!ok{
		b.Fatalf("Calculate error")
	}
}

//布尔值的逻辑运算 + 关系运算
func BenchmarkCalc3(b *testing.B) {
	rule3:="$age + 20 && $age < 100"
	m:=map[string]string{"name":"Mary","age":"30","number":"12345678","item":"a","ngx":"a"}
	if ok :=Calc(rule3,m).(bool);!ok{
		b.Fatalf("Calculate error")
	}
}

//列表运算
func BenchmarkCalc4(b *testing.B) {
	rule4:="$item in [1,2,3]"
	m:=map[string]string{"name":"Mary","age":"30","number":"12345678","item":"a","ngx":"a"}
	if ok :=Calc(rule4,m).(bool);!ok{
		b.Fatalf("Calculate error")
	}
}

//字符串的正则匹配
func BenchmarkCalc5(b *testing.B) {
	rule5:="$ngx re `[1-9]+`"
	m:=map[string]string{"name":"Mary","age":"30","number":"12345678","item":"a","ngx":"a"}
	if ok :=Calc(rule5,m).(bool);!ok{
		b.Fatalf("Calculate error")
	}
}

//联合条件
func BenchmarkCalc6(b *testing.B) {
	rule := "$name == Mary && $age > 20 && $age < 100 && $number == 12345678 &&  $item in [a,b,c] &&  $ngx re `[a-z]+` "
	m:=map[string]string{"name":"Mary","age":"30","number":"12345678","item":"a","ngx":"a"}
	if ok :=Calc(rule,m).(bool);!ok{
		b.Fatalf("Calculate error")
	}
}

