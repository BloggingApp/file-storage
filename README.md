## CDN

#### BloggingApp's own CDN for uploading avatars and post images

### API Docs

**POST** -> `/upload` - *upload image or video*
```
Request Body Type: multipart/form-data
Required fields:
    type: "IMAGE" || "VIDEO"
    file: File to upload
    path: Upload file path
```

**POST** -> `/move` - *move files to new path*
```
Request Body Type: json
Format:
oldPath: newPath
e.g.: public/post-images/temp/name.png: public/post-images/perm/name.png
```

**GET** -> `/public/**` - *get specified file by path*
