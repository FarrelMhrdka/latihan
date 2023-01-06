package main
import (
    "fmt"
 
)
func main(){
 
    var (
        sisi int
        luas int
    )
    fmt.Printf("Masukkan sisi : ")
    fmt.Scan(&sisi)
    luas = sisi*sisi
    fmt.Printf("Jadi luasnya adalah : %d", luas)
}
