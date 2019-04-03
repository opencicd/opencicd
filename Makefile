## Copyright 2019 OpenCICD Authors
##
## Licensed under the Apache License, Version 2.0 (the "License");
## you may not use this file except in compliance with the License.
## You may obtain a copy of the License at
##
##     http://www.apache.org/licenses/LICENSE-2.0
##
## Unless required by applicable law or agreed to in writing, software
## distributed under the License is distributed on an "AS IS" BASIS,
## WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
## See the License for the specific language governing permissions and
## limitations under the License.

# Include
include ./Makefile.defs

# Target all
.PHONY: all
all: build

# Target proto
.PHONY: proto
proto:
	$(PROTOTOOL) all

# Target clean
.PHONY: build
build: proto
	@for i in $(CMD_SUBDIRS); do \
		$(MAKE) -C $$i build; \
	done

# Target test
.PHONY: test
test:
	$(GINKGO) -r --randomizeAllSpecs --randomizeSuites --failOnPending --cover --trace --race --progress --compilers=2

# Target test-watch
.PHONY: test-watch
test-watch:
	$(GINKGO) watch -r

# Target clean
.PHONY: clean
clean:
	@echo Removing directory: $(OUT_DIR)
	@rm -Rf $(OUT_DIR)

	@echo Deleting generated protobuf files
	@find ./api -name *.pb.go -delete

	@echo Deleting cover files
	@find ./ -name *.coverprofile -delete
	
	@for i in $(CMD_SUBDIRS); do \
		$(MAKE) -C $$i clean; \
	done