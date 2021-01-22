#
import csv
import os
import sys

columns = 'path mode uid gid size atime mtime ctime'.split(' ')

writer = csv.writer(sys.stdout)


for dest in sys.argv[1:]:
	if not os.path.isdir(dest):
		continue
	for dirpath, dirnames, filenames in os.walk(dest):
		for name in dirnames + filenames:
			path = os.path.join(dirpath, name)
			try:
				stat = os.stat(path, follow_symlinks=False)
			except Exception as exc:
				print(exc, file=sys.stderr)
			try:
				writer.writerow(
					(path, stat.st_mode, stat.st_uid, stat.st_gid, stat.st_size, stat.st_atime, stat.st_mtime, stat.st_ctime)
				)
			except Exception as exc:
				print(exc, file=sys.stderr)
