package main

import (
	"blog_project/blog/blogpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Blog Client")
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:5001", opts)
	if err != nil {
		log.Fatalf("Could not connect to the server: %v\n", err)
	}
	defer cc.Close()

	// Create Blog
	c := blogpb.NewBlogServiceClient(cc)
	blog := &blogpb.Blog{
		AuthorId: "Tunaberk",
		Title:    "My First Blog",
		Content:  "Web3 is very exciting",
	}
	fmt.Println("Creating a new blog")
	resBlog, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{
		Blog: blog,
	})
	if err != nil {
		log.Fatalf("Error while creating the blog: %v\n", err)
	}
	fmt.Printf("Blog created with id: %v\n", resBlog.GetBlog().GetId())
}
