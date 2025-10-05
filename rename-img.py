import os
import sys
import random
import hashlib

# spent too mcuh time getting side tracked making this peice of shit

# renames any file in arg path to a base62*5 char name, uses hash for no collisions
# 50/50 chance it breaks and renames everything hehehe

BASE62 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
imgs = ['jpg', 'png', 'jpeg', 'webp']

arg_path = sys.argv[1]
files = os.listdir(arg_path)

def file_hash(path):
    h = hashlib.sha256()
    with open(path, "rb") as f:
        while chunk := f.read(8192):
            h.update(chunk)
    return h.digest()

for file in files:
    path = os.path.join(arg_path, file)
    
    if not os.path.isfile(path) or file == ".DS_Store":
        continue
    if file.split(".")[-1] not in imgs:
        continue
    
    h_bytes = file_hash(path)
    rng = random.Random(int.from_bytes(h_bytes, "big"))
    
    name = ''.join(rng.choices(BASE62, k=5)) + "." + file.split(".")[-1]
    name = os.path.join(arg_path, name)
    
    os.rename(path, name)
    print(f'{file:>12} -> {name}')