package main

import (
	"image/color"
	"os"

	"fyne.io/fyne/theme"

	"fyne.io/fyne"
)

func init() {
	// 设置字体的方式 1
	os.Setenv("FYNE_FONT", defaultFontName)
}

const defaultFontName = "./default.ttf"

func DefaultTheme() fyne.Theme {
	mt := theme.LightTheme()
	return mt
}

type myTheme struct {
}

func (m myTheme) BackgroundColor() color.Color {
	panic("implement me")
}

func (m myTheme) ButtonColor() color.Color {
	panic("implement me")
}

func (m myTheme) DisabledButtonColor() color.Color {
	panic("implement me")
}

func (m myTheme) HyperlinkColor() color.Color {
	panic("implement me")
}

func (m myTheme) TextColor() color.Color {
	panic("implement me")
}

func (m myTheme) DisabledTextColor() color.Color {
	panic("implement me")
}

func (m myTheme) IconColor() color.Color {
	panic("implement me")
}

func (m myTheme) DisabledIconColor() color.Color {
	panic("implement me")
}

func (m myTheme) PlaceHolderColor() color.Color {
	panic("implement me")
}

func (m myTheme) PrimaryColor() color.Color {
	panic("implement me")
}

func (m myTheme) HoverColor() color.Color {
	panic("implement me")
}

func (m myTheme) FocusColor() color.Color {
	panic("implement me")
}

func (m myTheme) ScrollBarColor() color.Color {
	panic("implement me")
}

func (m myTheme) ShadowColor() color.Color {
	panic("implement me")
}

func (m myTheme) TextSize() int {
	panic("implement me")
}

func (m myTheme) TextFont() fyne.Resource {
	// 设置字体 3
	//res, err = fyne.LoadResourceFromPath(fontpath) // storage.LoadResourceFromURI(url)

	panic("implement me")
}

func (m myTheme) TextBoldFont() fyne.Resource {
	panic("implement me")
}

func (m myTheme) TextItalicFont() fyne.Resource {
	panic("implement me")
}

func (m myTheme) TextBoldItalicFont() fyne.Resource {
	panic("implement me")
}

func (m myTheme) TextMonospaceFont() fyne.Resource {
	panic("implement me")
}

func (m myTheme) Padding() int {
	panic("implement me")
}

func (m myTheme) IconInlineSize() int {
	panic("implement me")
}

func (m myTheme) ScrollBarSize() int {
	panic("implement me")
}

func (m myTheme) ScrollBarSmallSize() int {
	panic("implement me")
}
