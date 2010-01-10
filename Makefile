include $(GOROOT)/src/Make.$(GOARCH)
include $(GOROOT)/src/Make.cmd

CLEANFILES=netsnail

all: netsnail

netsnail: main.$O
	$(QUOTED_GOBIN)/$(LD) -o netsnail main.$O

main.$O: netsnail.$O src/main.go
	$(QUOTED_GOBIN)/$(GC) -I . -o main.$O src/main.go

netsnail.$O: src/config.go
	$(QUOTED_GOBIN)/$(GC) -o netsnail.$O src/config.go

