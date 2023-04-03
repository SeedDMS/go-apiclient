PKGNAME=golang-seeddms-seeddms-apiclient
VERSION=0.0.2

test:
	go test -v .

build:
	go build seeddms.org/seeddms/apiclient

dist:
	rm -rf ${PKGNAME}-${VERSION}
	mkdir ${PKGNAME}-${VERSION}
	cp -r *.go Makefile go.mod go.sum ${PKGNAME}-${VERSION}
	mkdir -p ${PKGNAME}-${VERSION}/testdata/fixtures
	cp testdata/fixtures/*.json ${PKGNAME}-${VERSION}/testdata/fixtures
	tar czvf ${PKGNAME}-${VERSION}.tar.gz ${PKGNAME}-${VERSION}
	rm -rf ${PKGNAME}-${VERSION}

debian: dist
	mv ${PKGNAME}-${VERSION}.tar.gz ../${PKGNAME}_${VERSION}.orig.tar.gz
	debuild

