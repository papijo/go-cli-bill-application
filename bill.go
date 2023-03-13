package main

import (
	"fmt"
	"os"
)

//Custom Type/Structs
type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// Make new bills
func newBill(name string) bill {
	b := bill{
		name:  name + "'s bill" ,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

//method(receiver function) : format the bill
func (b bill) format() string {
fs:= "Bill Breakdown: \n"
var total float64 = 0

//list items
for k, v := range b.items {
	fs += fmt.Sprintf("%-25v....$%v \n", k+":", v)
	total += v
}

//add tip
fs += fmt.Sprintf("%-25v....$%v\n", "tip:", b.tip)

//total
fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total + b.tip)

return fs
}

//update tip
func (b *bill) updateTip(tip float64){
	b.tip = tip
}

//add an item to the bill
func (b *bill) addItem (name string, price float64){
	b.items[name] =  price
}

//save bill

func (b *bill) save() {
	data := []byte(b.format())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)

	if err!= nil{
		panic(err)
	}

	fmt.Println("Bill was saved to file")
}