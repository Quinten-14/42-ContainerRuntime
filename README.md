## 42-MinishellToolkit

**List of all the features I want to add**
| Feature | Status |
|--|--|
| Containers | :heavy_check_mark: |
| Tester| &#10007; |
| Docs| &#10007; |
| References| &#10007; |
| Utils| :heavy_check_mark: |

The project is still a work in progress and I will continue working on it. ETA: end of 2024

## Overview of the Features

**Containers**

In these containers you can run your minishell. They come packed together with valgrind, gdb and all the dependencies you need. These containers can be used with Docker and they have the purpose to limit the time spend on editing dependencies and code for different Operating Systems. This is a great tool for groups where both devs use different environments and OS.

[Link To Feature](https://github.com/Quinten-14/42-MinishellToolkit/tree/main/containers)

**Tester**

This will be a very extensive tester that compares your results with bash. This is the perfect starting point to start fixing your code when something is broken. The tester gives you back really extensive error messages and it is easy to navigate and see the test case. Fully written in GO but I well documented is the main focus behind this. Also an easy to use config file to add your own tests. The tester also tests for memory leaks and returns all the results in different log formats. These can then be used by the Comparer Util to see if you broke something accidently since your last test.

[Link To Feature](https://github.com/Quinten-14/42-MinishellToolkit/)

**Docs**

A pretty big registry of documentation about the different parts of Minishell and my thought proccess behind them. Not only does this include Minishell items but also some things about Data Structures and more. Nice for if you are a visual learner.

[Link To Feature](https://github.com/Quinten-14/42-MinishellToolkit/)

**References**

A list of all the references I used to complete Minishell and to find more info about how bash does things in the background. This also has some of my favorite Minishell's linked so you can check these out.

[Link To Feature](https://github.com/Quinten-14/42-MinishellToolkit/)

**Utils**

A bunch of small tools that can help you in Minishell. This can go from valgrindrc files to test tools. These are really made to make you find smaller issues or to make life a little bit less painfull when you mess up.

[Link To Feature](https://github.com/Quinten-14/42-MinishellToolkit/)
