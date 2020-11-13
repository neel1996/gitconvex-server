get:
	go get
compile:
	go run server.go
build:
	ifeq(,$(wildcard ./ui)){
		git clone https://github.com/neel1996/gitconvex-ui.git ui/
	}
	cd ui
	npm install
	npm i -g create-react-app tailwindcss@1.6.0
	npm run build:tailwind
	npm run build
	mkdir -p ./dist