binary:
	cd src && go build .

install:
	cd src && sudo mv dtcp /usr/bin