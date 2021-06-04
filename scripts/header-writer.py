import glob

files = glob.glob('/home/runner/work/Lux/Lux/*.go')

for file in files:
    if file.startswith("/pkg"):
        print("Excluded: " + file)
    else:
        print("Accepted: " + file)
