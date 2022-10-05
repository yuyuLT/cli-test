# idea-pot

This is a template for using gin and air in the Go language.

# How to use

1. touch app/src/go.mod

2. write go.mod 

module nothing-behind.com/sample_gin

go 1.19
 
3. docker-compose build

4. docker-compose run --rm app air init

5. docker-compose run --rm app go mod tidy

6. docker-compose up