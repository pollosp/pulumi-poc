package main

import (
	"html/template"
	"mime"
	"os"
	"path"

	"github.com/pulumi/pulumi-gcp/sdk/go/gcp/storage"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

func generateFile() error {
	var tmpl = `<html>
<head>
    <title>Hello World!</title>
</head>
<body>
    {{ . }}
</body>
</html>
`
	var indexFile *os.File
	var err error
	t := template.New("main")
	t, _ = t.Parse(tmpl)
	indexFile, err = os.Create("/tmp/index.html")

	t.Execute(indexFile, os.Getenv("GREETING"))

	indexFile.Close()
	if err != nil {
		return err
	}
	return nil
}

func createBucket(ctx *pulumi.Context, StorageClass string) (*storage.Bucket, error) {
	bucket, err := storage.NewBucket(ctx, "pulumibucket", &storage.BucketArgs{
		ForceDestroy: true,
		StorageClass: StorageClass,
		Name:         "my-bucket-pulumi-poc-olopez",
	})
	if err != nil {
		return bucket, err
	}
	return bucket, nil
}

func main() {
	//Generate file from given template
	generateFile()

	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a GCP resource (Storage Bucket) example extractig to a func
		bucket, err := createBucket(ctx, "COLDLINE")

		// Create a GCS Object from the template
		gcsObject, err := storage.NewBucketObject(ctx, "index.html", &storage.BucketObjectArgs{
			Bucket:      bucket.ID(),
			Source:      "/tmp/index.html",
			Name:        "index.html",
			ContentType: mime.TypeByExtension(path.Ext("index.html")), // set the MIME type of the file
		})
		if err != nil {
			return err
		}

		var RolesList [1]string
		RolesList[0] = "READER:allUsers"

		//Set as public the uploaded file to GCS
		if _, err := storage.NewObjectACL(ctx, "indexAllUsers", &storage.ObjectACLArgs{
			Bucket:       bucket.ID(),
			Object:       gcsObject.Name(),
			RoleEntities: RolesList,
		}); err != nil {
			return err
		}

		// Export the DNS name of the bucket
		ctx.Export("bucketName", bucket.Url())
		return nil
	})
}
