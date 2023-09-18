all:
	go run .

tailwind:
	tailwindcss -i views/style/index.css -o views/style/bundle.css --watch