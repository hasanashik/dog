# dog

echo "# dog" >> README.md
git init
git add README.md
git commit -m "first commit"
git branch -M main
git remote add origin https://github.com/hasanashik/dog.git
git push -u origin main

## Create go module named as git repo at the root of my dog package

go mod init github.com/hasanashik/dog
