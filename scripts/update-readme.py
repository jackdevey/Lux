import requests

url = "https://bandev.uk/api/lux/commands.json"

data = json.loads(requests.get(url))

for command in data["commands"]:
  print(command["name"])
