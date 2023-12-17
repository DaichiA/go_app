/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"net/http"
	"log"
	"io/ioutil"

	"database/sql"
	_ "github.com/lib/pq"

	"encoding/json"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("add called")

		entry_id     := args[0]
		access_token := args[1]
		url          := "https://cdn.contentful.com/spaces/2vskphwbz4oc/entries/" + entry_id + "?access_token=" + access_token
		res,err:=http.Get(url)
		if err!=nil {
			log.Fatal(err)
		}

		// deferは関数の終了時に実行される
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println(string(body));

		type breadData struct{
			Sys struct {
				Id string `json:"id"`
				CreatedAt string `json:"createdAt"`
			} `json:"sys"`
			Fields struct {
				Name string `json:"name"`
			} `json:"fields"`
		}

		bread := new(breadData)

		// bodyをjsonデコードして、breadに格納
		json.Unmarshal([]byte(body), bread);
		// fmt.Println(bread)

		// fmt.Println(bread.Sys.Id)
		// fmt.Println(bread.Sys.CreatedAt)
		// fmt.Println(bread.Fields.Name)


		// var stcData struct{}
		// json.Unmarshal([]byte(body), &stcData);
		// fmt.Println(stcData);

		// DB保存
		db(bread.Sys.Id, bread.Fields.Name, bread.Sys.CreatedAt)

		// result, err := db(bread.Sys.Id, bread.Fields.Name, bread.Sys.CreatedAt)
		// for result.Next() {
		// 	var id string
		// 	var name string
		// 	var created_at string
		// 	err := result.Scan(&id, &name, &created_at)
		// 	if err != nil {
		// 		fmt.Println(err)
		// 	}
		// 	fmt.Println(id, name, created_at)
		// }
	},
}

// func db(id string, name string, created_at string) (*sql.Rows, error) {
func db(id string, name string, created_at string) {
	// DB接続
	// ローカルで動作確認するだけなので全てベタ書き
	db, err := sql.Open("postgres", "host=db user=postgres dbname=go_app password=mypassword sslmode=disable")
    defer db.Close()

    if err != nil {
        fmt.Println(err)
    }

	fmt.Println("db接続成功")

	_, err = db.Exec("INSERT INTO breads(id, name, created_at) VALUES($1, $2, $3);", id, name, created_at)
	if err != nil {
        fmt.Println(err)
    } else {
		fmt.Println("insert成功")
	}

	// var rows *sql.Rows
	// rows, err = db.Query("SELECT * FROM breads;")
	// if err != nil {
    //     fmt.Println(err)
    // }
	// return rows, nil
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
