package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"image/color"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	go func() {
		w := app.NewWindow(app.Size(1000, 800))

		err := run(w)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

type C = layout.Context
type D = layout.Dimensions

var order string = "Nothing to print"
var uuidInput widget.Editor

func run(w *app.Window) error {
	th := material.NewTheme()
	button := new(widget.Clickable)

	var ops op.Ops
	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			layout.Flex{
				Axis:      layout.Vertical,
				Alignment: layout.Middle,
				Spacing:   layout.SpaceStart,
			}.Layout(gtx,
				layout.Rigid(
					func(gtx C) D {
						title := material.H1(th, "Your order")
						title.Color = color.NRGBA{R: 25, G: 4, B: 130, A: 255}
						title.Alignment = text.Middle
						return title.Layout(gtx)
					},
				),
				layout.Rigid(
					func(gtx C) D {
						y := 0
						if len(order) < 400 {
							y = 500
						}
						margins := layout.Inset{
							Top:    unit.Dp(0),
							Bottom: unit.Dp(y),
							Right:  unit.Dp(200),
							Left:   unit.Dp(200),
						}
						var orderText material.LabelStyle
						if len(order) > 200 {
							orderText = material.Overline(th, order)
						} else {
							orderText = material.Body1(th, order)
						}
						{
							material.Overline(th, order)
						}
						orderText.Color = color.NRGBA{R: 25, G: 4, B: 130, A: 255}

						return margins.Layout(gtx, orderText.Layout)
					},
				),
				layout.Rigid(
					func(gtx C) D {
						margins := layout.Inset{
							Top:    unit.Dp(0),
							Right:  unit.Dp(170),
							Bottom: unit.Dp(40),
							Left:   unit.Dp(170),
						}
						border := widget.Border{
							Color:        color.NRGBA{R: 204, G: 204, B: 204, A: 255},
							CornerRadius: unit.Dp(3),
							Width:        unit.Dp(2),
						}
						return margins.Layout(gtx,
							func(gtx C) D {
								ed := material.Editor(th, &uuidInput, "Order uuid")
								uuidInput.SingleLine = true
								uuidInput.Alignment = text.Middle
								return border.Layout(gtx, ed.Layout)
							})
					},
				),
				layout.Rigid(
					func(gtx C) D {
						gtx.Constraints.Min.X = 300

						b := material.Button(th, button, "Get order")
						return b.Layout(gtx)
					},
				),
				layout.Rigid(
					layout.Spacer{Height: unit.Dp(25)}.Layout,
				),
			)

			if button.Clicked() {
				go func() {
					url := fmt.Sprintf("http://127.0.0.1:8080/api/order/?uid=%s", uuidInput.Text())
					resp, err := http.Get(url)
					if err != nil {
						order = err.Error()
						return
					}
					defer resp.Body.Close()

					data, err := io.ReadAll(resp.Body)
					if err != nil {
						order = err.Error()
						return
					}

					var prettyJSON bytes.Buffer
					err = json.Indent(&prettyJSON, data, "", "    ")
					if err != nil {
						order = string(data)
						return
					}
					order = prettyJSON.String()
				}()
			}

			e.Frame(gtx.Ops)
		}
	}
}
