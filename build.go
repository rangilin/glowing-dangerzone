package main

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"sort"
)

// NewBlogBuilder create a BlogBuilder instance that will build blog from
// with files in specified directory and configuration
func NewBlogBuilder(parser PostParser, conf Configuration, dir string) BlogBuilder {
	return BlogBuilder{dir, parser, map[string]*template.Template{}, conf}
}

// A BlogBuilder that generate static files from posts/layouts
type BlogBuilder struct {
	// where posts, layouts directory exist
	blogDir string

	// Parser that parse Post from post files
	postParser PostParser

	// map contains compiled template for later use
	templates map[string]*template.Template

	// configuration
	conf Configuration
}

// Build the blog, it will parse and copy post files into specified output
// directory
func (b BlogBuilder) Build(output string) error {
	os.RemoveAll(output)
	os.Mkdir(output, os.ModePerm)

	if err := b.compileTemplates(); err != nil {
		return fmt.Errorf("Fail to compile templates due to %v", err)
	}

	paths, err := b.getPostPaths()
	if err != nil {
		return err
	}

	posts := []Post{}
	for _, path := range paths {
		post := b.postParser.Parse(path)
		posts = append(posts, post)

		postOutputDir := filepath.Join(output, Prettify(post.Title()))
		err := b.generatePost(post, postOutputDir)
		if err != nil {
			return fmt.Errorf("Fail to generate post due to %v", err)
		}
	}

	sort.Sort(PostsByDateDesc{posts})

	if err := b.generateBlogIndex(posts, output); err != nil {
		return fmt.Errorf("Fail to generate blog index due to %v", err)
	}

	if err := b.copyAssets(output); err != nil {
		return fmt.Errorf("Fail to copy asset files due to %v ", err)
	}
	return nil
}

// getPostPaths will return a list of post directories in the blog folder
func (b BlogBuilder) getPostPaths() ([]string, error) {
	paths, err := filepath.Glob(filepath.Join(b.blogDir, PostsDirName, "*"))
	if err != nil {
		return nil, err
	}

	dirs := []string{}
	for _, path := range paths {
		info, err := os.Stat(path)
		if err != nil {
			return nil, err
		}
		if info.IsDir() {
			dirs = append(dirs, path)
		}
	}
	return dirs, nil
}

// compileTemplates compile all templates for later use
func (b BlogBuilder) compileTemplates() error {
	templates := map[string]string{
		"index": IndexTemplateName,
		"post":  PostTemplateName,
	}

	baseTemplatePath := filepath.Join(b.blogDir, LayoutsDirName, BaseTemplateName)
	for name, filename := range templates {
		templatePath := filepath.Join(b.blogDir, LayoutsDirName, filename)
		template, err := template.ParseFiles(baseTemplatePath, templatePath)
		if err != nil {
			return fmt.Errorf("Fail to parse template %s due to %v", filename, err)
		}
		b.templates[name] = template
	}
	return nil
}

// generatePost will generate files for specified post and put under specified
// output directory
func (b BlogBuilder) generatePost(post Post, output string) error {
	if err := os.Mkdir(output, os.ModePerm); err != nil {
		return err
	}

	if err := b.copyPostFiles(post.Folder(), output); err != nil {
		return err
	}

	index := filepath.Join(output, "index.html")
	file, err := os.Create(index)
	if err != nil {
		return fmt.Errorf("Unable to create file %s", index)
	}

	data := map[string]interface{}{
		"Content": template.HTML(post.HtmlContent()),
		"Title":   post.Title(),
		"Date":    post.Date(),
	}

	if err := b.templates["post"].Execute(file, data); err != nil {
		return err
	}
	return nil
}

// generateBlogIndex will generate blog index file under specified output path
func (b BlogBuilder) generateBlogIndex(posts []Post, output string) error {
	index := filepath.Join(output, "index.html")
	file, err := os.Create(index)
	if err != nil {
		return fmt.Errorf("Unable to create file %s", index)
	}

	data := map[string]interface{}{
		"Posts": posts,
	}

	if err := b.templates["index"].Execute(file, data); err != nil {
		return err
	}
	return nil
}

// copyPostFiles will copy all files under post dir into output path, but
// exclude post.md
func (b BlogBuilder) copyPostFiles(postDir string, output string) error {

	postFilePath := filepath.Join(postDir, PostFileName)

	walkFn := func(path string, info os.FileInfo, err error) error {
		if path == postDir || path == postFilePath {
			return nil
		}

		rel, err := filepath.Rel(postDir, path)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return os.Mkdir(filepath.Join(output, rel), os.ModePerm)
		}
		return CopyFile(path, filepath.Join(output, rel))
	}

	return filepath.Walk(postDir, walkFn)
}

// copyAssets copy all files in assets to specified output directory
func (b BlogBuilder) copyAssets(output string) error {
	assetsDir := filepath.Join(b.blogDir, AssetsDirName)
	walkFn := func(path string, info os.FileInfo, err error) error {
		if path == assetsDir {
			return nil
		}

		rel, err := filepath.Rel(assetsDir, path)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return os.Mkdir(filepath.Join(output, rel), os.ModePerm)
		}
		return CopyFile(path, filepath.Join(output, rel))
	}

	return filepath.Walk(assetsDir, walkFn)
}
