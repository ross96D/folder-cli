/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// func main() {
// 	cmd.Execute()
// }

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	parent := filepath.Dir(home)
	parent2 := filepath.Dir(parent)
	parent3 := filepath.Dir(parent2)
	parent4 := filepath.Dir(parent3)
	fmt.Println("Home", home, "Parent directory:", parent, " 2", parent2, " 3", parent3, " 4", parent4)
}
