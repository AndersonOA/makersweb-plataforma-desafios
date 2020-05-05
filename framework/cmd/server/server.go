package main

import (
	"fmt"
	"github.com/AndersonOA/makersweb-plataforma-desafios/application/repositories"
	"github.com/AndersonOA/makersweb-plataforma-desafios/domain"
	"github.com/AndersonOA/makersweb-plataforma-desafios/framework/utils"
	"log"
)

func main() {

	db := utils.ConnectDB()

	user := domain.User{
		Name:     "Anderson O. Aristides",
		Email:    "aoaristides@gmail.com",
		Password: "123456",
	}

	userRepo := repositories.UserRepositoryDb{Db: db}
	result, err := userRepo.Insert(&user)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result.Name, result.Email, result.Token)

}
