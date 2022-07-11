"""
base64's `urlsafe_b64encode` uses '=' as padding.
These are not URL safe when used in URL paramaters.
Functions below work around this to strip/add back in padding.

See:
https://docs.python.org/2/library/base64.html
https://mail.python.org/pipermail/python-bugs-list/2007-February/037195.html

Referrence:
https://gist.github.com/cameronmaske/f520903ade824e4c30ab
"""

import base64


def base64_encode(string):
    """
    Removes any `=` used as padding from the encoded string.
    """
    encoded = base64.urlsafe_b64encode(string.encode())
    return encoded.rstrip(b"=").decode()


def base64_decode(string):
    """
    Adds back in the required padding before decoding.
    """
    padding = 4 - (len(string) % 4)
    string = string + ("=" * padding)
    print(string)
    print(type(string))
    return base64.urlsafe_b64decode(string.encode()).decode()
