# bloggy
A simple and stup*d static blog generator written in Go.

# Status
Currently, bloggy is in a very early stage of development. It is not ready for production use.

# Why
I wanted to create a simple blog generator that is easy to use. I prefer static sites because it's cheap to host and maintain. It also reduce security risks for your server.

# Usage
First, you need to prepare your blog posts content.
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

Run this command to generate markdown file to your designated output directory.

```
$ bloggy --content myblog/posts --layouts myblog/layouts --output myblog/html
Generating...
Done in 2.131284ms
```

It will generate html files in `myblog/html` directory.
```
$ tree myblog
myblog
├── html
│   ├── hello-world.html
│   ├── index.html
│   └── second-post.html
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

# License
MIT