package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v4"
)

type purchase struct {
	cardNo  string
	expDate string
	amt     int
	curr    string
}

var connP *pgx.Conn

func MainPurchase() {
	// defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	// defer profile.Start(profile.MemProfile, profile.MemProfileRate(1), profile.ProfilePath(".")).Stop()

	p1 := purchase{
		cardNo:  "123",
		expDate: "23/11/1976",
		amt:     1000,
		curr:    "INR",
	}
	fmt.Println("purchase is ", p1)

	start := time.Now()

	var err error
	// The first is the call to postgres the second to a aerospike
	connP, err = pgx.Connect(context.Background(), "postgresql://postgres:xmbd2311@localhost")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(" conn time ", time.Since(start))
	connTime := time.Now()

	if len(os.Args) == 1 {
		printHelp()
		os.Exit(0)
	}

	switch os.Args[1] {
	case "list":
		err = listTask()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to list tasks: %v\n", err)
			os.Exit(1)
		}

	case "add":
		err = addTasks(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to add task: %v\n", err)
			os.Exit(1)
		}

	case "update":
		n, err := strconv.ParseInt(os.Args[2], 10, 32)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable convert task_num into int: %v\n", err)
			os.Exit(1)
		}
		err = updateTasks(int(n), os.Args[3])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to update task: %v\n", err)
			os.Exit(1)
		}

	case "remove":
		n, err := strconv.ParseInt(os.Args[2], 10, 32)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable convert task_num into int: %v\n", err)
			os.Exit(1)
		}
		err = removeTasks(int(n))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to remove task: %v\n", err)
			os.Exit(1)
		}

	default:
		fmt.Fprintln(os.Stderr, "Invalid command")
		// printHelp()
		os.Exit(1)
	}
	fmt.Println(" completion time ", time.Since(connTime))

}

func listTask() error {
	rows, _ := conn.Query(context.Background(), "select * from tasks")

	for rows.Next() {
		var id int
		var description string
		err := rows.Scan(&id, &description)
		if err != nil {
			return err
		}
		fmt.Printf("%d. %s\n", id, description)
	}

	return rows.Err()
}

func addTasks(description string) error {
	_, err := conn.Exec(context.Background(), "insert into tasks(description) values($1)", description)
	return err
}

func updateTasks(itemNum int, description string) error {
	_, err := conn.Exec(context.Background(), "update tasks set description=$1 where id=$2", description, itemNum)
	return err
}

func removeTasks(itemNum int) error {
	_, err := conn.Exec(context.Background(), "delete from tasks where id=$1", itemNum)
	return err
}
