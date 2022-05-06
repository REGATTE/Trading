import logging
from kiteconnect import KiteConnect as kc

logging.basicConfig(level=logging.DEBUG)

kite_API_KEY = kc(api_key = "")

"""
Redirect user to the login URL obtained from kite.login_url(), 
and receive the request_topken from the registered redirect url after the login flow.
Once you have the request_token, obtain the access token as follows.
"""

data = kite_API_KEY.generate_session("request_token_here", secret="your_secret")
kite_API_KEY.set_access_token(data["access_token"])