package main

import (
	"fmt"
	"os/exec"
	"os"
	"time"
	"bytes"
	"strconv"
)

func main() {
	
	i := 0

	for i < 10 {
		if i == 0 || i == 2 || i == 4 || i == 6 || i == 8  {
			i++
			var out, stderr bytes.Buffer	
			cmd := exec.Command("docker", "-H", "tcp://10.5.60.1:2376", "run", "-v", "/home/smendes:/tmp/workdir", "-w=/tmp/workdir", "-c", "1024", "-m", "2000000000", "jrottenberg/ffmpeg", "-i", "BeachWaves1.flac", "resultadoo.dvd", "-y")

			cmd.Stdout = &out
			cmd.Stderr = &stderr

			start := time.Now()	
			if err := cmd.Run(); err != nil {
				fmt.Println("Not scheduled")
				fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
			} 	
			finish := time.Since(start)

			fileTime, err3 := os.OpenFile("makespan.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
             

        		if err3 != nil{
                		panic(err3)
        		}
        		defer fileTime.Close()
	
			finishTime := strconv.FormatFloat(finish.Seconds(), 'f', -1, 64)
			
			fmt.Println("Elapsed time: " + finishTime)

        		if _, err3 = fileTime.WriteString("FFmpeg time: " + finishTime + "\n"); err3 != nil {
                		panic(err3)
			}
			time.Sleep(time.Minute * 5)
		} else  {
			i++
			cmd := exec.Command("docker", "-H", "tcp://10.5.60.1:2376", "run", "-v", "/home/smendes:/ne/input", "-c", "1024", "-m", "2000000000", "alexjc/neural-enhance", "--zoom=2", "input/buga1.png")
			var out, stderr bytes.Buffer	
	
			cmd.Stdout = &out
			cmd.Stderr = &stderr
	
			start := time.Now()	
			if err := cmd.Run(); err != nil {
				fmt.Println("Not scheduled")
				fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
			} 
			finish := time.Since(start)

			fileTime, err3 := os.OpenFile("makespan.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
             

        		if err3 != nil{
                		panic(err3)
        		}
        		defer fileTime.Close()
	
			finishTime := strconv.FormatFloat(finish.Seconds(), 'f', -1, 64)

        		if _, err3 = fileTime.WriteString("enhance time: " + finishTime + "\n"); err3 != nil {
                		panic(err3)
			}
			fmt.Println("Elapsed time: " + finishTime)

			time.Sleep(time.Minute * 5)
		}
	}
}
