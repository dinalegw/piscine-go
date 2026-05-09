curl  https://acad.learn2earn.ng/assets/superhero/all.json | jq '.[] | select(.id== 70) | .name'
