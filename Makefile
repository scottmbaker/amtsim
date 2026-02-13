.PHONY: all amtsim rpcsim clean

APP_NAME_AMTSIM := amtsim
APP_NAME_RPCSIM := rpcsim

ifeq ($(OS),Windows_NT)
    AMTSIM_BIN := $(APP_NAME_AMTSIM).exe
    RPCSIM_BIN := $(APP_NAME_RPCSIM).exe
else
    AMTSIM_BIN := $(APP_NAME_AMTSIM)
    RPCSIM_BIN := $(APP_NAME_RPCSIM)
endif

all: amtsim rpcsim

amtsim:
	go build -o $(AMTSIM_BIN) ./cmd/amtsim

rpcsim:
	go build -o $(RPCSIM_BIN) ./cmd/rpcsim

clean:
	rm -f $(AMTSIM_BIN) $(RPCSIM_BIN)

lint:
	golangci-lint run

