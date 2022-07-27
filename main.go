package main

import (
    "fmt"

    g "github.com/AllenDang/giu"
)

var gamePath string = `C:\Program Files (x86)\Steam\steamapps\common\Legend Bowl\legend-bowl.exe`

func showGamePath() {
    fmt.Println(gamePath)
}

func loop() {
    g.SingleWindow().Layout(
        g.Row(
            g.Label("Game Path:"),
            g.InputText(&gamePath).Size(g.Auto),
        ),
        g.Row(
            g.Label(""),
        ),
        g.Row(
            g.Label("Detected Rosters:"),
            g.Button("Refresh List"),
        ),
        g.Row(
            g.RadioButton("Vanilla", true),
        ),
        g.Row(
            g.RadioButton("NFL_v4", false),
        ),
        g.Row(
            g.RadioButton("USFL_v1", false),
        ),
        g.Row(
            g.Label(""),
        ),
        g.Row(
            g.Button("Launch Legend Bowl").OnClick(showGamePath),
        ),
    )
}

func main() {
    window := g.NewMasterWindow("Legend Bowl Launcher", 800, 400, g.MasterWindowFlagsNotResizable)
    window.Run(loop)
}
