# server/python/bgrm.py
import sys, io
from PIL import Image
from rembg import remove

data = sys.stdin.buffer.read()
if not data:
    sys.exit("no input")
img = Image.open(io.BytesIO(data)).convert("RGBA")
out = remove(img)
buf = io.BytesIO()
out.save(buf, format="PNG")
sys.stdout.buffer.write(buf.getvalue())
