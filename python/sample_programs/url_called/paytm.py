import re
from urllib import request as r
req =  r.Request('https://www.paytm.com')
resp = r.urlopen(req)
lists = re.findall(r'href="http[s]://(.*?)"',str(resp.read()))
print('len',len(lists))
map={}
for l in lists:
	map[l]=1
for m in map:
	print(m)

