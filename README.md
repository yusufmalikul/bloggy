# bloggy
A simple and stup*d static blog generator written in Go.

# Status
Currently, bloggy is in a very early stage of development. It is not ready for production use.

# Why
I wanted to create a simple blog generator that is easy to use. I prefer static sites because it's cheap to host and maintain. It also reduce security risks for your server.

# Usage
First, you need to prepare your blog posts content. Layouts directory must have `index.html` and `post.html` see ([examples](examples/layouts)). Put your markdown files e.g. in `posts` directory.
```
$ tree myblog
myblog
├── layouts
│   ├── index.html
│   └── post.html
└── posts
    ├── hello-world.md
    └── second-post.md
```

Run this command to generate markdown files to your designated output directory.

```
$ bloggy --content myblog/posts --layouts myblog/layouts --output myblog/html
Generating...
Done in 952.06µs
```

It will generate html files in `myblog/html` directory.
```
$ tree myblog
myblog
├── html
│   ├── hello-world
│   │   └── index.html
│   ├── index.html
│   └── second-post
│       └── index.html
├── layouts
│   ├── index.html
│   └── post.html
└── posts
    ├── hello-world.md
    └── second-post.md
```

After that, you can start bloggy as a server.
```
$ bloggy server --dir myblog/html
Listening...
Visit http://localhost:8080
```

Visit http://localhost:8080 to see your blog. You can also deploy the static files (`myblog/html`) to your server.

# Styling
You can style your blog by editing `index.html` and `post.html` in `layouts` directory. Bloggy uses [Go's template](https://golang.org/pkg/text/template/). The available variables are:

- `Title` - title of the post
- `Body` - body of the post
- `Path` - path of the post

# License
MIT