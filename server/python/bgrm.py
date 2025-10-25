# /Users/bimap/Documents/Coding/Project/AngelaClothes/server/python/bgrm.py
import sys
try:
    import pillow_heif; pillow_heif.register_heif_opener()
except Exception:
    pass
from rembg import remove

data = sys.stdin.buffer.read()
sys.stdout.buffer.write(remove(data))
