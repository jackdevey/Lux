import os

files = os.listdir('/home/runner/work/Lux/Lux')

for file in files:
    if file.startswith("/pkg") OR NOT file.endswith('.go'):
        print("Excluded: " + file)
