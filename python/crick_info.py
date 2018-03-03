import json
import os
import subprocess
import datetime
from pycricbuzz import Cricbuzz
c = Cricbuzz()

for m in c.matches():
	if (m['mchstate'] == 'inprogress' or m['mchstate'] == 'lunch') and (m['mchdesc'].startswith('AUS vs') or m['mchdesc'].endswith('vs AUS')):
	#if True:	
		liveScores = c.livescore(m['id'])
		status = liveScores['matchinfo']['status'] + ' (' + liveScores['matchinfo']['mchstate'] + ')'
		

		scoreCard = c.scorecard(m['id'])
		#score = scoreCard['scorecard'][0]['runs']+'/'+scoreCard['scorecard'][0]['wickets']+ ' ('+  scoreCard['scorecard'][0]['overs']+' Ovs)'
		score = scoreCard['scorecard'][0]['runs']+'/'+scoreCard['scorecard'][0]['wickets']+ '('+  scoreCard['scorecard'][0]['overs']+'Ovs)'
		
		bat_0 ,bat_1 = "" ,""
		batsman0 = liveScores['batting']['batsman'][0]
		bat_0 = batsman0['name']+'-'+batsman0['runs']+'-'+batsman0['balls']+'-'+batsman0['fours']+'-'+batsman0['six']
		if len(liveScores['batting']['batsman']) > 1:
			batsman1 = liveScores['batting']['batsman'][1]
			bat_1 = batsman1['name']+'-'+batsman1['runs']+'-'+batsman1['balls']+'-'+batsman1['fours']+'-'+batsman1['six']
		
		bowl_0 , bowl_1 = "",""
		bowler0 = liveScores['bowling']['bowler'][0]
		bowl_0 = bowler0['name']+'-'+bowler0['overs']+'-'+bowler0['maidens']+'-'+bowler0['runs']+'-'+bowler0['wickets']
		if len(liveScores['bowling']['bowler']) > 1:
			bowler1 = liveScores['bowling']['bowler'][1]
			bowl_1 = bowler1['name']+'-'+bowler1['overs']+'-'+bowler1['maidens']+'-'+bowler1['runs']+'-'+bowler1['wickets']
		
		
		commentary = c.commentary(m['id'])
		comment = commentary['commentary'][0]

		print datetime.datetime.now()
		print 'Status:',status
		
		print 'Score:',score

		print 'Batsmen:'
		print bat_0,' ',bat_1 

		print 'Bowlers:'
		print bowl_0,' ',bowl_1

		print 'Commentary:'
		print(comment)
	
		new_val = status.replace(' ','_')		
		subprocess.call(['./shell.sh',new_val,score])
		#break
		
