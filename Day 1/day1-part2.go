package main

import ("fmt"; "bufio"; "os"; "log"; "strconv")

func Main_part2() {
    file, err := os.Open("input.txt")
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

        originalPosition := currentPosition

        numberPart, err := strconv.Atoi(line[1:])
        if err != nil {
            log.Fatal(err)
        }

        //count full rounds around the wheel
        timesAtZero += numberPart / 100
        numberPart = numberPart % 100
        
        if line[0] == 'L' {
            currentPosition -= numberPart
            if currentPosition < 0  {
                currentPosition = 100 + currentPosition
                if  originalPosition != 0 && currentPosition != 0{
                    timesAtZero += 1
                }
            } 
        } else {
            currentPosition += numberPart
            if currentPosition > 99 {
                currentPosition = currentPosition - 100
                if  originalPosition != 0 && currentPosition != 0 {
                    timesAtZero += 1
                }
            }
        }
        fmt.Println(currentPosition)

        if currentPosition == 0 {
            timesAtZero += 1
        }

        if currentPosition < 0 || currentPosition > 99 {
            log.Panic("wtff")
        }        

        fmt.Println("timesAtZero")
        fmt.Println(timesAtZero)

    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Println(timesAtZero)

}
