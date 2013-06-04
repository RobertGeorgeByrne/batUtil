package main

import (
	"fmt"
    "strings"
	"os/exec"
    "regexp"
)

func main() {
    r, err := exec.Command("acpi", "-V").Output()
    var batteryInfo string;
    var outputMsg string = ""
    if err != nil {
        fmt.Println("error: ", err)
        return
    }

    batteryInfo = string(r)
    
    percentageRegex, err := regexp.Compile(`\d+%`)

    if err != nil{
        fmt.Println("error: ", err)
        return
    }

    outputMsg += percentageRegex.FindString(batteryInfo) + " "

    if strings.Contains(batteryInfo, "on-line"){
        outputMsg = outputMsg + "" 
    }

    fmt.Println(outputMsg)
    fmt.Println("\n", batteryInfo)
    return 
}
