package main

import (
	"context"
	"fmt"
	simpleconnection "gitprc/featurePostgres/simpleConnection"
	simplesql "gitprc/featurePostgres/simpleSql"
	"time"
)

func main() {
	ctx := context.Background()
	conn, err := simpleconnection.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}

	if err := simplesql.CreateTable(ctx, conn); err != nil {
		panic(err)
	}

	tasks, err := simplesql.SelectRows(ctx, conn)

	if err != nil {
		panic(err)
	}

	for _, task := range tasks {
		if task.ID == 3 {
			task.Title = "Почистить снег"
			task.Description = "Взять лопату"
			task.Completed = true
			now := time.Now()
			task.CompletedAt = &now

			if err := simplesql.UpdateTask(ctx, conn, task); err != nil {
				panic(err)
			}
			break
		}
	}
	fmt.Println("Успешно")
}
