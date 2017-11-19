DIRS := $(wildcard */.)

BUILDDIRS = $(DIRS:%=build-%)
INSTALLDIRS = $(DIRS:%=install-%)
CHECKDIRS = $(DIRS:%=check-%)
FMTDIRS = $(DIRS:%=fmt-%)
TESTDIRS = $(DIRS:%=test-%)

all: $(BUILDDIRS)
$(DIRS): $(BUILDDIRS)
$(BUILDDIRS):
	$(MAKE) -C $(@:build-%=%)

install: $(INSTALLDIRS) all
$(INSTALLDIRS):
	$(MAKE) -C $(@:install-%=%) install

check: $(CHECKDIRS) all
$(CHECKDIRS): 
	$(MAKE) -C $(@:check-%=%) check

fmt: $(FMTDIRS) all
$(FMTDIRS): 
	$(MAKE) -C $(@:fmt-%=%) fmt

clean: $(CLEANDIRS)
$(CLEANDIRS): 
	$(MAKE) -C $(@:clean-%=%) clean

.PHONY: subdirs $(DIRS)
.PHONY: subdirs $(BUILDDIRS)
.PHONY: subdirs $(INSTALLDIRS)
.PHONY: subdirs $(CHECKDIRS)
.PHONY: subdirs $(FMTDIRS)
.PHONY: subdirs $(CLEANDIRS)
.PHONY: all install check fmt test
