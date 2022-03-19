package notif

import (
	"fmt"
	"os/exec"
)

func NOTIFY_ME(my_string ...string) {
	var res *exec.Cmd
	if len(my_string) > 0 {
		prompt := my_string[0]
		res = exec.Command("notify-send", prompt)
	} else {
		res = exec.Command("notify-send", "Время вышло")
	}
	stdin, _ := res.CombinedOutput()
	fmt.Println(string(stdin))
}
