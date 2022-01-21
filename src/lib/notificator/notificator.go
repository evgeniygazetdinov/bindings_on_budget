package notif

import (
	"os/exec"
    "log"
)


func NOTIFY_ME(my_string ... string){
    var cmd *exec.Cmd
    if len(my_string) > 0 {
        prompt := my_string[0]
        cmd = exec.Command(`notify-send  -t 0 "` + prompt + `"`)
    }else{
        cmd = exec.Command(`notify-send  -t 0 "Bringing down the system"`)
    }

    err := cmd.Run()

    if err != nil {
        log.Fatal(err)
    }
}