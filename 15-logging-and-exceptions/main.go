package main

import (
	"errors"
	"fmt"
	"log"
	"log/syslog"
	"os"
)

/*******
OOTB logging
*/

/*
	For more advanced logging, try out Logrus

	$ go get github.com/sirupsen/logrus

	Levels:
		log.Trace("going to mars")
		log.Debug("connected, received buffer from worker")
		log.Info("connection successful to db")
		log.Warn("something went wrong with x")
		log.Error("an error occurred in worker x")
		log.Fatal("impossible to continue exec")
		log.Panic("system emergency shutdown")
*/

func main() {

	playErrors()

	playLogging()

}

func playErrors() {
	// Golang does not have Exceptions like Java
	// To understand errors, use the mulitple return value from a function.
	fmt.Println("\n***Play Errors")
	x, err := divide(3, 4)
	if err != nil {
		fmt.Println("\nDivision Error: " + err.Error())

	} else {
		fmt.Printf("\n3/4 = %f\n", x)
	}

	y, err := divide(5, 0)
	if err != nil {
		fmt.Println("\nDivision Error: " + err.Error())

	} else {
		fmt.Printf("\n5/0 = %f\n", y)
	}

}

func divide(num, den float64) (float64, error) {
	// cannot divide by zero
	if den == 0 {
		return 0, errors.New(fmt.Sprintf("The denominator cannot be zero! [%f / %f ]", num, den))
	}
	return num / den, nil
}

func playLogging() {
	LOG_FILE := "log/app.log"

	fmt.Println("=======logging and exceptions======")
	setup(LOG_FILE)

	f, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	// set the output of the log to the newly opened file
	log.SetOutput(f)

	log.Println("Starting the play around with logging")
	log.Print("System checkpoint 1")
	log.Printf("System checkpoint %d", 2)

	// log.Fatal & log.Panic will terminate the program
	// log.Fatalf("System fatal checkpoint %d", -1000)
	// log.Panicf("System fatal checkpoint %d", -2000)

	// Syslog is an open standard for message logging.
	fmt.Println("***Syslogs")

	logwriter, err := syslog.New(syslog.LOG_WARNING, "loggingTestProgram")
	if err != nil {
		log.Fatal(err)
	}
	logwriter.Emerg("This is a log entry from GoPlayLogs... 1")

	// hostname := os.Hostname()

	// $ tail -f /var/log/syslog

	// Log to syslog
	logWriter2, err := syslog.New(syslog.LOG_SYSLOG, "Play Go App")
	if err != nil {
		log.Fatalln("Unable to set logfile:", err.Error())
	}
	// set the log output
	log.SetOutput(logWriter2)
	log.Println("Logging to syslog")

	// custom log file
	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)

	// optional: log date-time, filename, and line number
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	log.Println("Logging to custom file")
	log.Println("This is a log entry from GoPlayLogs... 2")
}

func setup(logFile string) {
	// prepare to run by deleting the log file (if it exists)
	fmt.Println("\nSetup...")

	if _, err := os.Stat(logFile); err == nil {
		e := os.Remove("log/app.log")
		if e != nil {
			fmt.Println("ERROR")
			log.Fatal(e)
		}
	}
}
