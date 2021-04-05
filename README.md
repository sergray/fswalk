# Unix Filesystem Walk

Quick implementation to compare speed and ease of Python and Go solutions for recursive enumeration of the files and printing out data in CSV format.

## Quickstart

	$ time python3 pyfswalk/fswalk.py /home 2>/dev/null | wc -l

	$ time go run gofswalk/fswalk.go /home | wc -l
