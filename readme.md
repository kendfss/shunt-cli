Rube
----
A game inspired by the notion of a 2D-Rubik's Cube  

The objective is to use flips and slides to sort the tiles in the grid.


### Start
```shell
> cd rube/cli
> go mod tidy
> go run game.go
{9}| 7 | 8 
 6 | 4 | 1 
 2 | 3 | 5
```
The objective is to sort the matrix into using as few moves as possible:  
```shell
 1 | 2 | 3 
 4 | 5 | 6 
 7 | 8 | 9 
```
 
### Keybindings
```json
[
    {
        "move_up": "w",
        "move_down": "s",
        "move_left": "a",
        "move_right": "d",
        
        "action_up": "arrow_up",
        "action_down": "arrow_down",
        "action_left": "arrow_left",
        "action_right": "arrow_right",
        "toggle_slide_flip": "space",
        
        "undo": "ctrl+z",
        "redo": "ctrl+y",
        "reset": "ctrl+r",
        
        "new_game": "n",
        "load": "l"
        "hide": "h"
        "info": "i"
        "save": "ctrl+s",
        "next_panel": "tab",
        "quit": "ctrl+c",
    }
]
```
