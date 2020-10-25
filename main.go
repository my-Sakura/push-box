package main

import (
	"github.com/nsf/termbox-go"
	"fmt"
	"os"
	"os/exec"
)

const (
	W = 7
	H = 8
  )
var (
	GameMap = [W][H]int{
		{0, 0, 0, 2, 0, 0, 0, 0}, 
		{0, 0, 3, 3, 3, 0, 0, 0}, 
		{0, 0, 0, 0, 0, 0, 0, 0}, 
		{1, 1, 0, 0, 0, 1, 1, 1}, 
		{0, 0, 0, 0, 0, 0, 0, 0}, 
		{0, 0, 0, 0, 0, 0, 0, 0}, 
		{0, 0, 0, 0, 0, 4, 4, 4},
	}
)

func initMap() {
	for i := 0; i < W; i++ {
		for j := 0; j < H; j++ {
			switch GameMap[i][j] {
			case 0:
				fmt.Printf(" ")
			case 1:
				fmt.Printf("1")
			case 2:
				fmt.Printf("2")
			case 3:
				fmt.Printf("3")
			case 4:
				fmt.Printf("4")
			case 6:
				fmt.Printf("6")
			case 7:
				fmt.Printf("7")
			}
		}
		fmt.Println()
	}
}
	

	func main() {		
			err := termbox.Init()
			if err != nil {
				panic(err)
			}
			defer termbox.Close()
		
		//clean control 
		
		initMap()
		var flag bool = true
		for flag {
		Control(&flag)
		}
	}


func changeMap() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		panic("run error")
	}
	initMap()
}

func complated() {
	fmt.Println("--恭喜通关--")
}

func Control(flag *bool) {
	var col, raw int
	for i := 0; i < W; i++ {
		for j := 0; j < H; j++ {
			if GameMap[i][j] == 2 || GameMap[i][j] == 6 {
				col = i
				raw = j
			}
		}
	}

	//keyboard input
	label:
	switch ev := termbox.PollEvent(); ev.Type {
	case termbox.EventKey:
		switch ev.Key{
		case termbox.KeyArrowUp:
			if col < 1 {
				goto label
			}

			if GameMap[col-1][raw] == 0 || GameMap[col-1][raw] == 4 {
				GameMap[col][raw] -= 2
				GameMap[col-1][raw] += 2 
				changeMap()	
				if GameMap[6][5] == 7 && GameMap[6][6] == 7 && GameMap[6][7] == 7 {
					complated()							
				}						
			} else if GameMap[col-1][raw] == 3 || GameMap[col-1][raw] == 7 {
				if col == 1 {
					goto label
				}
				if GameMap[col-2][raw] == 0 || GameMap[col-2][raw] == 4 {
					GameMap[col][raw] -= 2
					GameMap[col-1][raw] -= 1
					GameMap[col-2][raw] += 3
					changeMap()	
					if GameMap[6][5] == 7 && GameMap[6][6] == 7 && GameMap[6][7] == 7 {
						complated()							
					}					
				}										
			}
		case termbox.KeyArrowDown:
			if col > 5 {
				goto label
			}
			if GameMap[col+1][raw] == 0 || GameMap[col+1][raw] == 4 {
				GameMap[col][raw] -= 2
				GameMap[col+1][raw] += 2 
				changeMap()
				if GameMap[6][5] == 7 && GameMap[6][6] == 7 && GameMap[6][7] == 7 {
					complated()							
				}	
			} else if GameMap[col+1][raw] == 3 || GameMap[col+1][raw] == 7 {
				if col == 5 {
					goto label
				}
				if GameMap[col+2][raw] == 0 || GameMap[col+2][raw] == 4 {
					GameMap[col][raw] -= 2
					GameMap[col+1][raw] -= 1
					GameMap[col+2][raw] += 3
					changeMap()
					if GameMap[6][5] == 7 && GameMap[6][6] == 7 && GameMap[6][7] == 7 {
						complated()							
					}	
				}										
			}
			
		case termbox.KeyArrowLeft:
			if raw < 1 {
				goto label
			}
			if GameMap[col][raw-1] == 0 || GameMap[col][raw-1] == 4 {
				GameMap[col][raw] -= 2
				GameMap[col][raw-1] += 2 
				changeMap()
				if GameMap[6][5] == 7 && GameMap[6][6] == 7 && GameMap[6][7] == 7 {
					complated()							
				}	
			} else if GameMap[col][raw-1] == 3 || GameMap[col][raw-1] == 7 {
				if raw == 1 {
					goto label
				}
				if GameMap[col][raw-2] == 0 || GameMap[col][raw-2] == 4 {
					GameMap[col][raw] -= 2
					GameMap[col][raw-1] -= 1
					GameMap[col][raw-2] += 3
					changeMap()
					if GameMap[6][5] == 7 && GameMap[6][6] == 7 && GameMap[6][7] == 7 {
						complated()							
					}	
				}										
			}
		case termbox.KeyArrowRight:
			if raw > 6 {
				goto label
			}
			if GameMap[col][raw+1] == 0 || GameMap[col][raw+1] == 4 {
				GameMap[col][raw] -= 2
				GameMap[col][raw+1] += 2 
				changeMap()
				if GameMap[6][5] == 7 && GameMap[6][6] == 7 && GameMap[6][7] == 7 {
					complated()							
				}	
			} else if GameMap[col][raw+1] == 3 || GameMap[col][raw+1] == 7 {
				if raw == 6 {
					goto label
				}
				if GameMap[col][raw+2] == 0 || GameMap[col][raw+2] == 4 {
					GameMap[col][raw] -= 2
					GameMap[col][raw+1] -= 1
					GameMap[col][raw+2] += 3
					changeMap()
					if GameMap[6][5] == 7 && GameMap[6][6] == 7 && GameMap[6][7] == 7 {
						complated()							
					}	
				}										
			}
			//reset map
		case termbox.KeyF5:
			GameMap = [W][H]int{
				{0, 0, 0, 2, 0, 0, 0, 0}, 
				{0, 0, 3, 3, 3, 0, 0, 0}, 
				{0, 0, 0, 0, 0, 0, 0, 0}, 
				{1, 1, 0, 0, 0, 1, 1, 1}, 
				{0, 0, 0, 0, 0, 0, 0, 0}, 
				{0, 0, 0, 0, 0, 0, 0, 0}, 
				{0, 0, 0, 0, 0, 4, 4, 4},
			}
			changeMap()
		case termbox.KeyF1:
			exec.Command("")
			*flag = false
		}
	}
}




	