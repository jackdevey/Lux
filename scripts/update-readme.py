import requests
import json

url = "https://bandev.uk/api/lux/commands.json"

data = json.loads(requests.get(url).text)

for command in data["commands"]:
  print(command["name"])
