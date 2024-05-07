#! /bin/bash
curl https://01.gritlab.ax/assets/superhero/all.json | jq -r '.[] | select( .id == '$HERO_ID') | .connections.relatives' 