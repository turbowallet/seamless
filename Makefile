.PHONY: dep build docker release install test backup

run: 
	docker run -it --rm  -v $$PWD:/app -w /app -p 3000:3000 ruby ruby -run -e httpd . -p 3000
