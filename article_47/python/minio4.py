import argparse

from minio import Minio
from minio.error import ResponseError


def list_buckets(minioClient):
    buckets = minioClient.list_buckets()
    for bucket in buckets:
        print(bucket.name, bucket.creation_date)


def list_objects(minioClient, bucket):
    print("List of objects for bucket:", bucket)
    objects = minioClient.list_objects(bucket, prefix="", recursive=False)
    for object in objects:
        print(object.bucket_name, object.last_modified, object.etag,
              object.size, object.object_name)


parser = argparse.ArgumentParser()
parser.add_argument("--endpoint", default="127.0.0.1:9000",
                    help="MinIO service endpoint")
parser.add_argument("--accessKeyID", default="",
                    help="Access key ID for MinIO")
parser.add_argument("--secretAccessKey", default="",
                    help="Secret access key for MinIO")
parser.add_argument("--enable-ssl", dest="useSSL", action="store_true",
                    help="Use SSL for communication with MinIO")
parser.add_argument("--disable-ssl", dest="useSSL", action="store_false",
                    help="Don't SSL for communication with MinIO")
args = parser.parse_args()

minioClient = Minio(args.endpoint,
                    access_key=args.accessKeyID,
                    secret_key=args.secretAccessKey,
                    secure=args.useSSL)
list_buckets(minioClient)
list_objects(minioClient, "foo")
