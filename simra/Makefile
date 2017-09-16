
all:
	@make --no-print-directory graph
	@make --no-print-directory -C $(CURDIR)/internal/peer
	
graph:
	goviz -i github.com/pankona/gomo-simra/simra -d 2 | dot -Tsvg -o graph.svg

clean:
	rm -f graph.svg
