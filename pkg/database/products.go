package database

type Product struct {
    Id int
    Name string
    Price int64
}

var inMemoryProducts = []Product{}


func Save(product Product) {
    
}
