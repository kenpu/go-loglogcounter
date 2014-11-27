import sys
from time import time

# this code is take from:
# http://blog.notdot.net/2012/09/Dam-Cool-Algorithms-Cardinality-Estimation

def trailing_zeroes(num):
  """Counts the number of trailing 0 bits in num."""
  if num == 0:
    return 32 # Assumes 32 bit integer inputs!
  p = 0
  while (num >> p) & 1 == 0:
    p += 1
  return p

def estimate_cardinality(values, k, max_zeroes=None):
  """Estimates the number of unique elements in the input set values.

  Arguments:
    values: An iterator of hashable elements to estimate the cardinality of.
    k: The number of bits of hash to use as a bucket number; there will be 2**k buckets.
  """
  num_buckets = 2 ** k

  if not max_zeroes:
      max_zeroes = [0] * num_buckets
  else:
      if not len(max_zeroes) == num_buckets:
          raise Exception("k and len(max_zeroes) mismatch")

  if values:
      for value in values:
        h = hash(value)
        bucket = h & (num_buckets - 1) # Mask out the k least significant bits as bucket ID
        bucket_hash = h >> k
        max_zeroes[bucket] = max(max_zeroes[bucket], trailing_zeroes(bucket_hash))

  est = 2 ** (float(sum(max_zeroes)) / num_buckets) * num_buckets * 0.79402

  return est, max_zeroes

def value_stream(file):
    with open(file) as f:
        for x in f.xreadlines():
            yield x

# merge two max_zeros tables

def estimate_union(t1, t2, k):
    if not len(t1) == len(t2):
        raise Exception("two table lengths mismatch")

    t = [max(i, j) for (i,j) in zip(t1, t2)]
    return estimate_cardinality(None, k, t)

if not sys.argv[1:]:
    print "Usage: estimate.py [k] [files...]"
    print "k = 5 is a good choice to start with..."
    sys.exit(0)

k = int(sys.argv[1])

filenames = sys.argv[2:]

start = time()

# build the table set
tables = dict()
for f in filenames:
    v = value_stream(f)
    e, t = estimate_cardinality(v, k)
    print "INDEXED %s WITH EST %.1f" % (f, e)
    tables[f] = dict(e = e, t = t)

dur = time() - start

print "indexed %d files in %.2f seconds" % (len(filenames), dur)

# Compute pairwise union:

result = ["F1", "F2", "E1", "E2", "E", "M=%d" % k]
print ",".join(result)

start = time()
i = 0
for f1 in filenames:
    for f2 in filenames:
        if f1 < f2:
            e1 = tables[f1]["e"]
            e2 = tables[f2]["e"]
            e, t = estimate_union(tables[f1]["t"], tables[f2]["t"], k)
            sim = max(e1, e2)/e
            result = [f1, f2, "%.1f" % e1, "%.1f" % e2, "%.1f" % e]

            print ",".join(result)
            i += 1

dur = time() - start
print "JOINED %d PAIRS IN %.4f seconds" % (i, dur)
