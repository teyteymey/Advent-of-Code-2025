package main

import ("fmt"; "bufio"; "os"; "log"; "strconv")


func main() {
    file, err := os.Open("input2.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    // dial starts pointing at 50
    currentPosition := 50
    timesAtZero := 0
    for scanner.Scan() {
        fmt.Println("---------")
        line := scanner.Text()
        fmt.Println(line)

        numberPart, err := strconv.Atoi(line[1:])
        if err != nil {
            log.Fatal(err)
        }
        // 5 L10 -5 100-5 = 95
        // 98 R3 101 100-101 = 1
        if line[0] == 'L' {
            currentPosition -= numberPart
        } else {
            currentPosition += numberPart
        }
        fmt.Println(currentPosition)

        //normalize the results in the wheel
        if currentPosition < 0 {
            currentPosition = 100 + currentPosition
        } 
        if currentPosition > 99 {
            currentPosition = currentPosition - 100
        }

        fmt.Println(currentPosition)
        if currentPosition == 0 {
            timesAtZero += 1
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Println(timesAtZero)

}
