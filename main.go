package main
import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"encoding/json"

)
type Inventory struct {
	Name string
	Price float64
	Stock int
}

type Sale struct {
	ProductName string
	Quantity    int
	TotalPrice  float64
}



func main() {
	reader := bufio.NewReader(os.Stdin)
	var inventory[] Inventory
	var sale[] Sale
	

	fmt.Println("Welcome to Inventory and Sales Manager!")
	fmt.Println(`
-------------------------------------------------
1. add product
2. show products
3. sell product
4. show sales
5. total revenue
6. low stock alert
7. save
8. load
9. exit
--------------------------------------------------
	`)

fmt.Print("Enter Command: ")

for  {
	input,_:= reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "add product" {
		fmt.Print("Enter product name: ")
		name,_ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		exists := false

		for i := range inventory {
			if strings.EqualFold(name,inventory[i].Name) {
				exists = true
			}
		}
		if exists {
			fmt.Println("Name alreadi exists")
			continue
		}

		fmt.Print("Enter product price: ")
		priceStr,_ := reader.ReadString('\n')
		priceStr = strings.TrimSpace(priceStr)

		price,err := strconv.ParseFloat(priceStr,64)
		if err != nil {
			fmt.Println("Error.Please enter correct price")
		}

		fmt.Print("Enter product stock quantity: ")
		stockStr,_:= reader.ReadString('\n')
		stockStr = strings.TrimSpace(stockStr)

		stock,err := strconv.Atoi(stockStr)
		if err != nil {
			fmt.Println("Error. Please correct stock number.")
		}

		inventory = append(inventory, Inventory{name,price,stock})
		
		fmt.Printf("✅ Product added: %s - $%.2f (stock: %d)\n",name, price, stock)


	} else if input == "show products" {
		if len(inventory) == 0 {
			fmt.Println("Empty inventory data. ")
		} else {
			for i, e := range inventory {
				fmt.Printf("%d. %s - $%.2f (stock:%d)\n",i+1, e.Name, e.Price, e.Stock)
			}
		}
	} else if input == "sell product" {
		if len(inventory) == 0 {
			fmt.Println("Inventory data empty. No products.")
		} else {
			fmt.Print("Enter product name: ")
			name,_:= reader.ReadString('\n')
			name = strings.TrimSpace(name)

			found := false
			for i := range inventory {
				if strings.EqualFold(name,inventory[i].Name){
					found = true

					fmt.Print("Enter quantity: ")
					quantityStr,_:= reader.ReadString('\n')
					quantityStr = strings.TrimSpace(quantityStr)

					quantity, err := strconv.Atoi(quantityStr)
					if err != nil {
					fmt.Println("Error. Please enter correct quantity number.")
					}

					if quantity > inventory[i].Stock {
						fmt.Println("Not enough stock available.")
						return
					}

					total := float64(quantity) * inventory[i].Price

					inventory[i].Stock -= quantity

					sale = append(sale, Sale{
						ProductName: inventory[i].Name,
						Quantity: quantity,
						TotalPrice: total, 
					})

					fmt.Printf("✅ Sale recorded! Total = $%.2f\n", total)
				fmt.Printf("📦 Remaining stock for %s: %d\n", inventory[i].Name, inventory[i].Stock)
				break
				}

				if !found{
					fmt.Println("Product not found in inventory.")}
				}
			}
			
		} else if input == "show sales" {
			 if len(sale) == 0 {
				fmt.Println("No Sales record.")
			 } else {
				for i,e := range sale {
					fmt.Printf("%d. %s - %d sold (%.2f)",i+1, e.ProductName, e.Quantity, e.TotalPrice)
				}
			 }
		} else if input == "total revenue" {
			if len(sale) == 0 {
				fmt.Println("No sales recorded.")
			} else {
				for i := range sale {
					fmt.Printf("Total Revenue: $%.2f", sale[i].TotalPrice)
				}
			}
		} else if input == "save" {
			if len(inventory) == 0 {
				fmt.Println("No data to save. Empty:")
			} else {
				file,err := os.Create("inventory.json")
				if err != nil {
					fmt.Println("File not saved",err)
				}

				defer file.Close()
				
				data, err := json.MarshalIndent(inventory,"", " ")
				if err != nil {
					fmt.Println("Error converting .txt to json")
				}

				file.Write(data)

				fmt.Println("Flie Saved.")

			}
		} else if input == "load" {
			data, err := os.ReadFile("inventory.json")
			if err != nil {
				fmt.Println("Error",err)
			}
			err = json.Unmarshal(data, &inventory)
			if err != nil {
				fmt.Println("Error",err)
			}
			fmt.Println("Data load succefully.")
			
		} else if input == "exit" {
			fmt.Println("\nBye bye \n")
			break

	} else {
		fmt.Println("Error. Please enter correct command. ")
	}
	fmt.Println(`
-------------------------------------------------
1. add product
2. show products
3. sell product
4. show sales
5. total revenue
6. low stock alert
7. save
8. load
9. exit
--------------------------------------------------
	`)

fmt.Print("Enter Command: ")
	} 

	
}


	
