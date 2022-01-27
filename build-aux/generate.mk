# -*- fill-column: 102 -*-

# This file deals with creating files that get checked in to Git.  This is all grouped together in to
# one file, rather than being closer to the "subject matter" because this is a heinous thing.  Output
# files should not get checked in to Git -- every entry added to to this file is an affront to all
# that is good and proper.  As an exception, some of the Envoy-related stuff is allowed to live in
# envoy.mk, because that's a whole other bag of gross.

#
# `go mod tidy`
#
# This `go mod tidy` business only belongs in generate.mk because for the moment we're checking
# 'vendor/' in to Git.

go-mod-tidy:
.PHONY: go-mod-tidy

go-mod-tidy: go-mod-tidy/main
go-mod-tidy/main: $(OSS_HOME)/build-aux/go-version.txt
	rm -f go.sum
	GOFLAGS=-mod=mod go mod tidy -compat=$$(cut -d. -f1,2 < $<) -go=$$(cut -d. -f1,2 < $<)
.PHONY: go-mod-tidy/main

#
# The main `make generate` entrypoints and listings

# - Let $(generate/files) be a listing of all files or directories that `make generate` will create.
#
# - Let $(generate-fast/files) be the subset of $(generate/files) that can be generated "quickly".  A
#   file may NOT be considered fast if it uses the builder container, if it uses the network, or if it
#   needs to access the filesystem to evaluate the list of files (as the lines using `$(shell find
#   ...)` do).
#
# - Let $(generate/precious) be the subset of $(generate/files) that should not be deleted prior to
#   re-generation.

# Initialize
generate-fast/files  =
generate/files       = $(generate-fast/files)
generate/precious    =
# Whole directories with rules for each individual file in it
generate/files      += $(patsubst $(OSS_HOME)/api/%.proto,                   $(OSS_HOME)/pkg/api/%.pb.go                         , $(shell find $(OSS_HOME)/api/kat/              -name '*.proto')) $(OSS_HOME)/pkg/api/kat/
generate/files      += $(patsubst $(OSS_HOME)/api/%.proto,                   $(OSS_HOME)/pkg/api/%_grpc.pb.go                    , $(shell find $(OSS_HOME)/api/kat/              -name '*.proto'))
generate/files      += $(patsubst $(OSS_HOME)/api/%.proto,                   $(OSS_HOME)/pkg/api/%.pb.go                         , $(shell find $(OSS_HOME)/api/agent/            -name '*.proto')) $(OSS_HOME)/pkg/api/agent/
generate/files      += $(patsubst $(OSS_HOME)/api/%.proto,                   $(OSS_HOME)/pkg/api/%_grpc.pb.go                    , $(shell find $(OSS_HOME)/api/agent/            -name '*.proto'))
# Whole directories with one rule for the whole directory
generate/files      += $(OSS_HOME)/api/envoy/                # recipe in _cxx/envoy.mk
generate/files      += $(OSS_HOME)/pkg/api/envoy/            # recipe in _cxx/envoy.mk
generate/files      += $(OSS_HOME)/pkg/envoy-control-plane/  # recipe in _cxx/envoy.mk
# Individual files: Misc
generate/files      += $(OSS_HOME)/DEPENDENCIES.md
generate/files      += $(OSS_HOME)/DEPENDENCY_LICENSES.md
generate-fast/files += $(OSS_HOME)/CHANGELOG.md
generate-fast/files += $(OSS_HOME)/pkg/api/getambassador.io/v1/zz_generated.conversion.go
generate-fast/files += $(OSS_HOME)/pkg/api/getambassador.io/v1/zz_generated.conversion-spoke.go
generate-fast/files += $(OSS_HOME)/pkg/api/getambassador.io/v2/zz_generated.conversion.go
generate-fast/files += $(OSS_HOME)/pkg/api/getambassador.io/v2/zz_generated.conversion-spoke.go
generate-fast/files += $(OSS_HOME)/pkg/api/getambassador.io/v3alpha1/zz_generated.conversion-hub.go
# Individual files: YAML
generate-fast/files += $(OSS_HOME)/manifests/emissary/emissary-crds.yaml.in
generate-fast/files += $(OSS_HOME)/manifests/emissary/emissary-emissaryns.yaml.in
generate-fast/files += $(OSS_HOME)/manifests/emissary/emissary-defaultns.yaml.in
generate-fast/files += $(OSS_HOME)/pkg/api/getambassador.io/crds.yaml
generate-fast/files += $(OSS_HOME)/python/tests/integration/manifests/ambassador.yaml
generate-fast/files += $(OSS_HOME)/python/tests/integration/manifests/crds.yaml
generate-fast/files += $(OSS_HOME)/python/tests/integration/manifests/rbac_cluster_scope.yaml
generate-fast/files += $(OSS_HOME)/python/tests/integration/manifests/rbac_namespace_scope.yaml
# Individual files: Test TLS Certificates
generate-fast/files += $(OSS_HOME)/docker/test-auth/authsvc.crt
generate-fast/files += $(OSS_HOME)/docker/test-auth/authsvc.key
generate-fast/files += $(OSS_HOME)/docker/test-shadow/shadowsvc.crt
generate-fast/files += $(OSS_HOME)/docker/test-shadow/shadowsvc.key
generate-fast/files += $(OSS_HOME)/python/tests/selfsigned.py

generate: ## Update generated sources that get committed to Git
generate:
	$(MAKE) generate-clean
# This (generating specific targets early, then having a separate `_generate`) is a hack.  Because the
# full value of $(generate/files) is based on the listing of files in $(OSS_HOME)/api/, we need to
# make sure that those directories are fully populated before we evaluate the full $(generate/files).
	$(MAKE) $(OSS_HOME)/api/envoy
	$(MAKE) _generate
_generate:
	@echo '$(MAKE) $$(generate/files)'; $(MAKE) $(patsubst %/,%,$(generate/files))
.PHONY: generate _generate

generate-clean: ## Delete generated sources that get committed to Git
	rm -rf $(filter-out $(generate/precious),$(generate/files))
	find $(OSS_HOME)/pkg/api/getambassador.io -name 'zz_generated.*.go' -print -delete # generated as a side-effect of other files
.PHONY: generate-clean

generate-fast: ## Update the subset of generated-sources-that-get-committed-to-Git that can be updated quickly
generate-fast:
	$(MAKE) generate-fast-clean
	$(MAKE) $(patsubst %/,%,$(generate-fast/files))
.PHONY: generate-fast

generate-fast-clean: ## Delete the subset of generated-sources-that-get-committed-to-Git that can be updated quickly
	rm -rf $(filter-out $(generate/precious),$(generate-fast/files))
	find $(OSS_HOME)/pkg/api/getambassador.io -name 'zz_generated.*.go' -print -delete # generated as a side-effect of other files
.PHONY: generate-fast-clean

#
# Helper Make functions and variables

# Usage: $(call joinlist,SEPARATOR,LIST)
# Example: $(call joinlist,/,foo bar baz) => foo/bar/baz
joinlist=$(if $(word 2,$2),$(firstword $2)$1$(call joinlist,$1,$(wordlist 2,$(words $2),$2)),$2)

comma=,

gomoddir = $(shell cd $(OSS_HOME); go list -mod=readonly $1/... >/dev/null 2>/dev/null; go list -mod=readonly -m -f='{{.Dir}}' $1)

#
# Rules for downloading ("vendoring") sources from elsewhere


#
# `make generate` certificate generation

$(OSS_HOME)/docker/test-auth/authsvc.crt: $(tools/testcert-gen)
	$(tools/testcert-gen) --out-cert=$@ --out-key=/dev/null --hosts=authsvc.datawire.io
$(OSS_HOME)/docker/test-auth/authsvc.key: $(tools/testcert-gen)
	$(tools/testcert-gen) --out-cert=/dev/null --out-key=$@ --hosts=authsvc.datawire.io

$(OSS_HOME)/docker/test-shadow/shadowsvc.crt: $(tools/testcert-gen)
	$(tools/testcert-gen) --out-cert=$@ --out-key=/dev/null --hosts=demosvc.datawire.io
$(OSS_HOME)/docker/test-shadow/shadowsvc.key: $(tools/testcert-gen)
	$(tools/testcert-gen) --out-cert=/dev/null --out-key=$@ --hosts=demosvc.datawire.io

$(OSS_HOME)/python/tests/selfsigned.py: %: %.gen $(tools/testcert-gen)
	$@.gen $(tools/testcert-gen) >$@

#
# `make generate` protobuf rules

# proto_path is a list of where to look for .proto files.
proto_path += $(OSS_HOME)/api # input files must be within the path

# The "M{FOO}={BAR}" options map from .proto files to Go package names.
proto_options/go ?=
#proto_options/go += Mgoogle/protobuf/duration.proto=github.com/golang/protobuf/ptypes/duration

proto_options/go-grpc ?=

$(OSS_HOME)/pkg/api/%.pb.go: $(OSS_HOME)/api/%.proto $(tools/protoc) $(tools/protoc-gen-go)
	$(tools/protoc) \
	  $(addprefix --proto_path=,$(proto_path)) \
	  --plugin=$(tools/protoc-gen-go) --go_out=$(OSS_HOME)/pkg/api $(addprefix --go_opt=,$(proto_options/go)) \
	  $<
$(OSS_HOME)/pkg/api/%_grpc.pb.go: $(OSS_HOME)/api/%.proto $(tools/protoc) $(tools/protoc-gen-go-grpc)
	$(tools/protoc) \
	  $(addprefix --proto_path=,$(proto_path)) \
	  --plugin=$(tools/protoc-gen-go-grpc) --go-grpc_out=$(OSS_HOME)/pkg/api $(addprefix --go-grpc_opt=,$(proto_options/go-grpc)) \
	  $<

clean: _generate_clean
_generate_clean:
	rm -rf $(OSS_HOME)/_generate.tmp
.PHONY: _generate_clean

#
# `make generate` rules to update generated YAML files (and `zz_generated.*.go` Go files)

# Use `controller-gen` to generate Go & YAML
#
# - Enable a generator by setting the
#   `controller-gen/options/GENERATOR_NAME` variable (even to an empty
#   value).
# - Setting `controller-gen/output/GENERATOR_NAME` for an enabled
#   generator is optional; the default output for each enabled
#   generator is `dir=config/GENERATOR_NAME`.
# - It is invalid to set `controller-gen/output/GENERATOR_NAME` for a
#   generator that is not enabled.
#
#controller-gen/options/webhook     +=
#controller-gen/options/schemapatch += manifests=foo
#controller-gen/options/rbac        += roleName=ambassador
controller-gen/options/object      += headerFile=build-aux/copyright-boilerplate.go.txt
controller-gen/options/crd         += trivialVersions=false # Requires Kubernetes 1.13+
controller-gen/options/crd         += crdVersions=v1        # Requires Kubernetes 1.16+
controller-gen/output/crd           = dir=$@
$(OSS_HOME)/_generate.tmp/crds: $(tools/controller-gen) build-aux/copyright-boilerplate.go.txt FORCE
	@printf '  $(CYN)Running controller-gen$(END)\n'
	rm -rf $@
	mkdir -p $@
	cd $(OSS_HOME) && $(tools/controller-gen) \
	  $(foreach varname,$(sort $(filter controller-gen/options/%,$(.VARIABLES))), $(patsubst controller-gen/options/%,%,$(varname))$(if $(strip $($(varname))),:$(call joinlist,$(comma),$($(varname)))) ) \
	  $(foreach varname,$(sort $(filter controller-gen/output/%,$(.VARIABLES))), $(call joinlist,:,output $(patsubst controller-gen/output/%,%,$(varname)) $($(varname))) ) \
	  $(foreach p,$(wildcard ./pkg/api/getambassador.io/v*/),paths=$p...)

$(OSS_HOME)/%/zz_generated.conversion.go: $(tools/conversion-gen) build-aux/copyright-boilerplate.go.txt FORCE
	rm -f $@ $(@D)/*.scaffold.go
	GOPATH= GOFLAGS=-mod=mod $(tools/conversion-gen) \
	  --skip-unsafe \
	  --go-header-file=build-aux/copyright-boilerplate.go.txt \
	  --input-dirs=./$* \
	  --output-file-base=zz_generated.conversion
# Because v1 just aliases v2, conversion-gen will need to be able to see the v2 conversion functions
# when generating code for v1.
$(OSS_HOME)/pkg/api/getambassador.io/v1/zz_generated.conversion.go: $(OSS_HOME)/pkg/api/getambassador.io/v2/zz_generated.conversion.go

$(OSS_HOME)/%/handwritten.conversion.scaffold.go: $(OSS_HOME)/%/zz_generated.conversion.go
	{ \
	  gawk ' \
	    BEGIN { \
	      print("//+build scaffold"); \
	      print(""); \
	      print("package $(notdir $*)"); \
	      inFunc=0; \
	      curFunc=""; \
	    } \
	    match($$0, /^func auto(Convert_[^(]+)(\(.*)/, m) { \
	      if (inFunc) { \
	        print("  return nil"); \
	        print("}"); \
	        print(""); \
	        inFunc=0; \
	      } \
	      curFunc=\
	        "func " m[1] m[2] \
	        "  if err := auto" m[1] "(in, out, s); err != nil {" \
	        "    return err" \
	        "  }"; \
	    } \
	    /INFO|WARN/ { \
	      if (!inFunc) { \
	        print(curFunc); \
	        inFunc=1; \
	      } \
	      print; \
	    } \
	    END { \
	      if (inFunc) { \
	        print("  return nil"); \
	        print("}"); \
	      } \
	    }' | \
	  gofmt; \
	} <$< >$@

$(OSS_HOME)/%/zz_generated.conversion-hub.go: FORCE
	rm -f $@
	{ \
	  gawk ' \
	    BEGIN { \
	       print("package $(notdir $*)"); \
	       print(""); \
	       object=0; \
	    } \
	    /\/\/ \+kubebuilder:object:root=true/ { \
	       object=1; \
	    } \
	    /^type \S+ struct/ && object { \
	        if (!match($$2, /List$$/)) { \
	          print "func(*" $$2 ") Hub() {}"; \
	        } \
	        object=0; \
	    }' $(sort $(wildcard $(@D)/*.go)) | \
	  gofmt; \
	} >$@

$(OSS_HOME)/%/zz_generated.conversion-spoke.go: FORCE
	rm -f $@
	{ \
	  gawk ' \
	    BEGIN { \
	       print("package $(notdir $*)"); \
	       print(""); \
	       print("import ("); \
	       print("  k8sRuntime \"k8s.io/apimachinery/pkg/runtime\""); \
	       print("  \"sigs.k8s.io/controller-runtime/pkg/conversion\""); \
	       print(")"); \
	       print(""); \
	       print("func convertFrom(src conversion.Hub, dst conversion.Convertible) error {"); \
	       print("  scheme := conversionScheme()"); \
	       print("  var cur k8sRuntime.Object = src"); \
	       print("  for i := len(conversionIntermediates) - 1; i >= 0; i-- {"); \
	       print("    gv := conversionIntermediates[i]"); \
	       print("    var err error"); \
	       print("    cur, err = scheme.ConvertToVersion(cur, gv)"); \
	       print("    if err != nil {"); \
	       print("      return err"); \
	       print("    }"); \
	       print("  }"); \
	       print("  return scheme.Convert(cur, dst, nil)"); \
	       print("}"); \
	       print(""); \
	       print("func convertTo(src conversion.Convertible, dst conversion.Hub) error {"); \
	       print("  scheme := conversionScheme()"); \
	       print("  var cur k8sRuntime.Object = src"); \
	       print("  for _, gv := range conversionIntermediates {"); \
	       print("    var err error"); \
	       print("    cur, err = scheme.ConvertToVersion(cur, gv)"); \
	       print("    if err != nil {"); \
	       print("      return err"); \
	       print("    }"); \
	       print("  }"); \
	       print("  return scheme.Convert(cur, dst, nil)"); \
	       print("}"); \
	       print(""); \
	       object=0; \
	    } \
	    /\/\/ \+kubebuilder:object:root=true/ { \
	       object=1; \
	    } \
	    /^type \S+ struct/ && object { \
	        if (!match($$2, /List$$/)) { \
	          print "func(dst *" $$2 ") ConvertFrom(src conversion.Hub) error { return convertFrom(src, dst) }"; \
	          print "func(src *" $$2 ") ConvertTo(dst conversion.Hub) error { return convertTo(src, dst) }"; \
	        } \
	        object=0; \
	    }' $(sort $(wildcard $(@D)/*.go)) | \
	  gofmt; \
	} >$@

$(OSS_HOME)/manifests/emissary/emissary-crds.yaml.in: $(OSS_HOME)/_generate.tmp/crds $(tools/fix-crds)
	$(tools/fix-crds) --target=apiserver-kubectl $(sort $(wildcard $</*.yaml)) >$@

$(OSS_HOME)/python/tests/integration/manifests/crds.yaml: $(OSS_HOME)/_generate.tmp/crds $(tools/fix-crds)
	$(tools/fix-crds) --target=apiserver-kat $(sort $(wildcard $</*.yaml)) >$@

$(OSS_HOME)/pkg/api/getambassador.io/crds.yaml: $(OSS_HOME)/_generate.tmp/crds $(tools/fix-crds)
	$(tools/fix-crds) --target=internal-validator $(sort $(wildcard $</*.yaml)) >$@

# Names for all the helm-expanded.yaml files (and thence output.yaml and *.yaml.in files)
helm.name.emissary-emissaryns = emissary-ingress
helm.name.emissary-defaultns = emissary-ingress
helm.namespace.emissary-emissaryns = emissary
helm.namespace.emissary-defaultns = default
helm.name.emissary-emissaryns-agent = emissary-ingress
helm.namespace.emissary-emissaryns-agent = emissary
helm.name.emissary-defaultns-agent = emissary-ingress
helm.namespace.emissary-defaultns-agent = default
helm.name.emissary-emissaryns-migration = emissary-ingress
helm.namespace.emissary-emissaryns-migration = emissary
helm.name.emissary-defaultns-migration = emissary-ingress
helm.namespace.emissary-defaultns-migration = default

# IF YOU'RE LOOKING FOR *.yaml: recipes, look in main.mk.

$(OSS_HOME)/k8s-config/%/helm-expanded.yaml: \
  $(OSS_HOME)/k8s-config/%/values.yaml \
  $(boguschart_dir)
	helm template --namespace=$(helm.namespace.$*) --values=$(@D)/values.yaml $(or $(helm.name.$*),$*) $(boguschart_dir) >$@
$(OSS_HOME)/k8s-config/%/output.yaml: \
  $(OSS_HOME)/k8s-config/%/helm-expanded.yaml \
  $(OSS_HOME)/k8s-config/%/require.yaml \
  $(tools/filter-yaml)
	$(tools/filter-yaml) $(filter %/helm-expanded.yaml,$^) $(filter %/require.yaml,$^) >$@
k8s-config.clean:
	rm -f k8s-config/*/helm-expanded.yaml k8s-config/*/output.yaml
clean: k8s-config.clean

$(OSS_HOME)/manifests/emissary/%.yaml.in: $(OSS_HOME)/k8s-config/%/output.yaml
	cp $< $@

$(OSS_HOME)/python/tests/integration/manifests/%.yaml: $(OSS_HOME)/k8s-config/kat-%/output.yaml
	sed -e 's/«/{/g' -e 's/»/}/g' -e 's/♯.*//g' -e 's/- ←//g' <$< >$@

$(OSS_HOME)/python/tests/integration/manifests/rbac_cluster_scope.yaml: $(OSS_HOME)/k8s-config/kat-rbac-multinamespace/output.yaml
	sed -e 's/«/{/g' -e 's/»/}/g' -e 's/♯.*//g' -e 's/- ←//g' <$< >$@

$(OSS_HOME)/python/tests/integration/manifests/rbac_namespace_scope.yaml: $(OSS_HOME)/k8s-config/kat-rbac-singlenamespace/output.yaml
	sed -e 's/«/{/g' -e 's/»/}/g' -e 's/♯.*//g' -e 's/- ←//g' <$< >$@

#
# Generate report on dependencies

$(OSS_HOME)/build-aux/pip-show.txt: docker/base-pip.docker.tag.local
	docker run --rm "$$(cat docker/base-pip.docker)" sh -c 'pip freeze --exclude-editable | cut -d= -f1 | xargs pip show' > $@
clean: build-aux/pip-show.txt.rm

$(OSS_HOME)/build-aux/go-version.txt: $(_go-version/deps)
	$(_go-version/cmd) > $@
clean: build-aux/go-version.txt.rm

$(OSS_HOME)/build-aux/py-version.txt: docker/base-python/Dockerfile
	{ grep -o 'python3=\S*' | cut -d= -f2; } < $< > $@
clean: build-aux/py-version.txt.rm

$(OSS_HOME)/build-aux/go1%.src.tar.gz:
	curl -o $@ --fail -L https://dl.google.com/go/$(@F)
build-aux/go.src.tar.gz.clean:
	rm -f build-aux/go1*.src.tar.gz
clobber: build-aux/go.src.tar.gz.clean

$(OSS_HOME)/DEPENDENCIES.md: $(tools/go-mkopensource) $(tools/py-mkopensource) $(OSS_HOME)/build-aux/go-version.txt $(OSS_HOME)/build-aux/pip-show.txt
	$(MAKE) $(OSS_HOME)/build-aux/go$$(cat $(OSS_HOME)/build-aux/go-version.txt).src.tar.gz
	set -e; { \
		cd $(OSS_HOME); \
		$(tools/go-mkopensource) --output-format=txt --package=mod --application-type=external --gotar=build-aux/go$$(cat $(OSS_HOME)/build-aux/go-version.txt).src.tar.gz; \
		echo; \
		{ sed 's/^---$$//' $(OSS_HOME)/build-aux/pip-show.txt; echo; } | $(tools/py-mkopensource); \
	} > $@

$(OSS_HOME)/DEPENDENCY_LICENSES.md: $(tools/go-mkopensource) $(tools/py-mkopensource) $(OSS_HOME)/build-aux/go-version.txt $(OSS_HOME)/build-aux/pip-show.txt
	$(MAKE) $(OSS_HOME)/build-aux/go$$(cat $(OSS_HOME)/build-aux/go-version.txt).src.tar.gz
	echo -e "Emissary-ingress incorporates Free and Open Source software under the following licenses:\n" > $@
	set -e; { \
		cd $(OSS_HOME); \
		$(tools/go-mkopensource) --output-format=txt --package=mod --output-type=json --application-type=external --gotar=build-aux/go$$(cat $(OSS_HOME)/build-aux/go-version.txt).src.tar.gz | jq -r '.licenseInfo | to_entries | .[] | "* [" + .key + "](" + .value + ")"' ; \
		{ sed 's/^---$$//' $(OSS_HOME)/build-aux/pip-show.txt; echo; } | $(tools/py-mkopensource) --output-type=json | jq -r '.licenseInfo | to_entries | .[] | "* [" + .key + "](" + .value + ")"'; \
	} | sort | uniq | sed -e 's/\[\([^]]*\)]()/\1/' >> $@

#
# Misc. other `make generate` rules

$(OSS_HOME)/CHANGELOG.md: $(OSS_HOME)/docs/CHANGELOG.tpl $(OSS_HOME)/docs/releaseNotes.yml
	docker run --rm \
	  -v $(OSS_HOME)/docs/CHANGELOG.tpl:/tmp/CHANGELOG.tpl \
	  -v $(OSS_HOME)/docs/releaseNotes.yml:/tmp/releaseNotes.yml \
	  hairyhenderson/gomplate --verbose --file /tmp/CHANGELOG.tpl --datasource relnotes=/tmp/releaseNotes.yml > CHANGELOG.md
