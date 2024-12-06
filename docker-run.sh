#!/bin/bash
user="kokoro"
name="automod-bot"

TOKEN=${TOKEN:-$(read -p "Enter TOKEN: " token && echo "$token")}

chmod 777 config/banned_words.txt

docker build \
    --build-arg TOKEN="$TOKEN" \
    $@ -t $user/$name:latest . || exit

[ "$(docker ps | grep $name)" ] && docker kill $name
[ "$(docker ps -a | grep $name)" ] && docker rm $name

docker run \
	-itd \
	-u $(id -u):$(id -g) \
	--name $name \
    --network host \
	--restart=always \
	--volume $(pwd)/config/banned_words.txt:/app/config/banned_words.txt \
	$user/$name:latest