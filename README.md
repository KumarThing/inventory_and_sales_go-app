# 🧾 Inventory and Sales Manager (Go)

A simple command-line inventory and sales management system written in Go.
It allows you to manage products, record sales, calculate total revenue, and save or load data using JSON files.

---
# 📦 Features

✅ Add new products with price and stock
✅ Show all products in inventory
✅ Record product sales and update stock automatically
✅ Display all sales records
✅ Calculate total revenue from sales
✅ Save and load data from a .json file
✅ Simple text-based menu for user interaction

---

``` 
Run the Program
go run main.go

---

#  💻 Usage

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

---

## 🧠 Example Workflow
```
> add product
Enter product name: Laptop
Enter product price: 900
Enter product stock quantity: 5

> sell product
Enter product name: Laptop
Enter quantity: 2
✅ Sale recorded! Total = $1800.00
📦 Remaining stock for Laptop: 3

> total revenue
💵 Total Revenue: $1800.00

> save
File Saved.

> exit
Bye bye 👋

