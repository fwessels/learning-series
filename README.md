## Learning Series

How to run minio for learning series (on local laptop):

```
CI=true MINIO_STORAGE_CLASS_STANDARD=EC:3 MINIO_STORAGE_CLASS_RRS=EC:2 ./minio server data{1...8}
```
