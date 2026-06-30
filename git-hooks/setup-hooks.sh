rm -rf .git/hooks/*
cp -R ./git-hooks/* .git/hooks/
chmod +x .git/hooks/*
echo "Hooks are set up!"
