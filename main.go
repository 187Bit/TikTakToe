package main

import "fmt"

func main() {
 
  var Feld1, Feld2, Feld3 [3]string
  var Zähler int


  for Zähler < 3 {

    Feld1[Zähler] = "O"
    Feld2[Zähler] = "O"
    Feld3[Zähler] = "O"
    Zähler++

}

//Ausgabe der Matrix
fmt.Println(Feld1)
fmt.Println(Feld2)
fmt.Println(Feld3)


TikTakToe(Feld1,Feld2,Feld3)


}



func TikTakToe(Feld1, Feld2, Feld3 [3]string) {

var Game bool
Game = true
var Eingabe, Eingabe2 int
var Stelle int
var Leer = [3]string {"O","O","O"}




for Game == true {


fmt.Println("Welche Zeile?")
fmt.Scan(& Eingabe)
fmt.Println("Welche Stelle?")
fmt.Scan(& Eingabe2)


//Eingabe == 1

if Eingabe == 1 {


  if Eingabe2 == 1 {

    if Feld1[0] == Leer[0] {

        Feld1[0] = "X"
    }
  }

  if Eingabe2 == 2 {

    if Feld1[1] == Leer[0] {

      Feld1[1] = "X"
    
    }
  }

  if Eingabe2 == 3 {

    if Feld1[2] == Leer[0] {

    Feld1[2] = "X"
    
    }
  }
}


  
//Eingabe == 2

if Eingabe == 2 {

  if Eingabe2 == 1 {

    if Feld2[0] == Leer[0] {

    Feld2[0] = "X"


  }
}

  if Eingabe2 == 2 {

    if Feld2[1] == Leer[0] {

    Feld2[1] = "X"
    
  }
}

  if Eingabe2 == 3 {

    if Feld2[2] == Leer[0] {

    Feld2[2] = "X"
    
    }
  }
}
  

  //Eingabe == 3

if Eingabe == 3 {

    if Eingabe2 == 1 {

      if Feld3[0] == Leer[0] {

    Feld3[0] = "X"


  }
}

  if Eingabe2 == 2 {

    if Feld3[1] == Leer[0] {

    Feld3[1] = "X"
    
  }
}

  if Eingabe2 == 3 {

    if Feld3[2] == Leer[0] {

    Feld3[2] = "X"
    
    }
  }
}
  
  
  fmt.Println(Feld1)  
  fmt.Println(Feld2)
  fmt.Println(Feld3)

//Überprüfung:

var HorizonX = [3]string {"X","X","X"}
var HorizonT = [3]string {"T","T","T"}
//var Leer = [3]string {"O","O","O"}

if Feld1 == HorizonX {
  fmt.Println("Spieler 1 hat gewonnen! ", Feld1)
  Game = false
  continue
}
if Feld2 == HorizonX {
  fmt.Println("Spieler 1 hat gewonnen! ", Feld2)
  Game = false
  continue
}
if Feld3 == HorizonX {
  fmt.Println("Spieler 1 hat gewonnen! ", Feld3)
  Game = false
  continue
}

for j := 0; j < 1; j++ {

if Feld1[Stelle] == HorizonX[Stelle] {

  if Feld2[Stelle] == HorizonX[Stelle] {

    if Feld3[Stelle] == HorizonX[Stelle] {

          fmt.Println("Spieler 1 hat gewonnen! ")
          Game = false
          continue

    }
  }
}else{
  
if Feld1[1] == HorizonX[Stelle] {
       
  if Feld2[1] == HorizonX[Stelle] {
          
      if Feld3[1] == HorizonX[Stelle] {

          fmt.Println("Spieler 1 hat gewonnen! ")
          Game = false
          continue

}
    }
  }else{
    
    if Feld1[2] == HorizonX[Stelle] {

      if Feld2[2] == HorizonX[Stelle] {

        if Feld3[2] == HorizonX[Stelle] {

          fmt.Println("Spieler 1 hat gewonnen! ")
          Game = false
          continue

        }

      }

    }else{

      continue

    }
}
}
}
//Diagonal:

if Feld1[0] == HorizonX[Stelle] {

    if Feld2[1] == HorizonX[Stelle] {

        if Feld3[2] == HorizonX[Stelle] {

          fmt.Println("Spieler 1 hat gewonnen! ")
          Game = false
          continue
        }
    }
}

  if Feld1[2] == HorizonX[Stelle] {

    if Feld2[1] == HorizonX[Stelle] {

      if Feld3[0] == HorizonX[Stelle] {

        fmt.Println("Spieler 1 hat gewonnen! ")
        Game = false
        continue
      }
    }
  }


//Spieler2:

fmt.Println("Spieler2: Welche Zeile?")
fmt.Scan(& Eingabe)
fmt.Println("Spieler2: Welche Stelle?")
fmt.Scan(& Eingabe2)



//fmt.Println(Eingabe2)
//Eingabe == 1

if Eingabe == 1 {


  if Eingabe2 == 1 {

    if Feld1[0] == Leer[0] {

    Feld1[0] = "T"


  }
}

  if Eingabe2 == 2 {

    if Feld1[1] == Leer[0] {

      Feld1[1] = "T"
    
    }
  }

  if Eingabe2 == 3 {

    if Feld1[2] == Leer[0] {

    Feld1[2] = "T"
    
   }
  }
}

  
//Eingabe == 2

if Eingabe == 2 {

  if Eingabe2 == 1 {

    if Feld2[0] == Leer[0] {

    Feld2[0] = "T"

  }
}

  if Eingabe2 == 2 {

    if Feld2[1] == Leer[0] {

    Feld2[1] = "T"
    
  }
}

  if Eingabe2 == 3 {

    if Feld2[2] == Leer[0] {

    Feld2[2] = "T"
    
    }
  }
}
  

  //Eingabe == 3

if Eingabe == 3 {

    if Eingabe2 == 1 {

      if Feld3[0] == Leer[0] {

    Feld3[0] = "T"


  }
}

  if Eingabe2 == 2 {

    if Feld3[1] == Leer[0] {

    Feld3[1] = "T"
    
  }
}

  if Eingabe2 == 3 {

    if Feld3[2] == Leer[0] {

    Feld3[2] = "T"
    
    }
  }
}
  
  
  fmt.Println(Feld1)  
  fmt.Println(Feld2)
  fmt.Println(Feld3)

  //Überprüfung:

  if Feld1 == HorizonT {

  fmt.Println("Spieler 2 hat gewonnen! ",Feld1)  
  Game = false
  continue
  }

  if Feld2 == HorizonT {
  fmt.Println("Spieler 2 hat gewonnen! ", Feld2)
  Game = false
  continue
  }
  if Feld3 == HorizonT {
  fmt.Println("Spieler 2 hat gewonnen! ", Feld3)
  Game = false
  continue
  }

for j := 0; j < 1; j++ {

if Feld1[Stelle] == HorizonT[Stelle] {

  if Feld2[Stelle] == HorizonT[Stelle] {

    if Feld3[Stelle] == HorizonT[Stelle] {

          fmt.Println("Spieler 2 hat gewonnen! ")
          Game = false
          continue

    }
  }
}else{
  
if Feld1[1] == HorizonT[Stelle] {
        
  if Feld2[1] == HorizonT[Stelle] {
       
      if Feld3[1] == HorizonT[Stelle] {

          fmt.Println("Spieler 2 hat gewonnen! ")
          Game = false
          continue

}
    }
  }else{
    
    if Feld1[2] == HorizonT[Stelle] {

      if Feld2[2] == HorizonT[Stelle] {

        if Feld3[2] == HorizonT[Stelle] {

          fmt.Println("Spieler 2 hat gewonnen! ")
          Game = false
          continue

        }

      }

    }else{

      continue

    }
}
}
}

//Diagonal:

if Feld1[0] == HorizonT[Stelle] {

    if Feld2[1] == HorizonT[Stelle] {

        if Feld3[2] == HorizonT[Stelle] {

          fmt.Println("Spieler 2 hat gewonnen! ")
          Game = false
          continue
        }
    }
}

  if Feld1[2] == HorizonT[Stelle] {

    if Feld2[1] == HorizonT[Stelle] {

      if Feld3[0] == HorizonT[Stelle] {

        fmt.Println("Spieler 2 hat gewonnen! ")
        Game = false
        continue
      }
    }
  }


}


}
