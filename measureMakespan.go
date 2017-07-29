package main

import (
	"fmt"
	"os/exec"
	"os"
	"time"
	"bytes"
	"strconv"
	"net"
)

func main() {
	
	i := 0
	ip := getIPAddress()

	for i < 10 {
		if i == 0 || i == 2 || i == 4 || i == 6 || i == 8  {
			var out, stderr bytes.Buffer	
			cmd := exec.Command("docker", "-H", "tcp://" + ip + ":2376", "run", "-v", "/home/smendes/worker:/tmp/workdir", "-w=/tmp/workdir", "-c", "1024", "-m", "2000000000", "jrottenberg/ffmpeg", "-i", "BeachWaves1.flac", "resultadoo.dvd", "-y")

			cmd.Stdout = &out
			cmd.Stderr = &stderr

			start := time.Now()	
			if err := cmd.Run(); err != nil {
				fmt.Println("Not scheduled")
				fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
				time.Sleep(time.Second * 5)
				continue
			} 	
			finish := time.Since(start)

			fileTime, err3 := os.OpenFile("makespan.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
             

        		if err3 != nil{
                		panic(err3)
        		}
        		defer fileTime.Close()
	
			finishTime := strconv.FormatFloat(finish.Seconds(), 'f', -1, 64)
			tempo := time.Now()

        		if _, err3 = fileTime.WriteString("FFmpeg:" + finishTime + ",time:" + tempo.String() + "\n"); err3 != nil {
                		panic(err3)
			}
			i++
			time.Sleep(time.Minute * 5)
		} else  {
			cmd := exec.Command("docker", "-H", "tcp://" + ip + ":2376", "run", "-v", "/home/smendes/worker:/ne/input", "-c", "1024", "-m", "2000000000", "alexjc/neural-enhance", "--zoom=2", "input/buga1.png")
			var out, stderr bytes.Buffer	
	
			cmd.Stdout = &out
			cmd.Stderr = &stderr
	
			start := time.Now()	
			if err := cmd.Run(); err != nil {
				fmt.Println("Not scheduled")
				fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
				time.Sleep(time.Second * 5)
				continue
			} 
			finish := time.Since(start)

			fileTime, err3 := os.OpenFile("makespan.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
             

        		if err3 != nil{
                		panic(err3)
        		}
        		defer fileTime.Close()
	
			finishTime := strconv.FormatFloat(finish.Seconds(), 'f', -1, 64)
			 
			tempo := time.Now()

        		if _, err3 = fileTime.WriteString("enhance:" + finishTime + ",time:" + tempo.String() + "\n"); err3 != nil {
                		panic(err3)
			}
			i++

			time.Sleep(time.Minute * 5)
		}
	}
}

func getIPAddress() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err.Error())
	}
	count := 0
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil && count == 1{
				fmt.Println(ipnet.IP.String())
				return ipnet.IP.String()
			}
			count++
		}
	}
	return ""
}
