package dic
//Ladon Scanner for golang 
//Author: k8gege
//K8Blog: http://k8gege.org
//Github: https://github.com/k8gege
import (
	"fmt"
	"bufio"
	"strings"
	//"log"
	"os"
)

func UserDic() (users []string) {
	file, err := os.Open("user.txt")
	if err != nil {
		fmt.Println("Open user.txt error, %v", err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		user := strings.TrimSpace(scanner.Text())
		if user != "" {
			users = append(users, user)
		}
	}
	return users
}

func PassDic() (password []string) {
	file, err := os.Open("pass.txt")
	if err != nil {
		fmt.Println("Open pass.txt error, %v", err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		passwd := strings.TrimSpace(scanner.Text())
		if passwd != "" {
			password = append(password, passwd)
		}
	}
	return password
}