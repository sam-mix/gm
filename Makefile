.PHONY:
all: clean build run 


.PHONY:
clean:
	@rm -rf *.exe

.PHONY:
build:
	@go build .

.PHONY:
run:
	@./gm


.PHONY:
cfg:
	@cp conf/config.default.toml conf/config.toml 
