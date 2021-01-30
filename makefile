
test:
	echo "这个程序不太乖"

build:
	go build -ldflags -H=windowsgui

package:
	fyne package -os windows -icon Icon.png