# web-skijump
Web browser ski jumping multiplayer game.

## What is web-skijump?
It is an open-source application which is creating to improve skills about
communication between many users.
This game will be created with following technologies:
	- docker (docker-compose),
	- nginx (it is dockerized),
	- sass (dockerized sassc),
	- hugo (dockerized),
	- golang 1.8 (as backend),
	- websockets (for realtime communication),
	- postgresql (as database),
	- bower (to receive dependencies),
	- and much more!

What you can learn by follow our little project?
 - How to write frontend with javascript + jQuery and canvas,
 - How to develop small / medium project,
 - How to use newest technologies in small / medium project in backend,
 - How to prepare project to handle many requests.

## How to run it?
	1. Install docker and docker-compose. It is neccessary to change
		 docker Storage Driver to 'devicemapper'.

	2. Install bower.

	3. Download and install dependencies with bower (bower install command).

	4. Download golang 1.8 or newer.

	5. Run docker-compose (docker-compose up command).

	6. Go to /backend and run project (go run main.go command).

	7. Enjoy! :)
