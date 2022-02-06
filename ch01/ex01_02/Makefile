ARG = こんにちは 世界

TARGET = echo
SRC = echo.go
GO = go

.PHONY: run
run: $(TARGET)
	./$(TARGET) ${ARG}

$(TARGET):
	$(GO) build -o $(TARGET) $(SRC)

clean:
	rm $(TARGET)

