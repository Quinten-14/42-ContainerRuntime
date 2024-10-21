To build for Alpine use

docker build . -t minishell -f alpineBuild


To build for Ubuntu use (In progress)

docker build . -t minishell -f ubuntuBuild



To run build

docker run -it -v "$(pwd)/src:/app/src" -v "$(pwd)/libft:/app/libft" -v "$(pwd)/include:/app/include" minishell:latest

