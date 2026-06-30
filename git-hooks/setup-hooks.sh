rm -rf .git/hooks/*
cp -R ./scripts/hooks/* .git/hooks/
chmod +x .git/hooks/*
echo "Hooks are set up!"
