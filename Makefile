SRC_DIR := src
BIN_DIR := bin
LDFLAGS := -w -s 
OBJ_NAME := crawl
PORT := 8080

build:
	cd $(SRC_DIR); go build -v -ldflags "$(LDFLAGS)" -o $(OBJ_NAME);
	mv $(SRC_DIR)/$(OBJ_NAME) $(BIN_DIR) 

run:
	./$(BIN_DIR)/$(OBJ_NAME)

clean:
	rm $(BIN_DIR)/$(OBJ_NAME)

godoc: 
	godoc -http=:$(PORT)

