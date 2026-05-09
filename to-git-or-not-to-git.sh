curl https://acad.learn2earn.ng/assets/superhero/all.json | jq -r '.[] | select(.id== 170) | .name, .powerstats.durability, .appearance.gender'
