

#define all constant and variables
TargetName 	:= platform-realtime-plugin
BaseFolder	:=	../Build
BuildFolder	:=	$(BaseFolder)/$(TargetName)
TargetZip	:=	.zip
LinterFile	:=	static-analysis.xml
LogFile	    :=	RTPlugin.log
CoverAll := cover-all
CoverageFile	:=	cover-cobertura
sources	:=	$(wildcard *.go)
vendor := vendor



#Variable for getting the Revesion of the HEAD of the Repository for Automation-Engine
Build=`git rev-parse HEAD`

# go build command
build = go build  -o $(BuildFolder)/$(TargetName) ./


#Zipping the Target Folder
zip		=	cd $(BaseFolder) && zip -r $(TargetName)$(TargetZip) *

SERVICEPATH=$(GOPATH)/src/github.com/ContinuumLLC/$(TargetName)/services
SERVICEPACKPATH=$(SERVICEPATH)/continuum
SERVICEBINPATH=$(SERVICEPACKPATH)/realtimeService
SERVICELOGPATH=$(SERVICEPACKPATH)/log
SERVICECONFIGPATH=$(SERVICEPACKPATH)/config
SERVICEPOLICYPATH=$(SERVICECONFIGPATH)/policy
SERVICEPACKNAME=continuum.zip
SERVICEPACKTARGET=continuum
SERVICESCRIPTS=scripts

SOURCEPATH=$(GOPATH)/src/github.com/ContinuumLLC/$(TargetName)/src





#https://github.com/ContinuumLLC/Platform-Realtime-Service/wiki/Glossary [Understanding Phony]
.PHONY:	all	clean

#create directory for output file
dir:
	mkdir -p $(BuildFolder)

#Setup the -ldflag option for go build here, interpolate the variable values
#Ideally we are Passing value of the Head revesion of the code Repository to the main.go file to set the Build variable declared in main.go source file
LDFLAGS=-ldflags "-X main.Build=${Build}"

#build binary
#need to check what does it do??
BuildCmd:	$(sources)
	$(call build)
	$(call zip)
	
#Loading module dependencies through Glide (glide.yaml)
dependencies:
	glide install

#install binary
install:
	go install ${LDFLAGS}
		
	
#Installing Helpers (Static Code analyzer and Glide)	
install-helpers:
	echo "Installing GoMetaLinter"
	go get -u github.com/alecthomas/gometalinter
	echo "Installing Glide"
	curl https://glide.sh/get | sh	

	
# Environ commands
setup-env: install-helpers dependencies 


#Target to peroform all action in single step
all:	clean	dir dependencies BuildCmd test cover code-quality pack-service output-formatted

#Target to clean
clean:
	rm -rf $(BaseFolder)/$(CoverAll).out
	rm -rf $(BaseFolder)/$(CoverAll).html
	rm -rf $(BaseFolder)/$(CoverageFile).xml
	rm -rf $(BaseFolder)/$(LinterFile)
	rm -rf $(BaseFolder)/$(TargetName)$(TargetZip)
	rm -rf $(BaseFolder)/$(TargetName)
	rm -rf $(SERVICEBINPATH)/$(TargetName)
	rm -rf $(SERVICELOGPATH)/$(LogFile)
	rm -rf $(SERVICEPATH)/$(SERVICEPACKNAME)
	rm -rf ./$(vendor)




#Target for consolidate code quality check, For more info check below link:
#https://continuum.atlassian.net/wiki/spaces/C2E/pages/170754059/GO+Best+Practices#GOBestPractices-CodeQuality-GoMetaLinter
code-quality:
	- gometalinter --vendor --skip=mocks --disable=dupl --disable=gotype --disable=errcheck --disable=gas --deadline=5000s --checkstyle --sort=linter ./... > $(BaseFolder)/$(LinterFile) 

#better reprsentation for code coverage output (Html Format)	
output-formatted:
	go tool cover -html=./$(BaseFolder)/$(CoverAll).out -o ./$(BaseFolder)/$(CoverAll).html
	go tool cover -func=./$(BaseFolder)/$(CoverAll).out

pack-service:
	mkdir -p $(SERVICEBINPATH)
	mkdir -p $(SERVICELOGPATH)
	
	cp $(BaseFolder)/$(TargetName)/$(TargetName)  $(SERVICEBINPATH)
	
	cd $(SERVICEPATH); rm -f $(SERVICEPACKNAME); zip -r $(SERVICEPACKNAME) $(SERVICEPACKTARGET)
