import glob

files_tmp = glob.glob('/home/runner/work/Lux/Lux/*/*/*.go')

for file_tmp in files_tmp:
    if file_tmp.startswith("/home/runner/work/Lux/Lux/pkg"):
        print("Excluded: " + file_tmp)
    else:
        print("Accepted: " + file_tmp)
        files.append(file_tmp)
        
for file in files:
    text = open(file).read()
    if text.startswith("""/**

Lux
Copyright (C) 2021  BanDev

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.

*/"""):
        print("Header found")
