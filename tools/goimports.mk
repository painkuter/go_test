FILESDIFFMASTER = $(shell git diff --name-only master | grep /*\.go$/)
PROJECTFOLDERNAME = $(shell basename `pwd`)

.PHONY: goimports-sort-diff-master
goimports-sort-diff-master:
	$Q \
	for f in $(FILESDIFFMASTER); do \
	    if grep -q -E '^// Code generated .* DO NOT EDIT\.$$' $$f; then \
	      continue; \
	    fi; \
	    diff_file=$$(goimports -local $(PROJECTFOLDERNAME) -d $$f); \
        if [ "$${#diff_file}" = 0 ]; then \
        	newline_count=$$(cat $$f | perl -e 'while(<STDIN>){ $$in .= $$_ }; ($$r) = $$in =~ /(import \(.*?(\s{3,}).*?\))/gms; if($$r) { @n = $$r =~ /(\s{3,})/g; print scalar(@n) } else { print 0 }'); \
        	if [ $$newline_count -lt "3" ]; then \
        		continue; \
        	fi; \
        fi; \
	    sed '/^import/,/)/ { /^[\s \t]*$$/ d; }' $$f > $$f.tmp; mv $$f.tmp $$f; \
	    goimports -local $(PROJECTFOLDERNAME) -w $$f $$f; \
	    echo "File '$$f' is changed!"; \
  	done

.PHONY: goimports-check-sort-diff-master
goimports-check-sort-diff-master:
	$Q \
	is_exists_diff=""; \
	for f in $(FILESDIFFMASTER); do \
	    if grep -q -E '^// Code generated .* DO NOT EDIT\.$$' $$f; then \
	      continue; \
	    fi; \
	    diff_file=$$(goimports -local $(PROJECTFOLDERNAME) -d $$f); \
		if [ "$${#diff_file}" != 0 ]; then \
		  	is_exists_diff="true"; \
		  	echo "File '$$f' is needed to sort imports!"; \
		  	continue; \
		fi; \
		newline_count=$$(cat $$f | perl -e 'while(<STDIN>){ $$in .= $$_ }; ($$r) = $$in =~ /(import \(.*?(\s{3,}).*?\))/gms; if($$r) { @n = $$r =~ /(\s{3,})/g; print scalar(@n) } else { print 0 }'); \
		echo "$$newline_count newline_count"; \
		if [ $$newline_count -gt "2" ]; then \
		  	is_exists_diff="true"; \
          	echo "File '$$f' is needed to sort imports!"; \
		fi; \
  	done; \
  	if [ $$is_exists_diff ]; then \
  		exit 1; \
  	else \
  		echo "Not files to sort!"; \
  	fi
