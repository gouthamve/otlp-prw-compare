import json

targets = []

for i in range(0, 100):
    target = {
        "labels": {
            "instance": "instance-{}".format(i)
        },
        "targets": [
            "avalanche:9001",
        ]
    }
    targets.append(target)


json_object = json.dumps(targets, indent=2)

with open("./prometheus/targets.json", "w+") as outfile:
    outfile.write(json_object)

with open("./otelcol/targets.json", "w+") as outfile:
    outfile.write(json_object)