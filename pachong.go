package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"io"
)

func checkFileIsExist(filename string) (bool) {
	var exist = true;
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false;
	}
	return exist;
}

func main() {
	
	for _, url := range os.Args[1:]{
		resp, err := http.Get(url)
		if err != nil{
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		content := string(b)
//		fmt.Printf("%s",content)
		
		dir := "files"
		_, err = os.Stat(dir)
		
		if err != nil || os.IsExist(err){
 			err = os.Mkdir(dir, os.ModePerm)
			if err != nil {
				fmt.Println(err)
				return
			} 
			fmt.Printf("Mkdir(%s)",dir)
		}else{
			fmt.Printf("dir(%s) is exist! ",dir)
		}
		name := "filename"
		file := dir + "/" + name
		fmt.Println(file)
		
		var fs    *os.File
		_, err = os.Stat(file)
		 if !checkFileIsExist(file) { 
			fs, err = os.Create(file);
			if err != nil {
				fmt.Fprintf(os.Stderr, "Create: %v\n", err)
			}
			fmt.Println("Create file", file);
		}else {
			fs, err = os.OpenFile(file, os.O_APPEND, 0666)
			if err != nil {
				fmt.Fprintf(os.Stderr, "OpenFile: %v\n", err)
			}
			fmt.Println("file exist");
		}
		defer fs.Close()
		n, err := io.WriteString(fs, content) //
		if err != nil {
			fmt.Println("WriteString err");
		}
		fmt.Printf("write %d Byte", n);
		
	}
	
}