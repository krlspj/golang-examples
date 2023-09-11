.PHONY: all clean

CC = gcc
CFLAGS = -Wall -shared -fPIC

SRC = mylibrary.c

OUT_DIR = .
OUT_LIB = mylibrary.so

all: $(OUT_LIB)

$(OUT_LIB): $(SRC)
	$(CC) $(CFLAGS) -o $(OUT_DIR)/$@ $^

clean:
	rm -f $(OUT_DIR)/$(OUT_LIB)
