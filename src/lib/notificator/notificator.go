package notif

import (
	"os/exec"
    "log"
    "bytes"
)

const ShellToUse = "bash"

func shellout(command string) (error, string, string) {
    var stdout bytes.Buffer
    var stderr bytes.Buffer
    cmd := exec.Command(ShellToUse, "-c", command)
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    err := cmd.Run()
    return err, stdout.String(), stderr.String()
}


func NOTIFY_ME(my_string ... string){
    var cmd *exec.Cmd
    if len(my_string) > 0 {
        prompt := my_string[0]
        shellout(prompt)
    }else{
        shellout(`notify-send  -t 0 "Bringing down the system"`)
    }

    err := cmd.Run()

    if err != nil {
        log.Fatal(err)
    }
}