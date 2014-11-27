import os

def split_tbl(tblfile):
    colcount = 0
    with open(tblfile) as f:
        line = f.readline()
        colcount = len(line.split("|")) - 1
    
    filehdlr = [open("%s.%d" % (tblfile, i), "w") for i in range(colcount)]

    with open(tblfile) as f:
        for line in f.xreadlines():
            for i,x in enumerate(line.split("|")):
                if i < colcount:
                    h = filehdlr[i]
                    print >>h, x

    for f in filehdlr:
        f.close()

datadir = "./output"
for file in os.listdir(datadir):
    if file.endswith(".tbl"):
        file = os.path.join(datadir, file)
        print "Splitting %s" % file
        split_tbl(file)

print "Done"
