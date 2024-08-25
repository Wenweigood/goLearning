package tempconv

import "fmt"

//摄氏度
type Celsius float64

//华氏度
type Fahrenheit float64

//开尔文
type Kelvin float64

const (
	AbsoluteZero Celsius = -273.15
	FreezingC    Celsius = 0
	BoillingC    Celsius = 100
)

//定义在类型Celsius上的方法
func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

//定义在类型Fahrenheit上的方法
func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%gK", k)
}
