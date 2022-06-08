package webServer

import (
	"fmt"
	"net/http"
	"strconv"
)

type Human struct {
	FName string
	LName string
	Age   int
}

func (hum Human) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	hum.FName = "Deneme"
	hum.LName = "Surname"
	hum.Age = 12

	r.ParseForm()

	fmt.Println(r.Form)
	fmt.Fprint(w, "<table><tr><td><b>Ad</b></td><td><b>Soyad</b></td><td><b>Yaş</b></td></tr><tr><td>"+hum.FName+"</td><td>"+hum.LName+"</td><td>"+strconv.Itoa(hum.Age)+"</td></tr><tr></tr><tr></tr><tr><td><b>Path</b></td><td>"+r.URL.Path+"</td></tr></table>")

	//fmt.Fprint(w, "<table>
	// <tr>
	//    <td><b>Ad</b></td>
	//    <td><b>Soyad</b></td>
	//    <td><b>Yaş</b></td>
	// </tr>
	// <tr>
	//    <td>" + hum.FName + "</td>
	//    <td>" + hum.LName + "</td>
	//    <td>" + hum.Age + "</td>
	// </tr>
	// <tr> </tr>
	// <tr> </tr>
	// <tr>
	//    <td> <b>Path</b> </td>
	//    <td> "+ r.URL.Path +" </td>
	// </tr>
	// </table>")
}
