package main

import (
  "fmt"
  "strings"
  "math"
  "bufio"
  "os"
)

func main() {
  //gathering user input
  encrypt := true
  fmt.Println("Masukkan kata kunci")
  scan := bufio.NewScanner(os.Stdin)
  scan.Scan()
  keyword := scan.Text()
  data:=initializeDataSet(keyword)
  fmt.Println("Masukkan pesan yang ingin Anda enkripsi/dekripsi (tanpa spasi)")
  scan.Scan()
  plainText := scan.Text()
  fmt.Println("Masukkan E untuk enkripsi atau D untuk mendekripsi")
  scan.Scan()
  hold := scan.Text()

  //converting all to lowercase
  hold=strings.ToLower(hold)
  keyword=strings.ToLower(keyword)
  plainText=strings.ToLower(plainText)

  //runs either encryption or decryption and displays the data
  if strings.Compare(hold,"e") ==0 {
    encrypt=true
    encryptedT:=(splitLetters(data,plainText,encrypt))
    fmt.Println("Data Terenkripsi: " , encryptedT)

    //gives the option to decrypt the data (testing purposes only)
    fmt.Println("Apakah Anda ingin mendekripsi data sekarang? iya/tidak")
    scan.Scan()
    if strings.Compare(scan.Text(),"iya") ==0{
      decryT:=(splitLetters(data,encryptedT,false))
      fmt.Println("Data yang Didekripsi: " , decryT)
    }
  } else {
    encrypt=false
    encryptedT:=(splitLetters(data,plainText,encrypt))
    fmt.Println("Data yang Didekripsi: " , encryptedT)
  }
}

func initializeDataSet(keyword string) [5][5]string {
  pos := 1
  alpos :=1
  //alph - q
  alph := "abcdefghijklmnoprstuvwxyz"
  letters := [5][5]string{}

  //makes sure there are no duplicates of letters in the keyword
  for q :=1; q <=len(keyword)-1;q++ {
      substring := keyword[q:len(keyword)]
      holder:= keyword[0:q+1]
      substring = strings.Replace(substring ,(string([]rune(keyword)[q])) ,"" , -1)
      keyword = holder + substring
  }

  //makes sure there are no dubpliate letters in the alph
  for c :=0;c<=len(keyword)-1;c++ {
    if strings.Contains(alph , (string([]rune(keyword)[c]))){
      alph = strings.Replace(alph , (string([]rune(keyword)[c])) , "" , -1)
    }
  }

  //loads the 5x5 array
  for x :=0; x<=4; x++ {
    for z :=0;z<=4;z++ {
      if len(keyword)+1 <= pos {
          letters[x][z] = (string([]rune(alph)[alpos-1]))
        alpos++
      } else {
          letters[x][z] = (string([]rune(keyword)[pos-1]))
        pos++
      }
    }
  }
  return letters
}

func encrypt(data [5][5]string, letters string) string {
  //variables of the positions of the letters in the grid
  l1x:=0
  l2x:=0
  l1y:=0
  l2y:=0
  found := false
  l1 := (string([]rune(letters)[0]))
  l2 := (string([]rune(letters)[1]))

  //searches for the letters in the grid, breaks the loop if found
  for l1x <=4 {
    l1y=0
    for l1y<=4 {
      if l1 == data[l1x][l1y] {
        found = true
        break
      }
      l1y++
    }
    if found == true {
      found = false
      break
    }
    l1x++
  }

  for l2x <=4 {
    l2y =0
    for l2y <=4 {
      if l2 == data[l2x][l2y] {
        found = true 
        break
      }
      l2y++
    }
    if (found == true) {
      break 
    }
    l2x++
  }

  //case 1
  if l2x == l1x {
    l2x++
    l1x++
    if l2x > 4 {
      l2x =0
    }
    if l1x > 4 {
      l1x =0
    }
  }

  //case 2
  if l2y == l1y {
    l2y++
    l1y++
    if l1y > 4 {
      l1y =0
    }
    if l2y > 4 {
      l2y =0
    }
  }

  //case 3
  if l2y != l1y && l2x != l1x {
    holder := l1x
    l1x = l2x
    l2x = holder
  }
  returnVal := data[l1x][l1y]+data[l2x][l2y]
  return (returnVal)
}

func unEncrypt(data [5][5]string, letters string) string{
  //same as the encrypt function
  l1x:=0
  l2x:=0
  l1y:=0
  l2y:=0
  found := false
  l1 := (string([]rune(letters)[0]))
  l2 := (string([]rune(letters)[1]))
  for l1x <=4 {
    l1y=0
    for l1y<=4 {
      if l1 == data[l1x][l1y]{
        found = true
        break
      }
      l1y++
    }
    if found == true {
      found = false
      break
    }
    l1x++
  }

  for l2x <=4 {
    l2y =0
    for l2y <=4 {
      if l2 == data[l2x][l2y] {
        found = true 
        break
      }
      l2y++
    }
    if (found == true) {
      break 
    }
    l2x++
  }

  //case 1
  if l2x == l1x {
    if l2x >0 {
      l2x--
    } else {
      l2x=4
    }
    if l1x > 0 {
      l1x--
    } else{ 
      l1x=4
    }
  }

  //case 2
  if l2y == l1y {
    if l2y >0 {
      l2y--
    } else {
      l2y=4
    }
    if l1y > 0 {
      l1y--
    } else {
      l1y=4
    }
  }

  //case 3
  if l2y != l1y && l2x != l1x{
    holder := l1x
    l1x = l2x
    l2x = holder
  }
  returnVal:=data[l1x][l1y]+data[l2x][l2y]
  return (returnVal)
}

func splitLetters(data [5][5]string , keyword string , encryption bool) string{
  returnVal := ""
  z:=0 

  //converting the length from int to float because mod only works if its float64
  leng:=float64(len(keyword))
  if math.Mod(leng,2)!=0 {
    keyword+="x"
  }

  //splits up the letters into pairs and encrypts/decrypts the data then returns the string
  for z <=len(keyword)-1{
    if (encryption ==true) {
      returnVal+=encrypt(data , keyword[z:z+2])
      z+=2
    } else {
      returnVal+=unEncrypt(data , keyword[z:z+2])
      z+=2
    }
  }
  return returnVal
}