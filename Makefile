.PHONY: %.idl
%.idl:
	@$(DUMP_RUNNER) $@

.PHONY: %.h
%.h:
	@$(DUMP_RUNNER) $@

.PHONY: test-idl
test-idl:
	@for f in $(shell pwd)/idl/*.idl; do $(DUMP_RUNNER) $$f &>/dev/null && echo ok $$f || echo fail $$f; done

MSIDLPATH ?= $(shell pwd)/idl:$(shell pwd)/idl/h

DOCKER_IMAGE ?= ghcr.io/oiweiwei/midl-gen-go

# verbose mode
ifeq ($(DEBUG),1)
DOCKER_RUNNER_FLAGS += --verbose
endif

DUMP_RUNNER ?= docker run --rm \
	-v $(shell pwd):/work \
	-u $(shell id -u):$(shell id -g) \
	$(DOCKER_IMAGE) dump \
	-I /work/idl \
	-I /work/idl/h

DOCKER_RUNNER ?= docker run --rm \
	-v $(shell pwd):/work \
	-u $(shell id -u):$(shell id -g) \
	$(DOCKER_IMAGE) generate \
	--pkg "github.com/oiweiwei/go-msrpc/msrpc/" \
	-I /work/idl \
	-I /work/idl/h \
	--output /work/msrpc/ \
	--msdn-openspecs-indexer-file /work/msdn/index.yaml \
	$(DOCKER_RUNNER_FLAGS) \
	--msdn-openspecs-indexer-extra-file /work/msdn/extra.yaml \
	--msdn-openspecs-cache-dir /work/msdn/.cache/

.PHONY: %.go
%.go:
	$(DOCKER_RUNNER) $(basename $@).idl

.PHONY: %.json
%.json:
	$(DUMP_RUNNER) $(basename $@).idl

.PHONY: all
all:
	$(MAKE) \
		adts.go \
		adts/claims.go \
		bkrp.go \
		bpau.go \
		brwsa.go \
		capr.go \
		cmpo.go \
		cmrp.go \
		conv.go \
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
		dcom/mqac.go \
		dcom/oaut.go \
		dcom/ocspa.go \
		dcom/pla.go \
		dcom/rai.go \
		dcom/rdpesc.go \
		dcom/rrasm.go \
		dcom/rsmp.go \
		dcom/scmp.go \
		dcom/tpmvsc.go \
		dcom/tsrap.go \
		dcom/uamg.go \
		dcom/urlmon.go \
		dcom/vds.go \
		dcom/wcce.go \
		dcom/wmi.go \
		dcom/wsrm.go \
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
		frs1.go \
		frs2.go \
		fsrvp.go \
		gkdi.go \
		icpr.go \
		irp.go \
		lrec.go \
		lsad.go \
		lsat.go \
		mimicom.go \
		mgmt.go \
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
		tsts.go \
		w32t.go \
		wdsc.go \
		wkst.go

.PHONY: test
test:
	go test ./msrpc/...
