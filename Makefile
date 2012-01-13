include $(GOROOT)/src/Make.inc
CLEANFILES=netsnail

NETSNAIL_SRC=config.go delayproxy.go netsnailhelpers.go

all: netsnail

netsnail: main.$O
	$(QUOTED_GOBIN)/$(LD) -o netsnail main.$O

main.$O: netsnail.$O main.go
	$(QUOTED_GOBIN)/$(GC) -I . -o main.$O main.go

netsnail.$O: $(NETSNAIL_SRC)
	$(QUOTED_GOBIN)/$(GC) -o netsnail.$O $(NETSNAIL_SRC)

include $(GOROOT)/src/Make.cmd
