package Parse

import (
	"fmt"
	"io"
	http "net/http"
	// os "os"
)

func ApiCall(url string) (string,error){
	fmt.Println("received ",url)
	resp,err := http.Get(url);
	if(err != nil){
		fmt.Println("Error: ",err)
		return "",err;
	}
	fmt.Println("after get")
	defer resp.Body.Close()

	body,err:=io.ReadAll(resp.Body)
	if(err != nil){
		fmt.Println("Error: ",err)
		return "",err;
	}
	return string(body),nil

}