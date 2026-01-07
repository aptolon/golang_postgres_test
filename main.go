package main

import (
	"fmt"
	httpserver "study/http_server"
)

func main() {
	// ctx := context.Background()
	// conn, err := simple_connection.CreateConnection(ctx)
	// if err != nil {
	// 	panic(err)
	// }
	// if err := simple_sql.CreateTable(ctx, conn); err != nil {
	// 	panic(err)
	// }
	// // if err := simple_sql.InsertRow(
	// // 	ctx,
	// // 	conn,
	// // 	"поесть суп",
	// // 	"мм вкусно",
	// // 	false,
	// // 	time.Now(),
	// // ); err != nil {
	// // 	panic(err)
	// // }

	// tasks, err := simple_sql.SelectRows(ctx, conn)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(tasks)
	// fmt.Println("создано успешно")

	fmt.Println("Запуск сервера")
	if err := httpserver.StartHTTPServer(); err != nil {
		fmt.Println("ошибка ", err)
	} else {
		fmt.Println("сервер завершился успешно ")
	}

	// _, err := os.Create("out/new_file.txt")
	// if err != nil {
	// 	panic(err)
	// }
}
