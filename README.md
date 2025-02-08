## CDN

#### BloggingApp's own CDN for uploading avatars and post images

### API Docs

**POST** -> `/upload` - *upload image or video*
```
Request Type: multipart/form-data
Required fields:
    type: "IMAGE" || "VIDEO"
    file: File to upload
    path: Upload file path
```

**GET** -> `/public/**` - *get specified file by path*
