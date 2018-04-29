import re
from urllib import request as r
req =  r.Request('http://www.goibibo.com/')
resp = r.urlopen(req)
lists = re.findall(r'href="http[s]://(.*?)"',str(resp.read()))
print('len',len(lists))
for l in lists:
	print(l)

