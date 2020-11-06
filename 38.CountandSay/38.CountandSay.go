package main

import "fmt"

func main() {
	res:= countAndSay(2)
	fmt.Println(res);
}
func countAndSay(n int) string {
	res :="1";
	for i:=0;i<n-1;i++{
		res = say(res)
	}
	return res
}
func say(v string) string{
	 count := 1
	 prev :=v[0] 
	 var res string;

	 for i:=1;i<len(v);i++{
		if v[i]==prev {
			count += 1
		}else {
			res = fmt.Sprint(res, count,string(prev))
			count = 1
			prev = v[i]
		}
	}
	return fmt.Sprint(res, count,string(prev))
}