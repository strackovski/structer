all:
	./scripts/build.sh
publish:
	./scripts/build.sh publish
install:
	./scripts/install.sh
web_install:
	./scripts/install_web.sh