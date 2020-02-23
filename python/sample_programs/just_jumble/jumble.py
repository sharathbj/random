'''
hello world ,
	have you tried "Just Jumble" game in android ?
	its a very good app ,which gives a jumbled words
	and we need to find the current sequence of
	characters to find a meaningful word (they also
	provide the scenario).

	But sometimes the characters are so difficult
	that, we can find exact word ,after so many 
	tries ,and sad thing is u cant proceed further
	unless you solve and need to pay them, if 
	you want an answer :(
	
	So i have made a sample scipt ,that generates
	all posible words and call a rest api
	(from dictionary.com ) and checks if it a 
	proper word,after that you can choose the word,
	which suites as the best answer
'''

import sys
import json
import requests as r
word = sys.argv[1]
print('word:',word)
dicts = {}

def callApi(keyWord):
	query = 'https://api-portal.dictionary.com/dcom/pageData/'+keyWord
	resp = r.get(query).json()
	return resp

def swap(lists,i,j):
	temp = lists[i]
	lists[i] = lists[j]
	lists[j]= temp
	return ''.join(lists)

def permutation(word , si ,ei):
	index = -1
	while si <= ei:
		index+=1
		if index >= len(word):
			break
		word1 = swap(list(word),si,index)
		permutation(word1,si+1,ei)
	global dicts
	dicts[word]=1

permutation(word,0,len(word)-1)
count = 0
for val in dicts:
	count+=1
	try:
		respData = callApi(val)
	except Expection as e:
		print('exception',val,e)
		continue
	if respData['data'] is not None:
		print('try this word:',val)
	
