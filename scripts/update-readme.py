import requests
import json

url = "https://bandev.uk/api/lux/commands.json"

data = json.loads(requests.get(url).text)

readme = open("/home/runner/work/Lux/Lux/README.md", "r").read()

a1 = readme.split("<ul id='EDTCMDS'>")[1]
a2 = a1.split("</ul>")[0]
commands = a2.split("<li>")

i = 0

for command in data["commands"]:
  commands[i] = "<code>"+command["examples"][0]+"</code><p>"+command["description"]+"</p>"
  print(commands[i])
  i = i + 1
