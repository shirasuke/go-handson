package output

import (
	"fmt"
	"os"
	"strings"
)




func Output(files []os.DirEntry){
	for _,file := range files{
		if strings.HasPrefix(file.Name(), ".") {
            continue
        }
		fmt.Println(file.Name())
	}
}

func OutputAll(files []os.DirEntry){
	for _,file := range files{
		fmt.Println(file.Name())
	}
}