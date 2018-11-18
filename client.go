package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func client() {
	password := input("Set password: ")
	passwordConfirm := input("Password confirm: ")

	if password != passwordConfirm {
		fmt.Println("Password and password confirm not match!")
		os.Exit(0)
	}

	hashedPassword := hashPassword(password)
	fmt.Printf("Password set success! (%s)\n\n", hashedPassword)

	if pwd := input("Login password: "); checkPassword(pwd, hashedPassword) {
		fmt.Println("Login success!!!")
	} else {
		fmt.Println("Login fail!!!")
	}
}

func input(s string) string {
	fmt.Print(s)

	scanner := bufio.NewReader(os.Stdin)
	result, err := scanner.ReadString('\n')

	if err != nil {
		log.Fatalln(err)
	}

	return strings.Replace(result, "\n", "", -1)
}

func hashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatalln(err)
	}

	return string(hash)
}

func checkPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
