module tinygo.org/x/bluetooth

go 1.15

//replace github.com/muka/go-bluetooth => github.com/skast96/go-bluetooth v0.0.0-20220323113240-9694dc08e349

require (
	github.com/JuulLabs-OSS/cbgo v0.0.2
	github.com/go-ole/go-ole v1.2.4
	github.com/godbus/dbus/v5 v5.0.3
	github.com/muka/go-bluetooth v0.0.0-20220323170840-382ca1d29f29
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	golang.org/x/sys v0.0.0-20200826173525-f9321e4c35a6 // indirect
	tinygo.org/x/drivers v0.15.1
	tinygo.org/x/tinyterm v0.1.0
)
