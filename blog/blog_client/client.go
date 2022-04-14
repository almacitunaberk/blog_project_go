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

	c := blogpb.NewBlogServiceClient(cc)

	// Create Blog
	/*

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
	*/

	// Read Blog
	res, err := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{
		BlogId: "6257ab411d78625941baf004",
	})
	if err != nil {
		log.Fatalf("Error happened when requested a blog: %v\n", err)
	}
	fmt.Printf("Read the blog: %v\n", res.GetBlog().GetContent())

	// Update Blog

	newBlog := &blogpb.Blog{
		Id:       "6257ab411d78625941baf004",
		AuthorId: "Tunaberk Almaci",
		Title:    "Edited Blog",
		Content:  "Web3, Metaverse and Blockchain are awesome (2)",
	}
	updateRes, updateErr := c.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{
		Blog: newBlog,
	})
	if updateErr != nil {
		fmt.Printf("Error happened while returning updated blog: %v\n", err)
	}
	fmt.Printf("Read the updated blog: %v\n", updateRes)

	// Delete Blog
	deleteRes, deleteErr := c.DeleteBlog(context.Background(), &blogpb.DeleteBlogRequest{
		BlogId: "6257ab411d78625941baf004",
	})
	if deleteErr != nil {
		fmt.Printf("Error happened while deleting a blog: %v\n", err)
	}

	fmt.Printf("Deleted the blog with ID: %v\n", deleteRes)

}
