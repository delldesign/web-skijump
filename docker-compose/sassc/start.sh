#!/bin/ash

sassc --sourcemap /src/application.scss /dst/application.css

inotifyd - /src:ydncw | while read line
do
  if [ -e /src/application.scss ]
  then
    sassc --sourcemap /src/application.scss /dst/application.css
  fi
done
