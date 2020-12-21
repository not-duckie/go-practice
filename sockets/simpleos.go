package main

import "fmt"
import "log"
import "os/exec"

func main(){
	cmd := exec.Command("sh","-c","ls -la")
	stdoutStderr,err := cmd.CombinedOutput()
	if err!= nil{
		log.Fatal(err)
	}
	fmt.Printf("%s\n",stdoutStderr)

}
