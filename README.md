glowing-dangerzone
==================

Just my home made static blog generator

## Usage ##

### Create blog ###

To create a new blog :

    cd /path/to/my/blog/
    glowing-dangerzone new

This will generate two folders :

    layouts : place for HTML layout files
    posts : place for your posts
    assets : place for your css, js, the whole directory will be copied to the build directory

### Create post ###

To create a post :

    glowing-dangerzone post -title='This is a pen, that is a book'

It will generate a folder named `this-is-a-pen-that-is-a-book` under `posts`. Inside the folder, a `post.md` will also be generated with following content :

    ---
    date: 2015-01-12
    title: This is a pen, that is a book
    ---

Then you can start to write your post.

### Build blog ###

To generate static files for your blog :

    glowing-dangerzone build

It will create a folder `blog` and put all generated files under it.


### Serve your blog files ###

    glowing-dangerzone serve -port=80

Then it will run a file server on `blog` folder.

## Configuration ##

All configuration values should be stored as system environment variables, there are:

**GD_GITHUB_ACCESS_TOKEN** : Required. Access token of github, it will be used when convert markdown to HTML
