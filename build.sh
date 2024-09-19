#!/usr/bin/env zsh

acr_host="acr.internal.attains.cn"
acr_project="cloud"
acr_user="admin"
acr_pwd="Y1A3CubC1K2M4w9f"

remote_user="root"
remote_ip="103.45.130.157"

current_dir=$(git rev-parse --show-toplevel)
project_name=$(basename "$current_dir")
current_branch=$(git branch --show-current)
latest_commit_hash=$(git show -s --format='%h')
image_base_tag="$project_name-$current_branch:$latest_commit_hash"
image_tag="$acr_host/$acr_project/$image_base_tag"

echo "\033[32m branch: $current_branch \033[0m"
echo "\033[32m commit: $latest_commit_hash \033[0m"
echo "\033[32m name: $project_name \033[0m"
echo "\033[32m tag: $image_tag \033[0m"

echo "\033[32m start build image... \033[0m"
docker build -t "$image_tag" .
# shellcheck disable=SC2181
if [ $? -eq 0 ]; then
  echo "\033[32m image built! start push image... \033[0m"
else
  echo "\033[31m image build failed! \033[0m"
  exit
fi

docker login -u "$acr_user" -p "$acr_pwd" "$acr_host"
docker push "$image_tag"
# shellcheck disable=SC2181
if [ $? -eq 0 ]; then
  echo "\033[32m image pushed! start login remote... \033[0m"
else
  echo "\033[31m image push failed! \033[0m"
  exit
fi

ssh $remote_user@$remote_ip << remotessh
echo "remote login succeed!"
docker rm -f "$project_name"
docker rmi -f "$image_tag"
docker login -u "$acr_user" -p "$acr_pwd" "$acr_host"
docker pull "$image_tag"
docker run -p 9102:9102 --name "$project_name" -d $image_tag
exit
remotessh
# shellcheck disable=SC2181
if [ $? -eq 0 ]; then
  echo "ok"
else
  echo "\033[31m remote login failed! \033[0m"
fi