package main

import (
	"computer_shop/dao"
	"fmt"
)

func main() {
	var CategoryDAO dao.CategoryDAO
	CategoryDAO.Delete(1)
	fmt.Println(CategoryDAO.FindAll())
}
