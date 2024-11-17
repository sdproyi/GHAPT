.PHONY: templ air dev kill tailwind
templ:
	templ generate --watch --proxy="http://localhost:8080"
tailwind:
	tailwindcss -i ./tailwind.css -o ./static/style/styles.css --watch --silent
air:
	air -c .air.toml

dev:
	$(MAKE) -j2 templ air

kill:
	bun kill-port 8080