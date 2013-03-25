import sys
from time import time

t0 = time()
f = open(sys.argv[1], 'r')
lookup = {}
i = 0
for line in f.xreadlines():
    for w in line.strip().split():
        i += 1
        if w in lookup:
            lookup[w] += 1
        else:
            lookup[w] = 1
f.close()
print "distinct words = ", len(lookup.keys())
print "processed %d words in %.2f seconds" % (i, time()-t0)
