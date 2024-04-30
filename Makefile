default: build

.PHONY: gen
gen:
	@goyacc -v midl/y.out -o midl/parse.go -p RPC midl/parse.y

.PHONY: gen-oem
gen-oem:
	@go generate ./...

.PHONY: bin
bin:
	@mkdir -p bin/

.PHONY: build
build: bin gen
	@CGO_ENABLED=0 go build -o bin/parse codegen/main.go

.PHONY: %.idl
%.idl:
	@./bin/parse -j $@

.PHONY: %.h
%.h:
	@./bin/parse -j $@

.PHONY: test-idl
test-idl:
	@for f in $(shell pwd)/idl/*.idl; do ./bin/parse -j $$f &>/dev/null && echo ok $$f || echo fail $$f; done

FORMAT ?= true

MSIDLPATH ?= $(shell pwd)/idl:$(shell pwd)/idl/h

.PHONY: %.go
%.go:
	MSIDLPATH=$(MSIDLPATH) ./bin/parse \
		-I "github.com/oiweiwei/go-msrpc/msrpc/" \
		-dir ./msrpc/ \
		-format=$(FORMAT) \
		-doc-cache ./.cache/doc/ \
		-f "$(basename $@).idl"

.PHONY: %.json
%.json:
	MSIDLPATH=$(MSIDLPATH) ./bin/parse \
		-j -f "$(basename $@).idl"


.PHONY: all
all:
	$(MAKE) build \
		adts.go \
		adts/claims.go \
		bkrp.go \
		bpau.go \
		brwsa.go \
		capr.go \
		cmpo.go \
		cmrp.go \
		dcetypes.go \
		dcom.go \
		dcom/adtg.go \
		dcom/ccfg.go \
		dcom/com.go \
		dcom/coma.go \
		dcom/comev.go \
		dcom/comt.go \
		dcom/csra.go \
		dcom/csvp.go \
		dcom/dfsrh.go \
		dcom/dmrp.go \
		dcom/fsrm.go \
		dcom/iisa.go \
		dcom/iiss.go \
		dcom/imsa.go \
		dcom/ioi.go \
		dcom/rai.go \
		dcom/rdpesc.go \
		dcom/oaut.go \
		dcom/vds.go \
		dcom/wmi.go \
		dcom/wcce.go \
		dfsnm.go \
		dhcpm.go \
		dltm.go \
		dltw.go \
		dnsp.go \
		dnsp/record.go \
		drsr.go \
		dssp.go \
		dtyp.go \
		eerr.go \
		epm.go \
		even.go \
		even6.go \
		fasp.go \
		fax.go \
		fsrvp.go \
		gkdi.go \
		icpr.go \
		irp.go \
		lrec.go \
		lsad.go \
		lsat.go \
		mqds.go \
		mqmp.go \
		mqmq.go \
		mqmr.go \
		mqqp.go \
		mqrr.go \
		msrp.go \
		negoex.go \
		nrpc.go \
		nspi.go \
		oxabref.go \
		oxcrpc.go \
		pac.go \
		pan.go \
		par.go \
		pcq.go \
		raa.go \
		raiw.go \
		rpcl.go \
		rprn.go \
		rrp.go \
		rsp.go \
		samr.go \
		sch.go \
		scmr.go \
		srvs.go \
		ssp.go \
		swn.go \
		trp.go \
		tsch.go \
		tsgu.go \
		w32t.go \
		wkst.go

.PHONY: test
test:
	go test ./example/...
