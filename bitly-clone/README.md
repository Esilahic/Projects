# GOLY - bitly clone

### Description
I am following along with this [tutorial](https://www.youtube.com/watch?v=bTLQT7W12dQ)

This is a clone of bitly, a link shortening service. It is a simple web application that allows users to shorten long URLs much like TinyURL.com and bit.ly do.

### Docker
taken from tutorial above:
Using docker to run postgres with the command below:
```
$ docker run --name name-of-container -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=test -d postgres:14
```