package banner

func Show() {
	bar := "_________________________________________________________________\n"
	print(bar)
	println(`
   ██     ██  █████  ███████ ████████ ██████   █████  ██████  
   ██     ██ ██   ██ ██         ██    ██   ██ ██   ██ ██   ██ 
   ██  █  ██ ███████ ███████    ██    ██████  ███████ ██████  
   ██ ███ ██ ██   ██      ██    ██    ██   ██ ██   ██ ██      
    ███ ███  ██   ██ ███████    ██    ██   ██ ██   ██ ██

        ATRI - Advanced Technology Research Institute
                      Version: 1.0.0
       Github: https://github.com/institute-atri/wastrap`)
	println(bar)
}
