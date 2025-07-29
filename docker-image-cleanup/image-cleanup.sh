#!/bin/bash

# verify docker is running
if ! command -v docker &> /dev/null
then
    echo "Are you working with docker ?"
    exit 1
fi

# explicitly declare an indexed array
declare -a docker_images
declare -a images_to_delete

# collect and store docker image in an array with specific format
mapfile -t docker_images < <(docker images --format '{{.Repository}}:{{.Tag}}')

for image in "${docker_images[@]}"; do
    # display a prompt message before accepting user input intead of using echo
    read -p "Do you want to delete '$image'? [y/N]: " confirm
    if [[ "$confirm" == "y" || "$confirm" == "Y" ]]; then
        images_to_delete+=("$image")
        echo "Marked '$image' for deletion."
    else
        echo "Skipping '$image'."
    fi
    echo -e "\n"
done

if [ ${#images_to_delete[@]} -gt 0 ]; then
    echo "Deleting images: ${images_to_delete[*]}"
    docker rmi "${images_to_delete[@]}"
else
    echo "No images were marked for deletion."
fi