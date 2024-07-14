# Initialize a git repository and push it to GitHub and Gitea
git init

# Add all files to the repository
git add .

# Status of the repository
git status

# Commit the files
git commit -m "Initial commit"

# Push the repository to GitHub
git remote add origin1 https://github.com/pro12x/social-network.git

# Push the repository to Gitea
git remote add origin2 https://learn.zone01dakar.sn/git/fmokomba/social-network.git

# Push the repository to both GitHub and Gitea
git push -u origin1 master
git push -u origin2 master

# End of script