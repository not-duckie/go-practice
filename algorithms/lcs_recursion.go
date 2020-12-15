package main

import ("fmt")

func max(a,b int) int {
	if a>b{
		return a
	} else { return b}
}

func lcs(i, j int) int{
	a := "bd"
	b := "abcd"
	if(i==len(a) || j==len(b)){
		return 0;
	} else if (a[i]==b[j]){
		return 1+lcs(i+1,j+1)
	} else { return max(lcs(i+1,j),lcs(i,j+1))}

}

func main(){
	res := lcs(0,0)
	fmt.Println(res)

}
