package main

import ("fmt"
"math"
)

func main() {
	res:= divide(-2147483648,-1)
	fmt.Println("res:",res)
}
func divide ( dividend int, divisor int) (x int) {
	
	if dividend==0{
		return;
	} 
	isNegativeDivisor,isNegativeDividend := false,false;
	if dividend<0{
		isNegativeDividend = true;
		dividend =-dividend
	}    
	if divisor<0 {
		isNegativeDivisor = true;
		divisor=-divisor
	}
	temp:=0 
	for {
		if dividend>=temp{
			temp += divisor
			x+=1
		}else {
			x-=1
			if !((isNegativeDividend && isNegativeDivisor) ||(isNegativeDividend==false && isNegativeDivisor==false)) {
				x = -x;
				if x<math.MinInt32{
					x = math.MinInt32
				}
			}else {
				if x>math.MaxInt32{
					x = math.MaxInt32
				}
			}
			 return;
		}
	} 
}

