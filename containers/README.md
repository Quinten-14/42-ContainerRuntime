# Containers for Minishell Development

These containers allow you to run your code and conduct tests easily. Both an Alpine and an Ubuntu Dockerfile are included, featuring minimal base images for quicker builds.

## Dependencies

- **Docker**: Ensure you have Docker installed on your machine.
- **Your Source Files**: Prepare your project structure as described below.

## What’s Included in the Containers?

- **Compiler**: GCC for compiling your code.
- **Dependencies**: Necessary libraries and tools.
- **Valgrind**: For memory leak detection and profiling.
- **GDB**: The GNU Debugger for debugging your applications.
- **Vim**: A text editor for in-container editing.

## Project Structure

Organize your project files as follows:

Put your src files in a src directory. Include files in an include directory and your libft files in a libft directory. The Makefile can be in the root of the project. So all these dirs and Makefile need to be in the same place where the ubuntuBuild and alpineBuild are located.

You can also always change the location in the build file to the specific location you want.

(A fancy tree is coming soon)

## Build

To build the Docker images, use the following commands:

### For Alpine

```bash
docker build . -t minishell -f alpineBuild
```

### For Ubuntu

```bash
docker build . -t minishell -f ubuntuBuild
```

## Run

To run the Docker images in a container, use this command:

```bash
docker run -it -v "$(pwd)/src:/app/src" -v "$(pwd)/libft:/app/libft" -v "$(pwd)/include:/app/include" minishell:latest
```


## Support

If you need any support or something is not working, contact me on the 42 Slack at the name qraymaek or open a ticket on GitHub.




## FAQ

#### Can I contribute to this project?

Yes, We would even love it if you contribute so we can all make minishell less stressfull for students after us. If you have any contributes feel free to open a Pull Request. I will review them within 24 hours and give you the credit you deserve.

#### Are these Containers usefull for evaluations?

Normally they should be fine because all the dependencies are the same as the ones installed on my campus. If the dependencies would be different or you need a different build let me know in an issue and I will try to make one for you within 48 hours.


#### Why should I use containers?

Containers can be a handy tool for when you work on a different OS compared to your groupmate. I found it very painfull working on MacOS while my friend was working on Linux and always needing to change some dependencies and small code pieces. Containers are also a lot more lightweight and portable than VM's. Also this Valgrind and GDB will work on unsupported platforms like the new Mac's with the ARM chips.
