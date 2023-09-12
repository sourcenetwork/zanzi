src := $(wildcard *.dot)
images := $(src:%.dot=%.png)

.PHONY: all
all: $(images)

%.png: %.dot
	dot -Tpng $< > $@
