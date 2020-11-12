# 规则引擎

## 技术背景

golang 、antrl 4

## 实现功能

- ⽀持数值类型，字符串类型，布尔类型的表达

- ⽀持ipv4，cidr类型的表达

- ⽀持容器类型的表达

- 支持数值运算

- 支持布尔运算

- 支持关系运算

- 支持列表运算

- ⽀持字符串的正则匹配

## 测试用例

```
//简单运算
func TestCalc1(t *testing.T) {
	rule1:="$name == Mary && ( $age < 30 && $age > 100 || ($age - 30) < 0 ) "
	m := map[string]string{"age": "30", "name": "Mary"}
	if ok :=Calc(rule1,m).(bool);ok{
		t.Fatalf("Calculate error")
	}
}

//集合运算 + 布尔运算
func TestCalc2(t *testing.T) {
	rule:="$name == Mary && ( $age > 20 && $age < 100 ) || $number == 1234567890 && $item in [aa,bb,cc] || $flag == false"
	m := map[string]string{"age": "30", "name": "Mary", "number": "1234567890", "item": "dd", "flag": "true"}
	if ok :=Calc(rule,m).(bool);ok{
		t.Fatalf("Calculate error")
	}
}

//+正则运算
func TestCalc3(t *testing.T) {
	rule:="$name == Mary && ( $age > 20 && $age < 100 ) || $number == 1234567890 && $item in [aa,bb,cc] || ( abcd like `ab.*`)"
	m := map[string]string{"age": "30", "name": "Mary", "number": "1234567890", "item": "dd"}
	if ok :=Calc(rule,m).(bool);!ok{
		t.Fatalf("Calculate error")
	}
}

//+IP 段 验证
func TestCalc4(t *testing.T) {
	rule := "$name == Mary && ( $age > 20 && $age < 100 ) || $number == 1234567890 && $item in [aa,bb,cc] || ( abcd like `ab.*`) && $ip at 192.168.23.11-192.168.23.16"
	m := map[string]string{"age": "30", "name": "Mary", "number": "1234567890", "item": "dd", "ip": "192.168.23.14"}
	if ok :=Calc(rule,m).(bool);!ok{
		t.Fatalf("Calculate error")
	}
}

//cidr ip段验证
func TestCalc5(t *testing.T) {
	rule := "$name == Mary && ( $age > 20 && $age < 100 ) || $number == 1234567890 && $item in [aa,bb,cc] || ( abcd like `ab.*`) && $ip1 at 192.168.23.11-192.168.23.16 && $ip2 belong 192.168.0.0/23"
	m := map[string]string{"age": "30", "name": "Mary", "number": "1234567890", "item": "dd", "ip1": "192.168.23.14", "ip2": "192.168.1.2"}
	if ok :=Calc(rule,m).(bool);!ok{
		t.Fatalf("Calculate error")
	}
}
```

## 基准测试结果

```
goos: windows
goarch: amd64
pkg: engine/CMD
BenchmarkCalcCalc1-8    1000000000               0.000997 ns/op        0 B/op          0 allocs/op
BenchmarkCalcCalc2-8    1000000000               0.000999 ns/op        0 B/op          0 allocs/op
BenchmarkCalcCalc3-8    1000000000               0.000991 ns/op        0 B/op          0 allocs/op
BenchmarkCalcCalc4-8    1000000000               0.000998 ns/op        0 B/op          0 allocs/op
BenchmarkCalcCalc5-8    1000000000               0.00104 ns/op         0 B/op          0 allocs/op
PASS
ok      engine/CMD      0.473s

```


## 详细设计

**Expr 文件夹**

------

```
├── Expr.g4
├── Expr.interp
├── Expr.tokens
├── expr_base_listener.go
├── expr_base_visitor.go
├── expr_lexer.go
├── expr_listener.go``
├── expr_parser.go
├── expr_visitor.go
├── ExprLexer.tokens
└── ExprLexer.interp
```

> 表达式解析通过 antrl4 工具生成

**具体实现方式**

```
1. 先在expr.g4文件中定义词法与语法
2. 通过命令生成 .go 文件接口 及其他相关文件
3.语法树将token进行解析
4. 在expr.go 实现 basevisitor 接口文件中定义具体运算规则及之后的操作 
5. 通过字符串替换的方式实现变量值进行注入
```

