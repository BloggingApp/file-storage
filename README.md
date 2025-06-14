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

**POST** -> `/delete` - *delete files*
```
Request Body Type: json
Format:
array of strings with path
e.g.: ["some/path", "some/other/path"]
```

**GET** -> `/public/**` - *get specified file by path*
