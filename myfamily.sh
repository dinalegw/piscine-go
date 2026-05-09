#! /bin/bash
curl -s "https://acad.learn2earn.ng/assets/superhero/all.json" | jq ".[] | select(.id == $HERO_ID) | .connections.relatives" | sed 's/"//g'