import base64
import urllib.parse

# any string or file
message = "Python is fun"
print(f"message: {message}")
message_bytes = message.encode('ascii')
# 1. encode string to base64
base64_bytes = base64.b64encode(message_bytes)
base64_message = base64_bytes.decode('ascii')
print(f"base64 encoded message: {base64_message}")
# 2. url encode the base64
encoded = urllib.parse.quote(base64_message)
print(f"URL encoded message: {encoded}")
# 3. travel the network
# 4. url decode to base64
decoded = urllib.parse.unquote(encoded)
print(f"URL decoded message: {decoded}")
# 5. decode base64 to string
base64_decade = base64.standard_b64decode(decoded).decode("utf-8")
print(f"message: {base64_decade}")
