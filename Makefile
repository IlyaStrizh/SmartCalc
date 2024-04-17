CC= ~/go/bin/fyne
NAME= SmartCalc_v3.0

ifeq ($(shell uname -s), Linux)
OS= linux

else ifeq ($(shell uname -s), Darwin)
OS= darwin

endif

all: build test

deps:
	@command -v $(CC) >/dev/null 2>&1 || { echo >&2 "Fyne is not installed. Installing..."; go install fyne.io/fyne/v2/cmd/fyne@latest; }

install: deps build
	@cd cmd && $(CC) package -os $(OS) -name $(NAME) -icon ../images/calc.png && \
	cp $(PWD)/internal/model/so/calc.so $(PWD)/cmd/$(NAME).app/Contents/Resources/. && \
    cp $(PWD)/md/info.md $(PWD)/cmd/$(NAME).app/Contents/Resources/. && \
    cp $(PWD)/txt/history.txt $(PWD)/cmd/$(NAME).app/Contents/Resources/. && \
	ln -s $(PWD)/cmd/$(NAME).app ~/Desktop/$(NAME).app

dist: deps build
	@cd cmd && $(CC) package -os $(OS) -name $(NAME) -icon ../images/calc.png && \
	cp $(PWD)/internal/model/so/calc.so $(PWD)/cmd/$(NAME).app/Contents/Resources/. && \
    cp $(PWD)/md/info.md $(PWD)/cmd/$(NAME).app/Contents/Resources/. && \
    cp $(PWD)/txt/history.txt $(PWD)/cmd/$(NAME).app/Contents/Resources/. && \
	cd $(PWD)/cmd && tar -czf ~/Desktop/$(NAME).tar.gz $(NAME).app

uninstall: clean
	rm -rf $(PWD)/cmd/$(NAME).app
	rm -f ~/Desktop/$(NAME).app ~/Desktop/$(NAME).app.tar.gz

build:
	cd $(PWD)/internal/model/so && go build -buildmode=plugin -o calc.so

test: clean build
	cd $(PWD)/test && go test

clean:
	rm -f $(PWD)/internal/model/so/calc.so