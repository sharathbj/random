My Doctor has adviced me to close my Eyes frequency ,since my eye surface is very dry ,but i used to forgot this,as usual.
So i just created a desktop notifier which will remind me to blink my eyes!!

Steps :
1) install any desktop notifiers ,i have used "terminal-notifier" (i tried with notify,notify2 and many other package,
nothing worked i just wasted time ( in MacOS)
2) follow the code in blink_eye.sh
3) create a cron job ,which schedules the shell script run(in my case it runs every 10 min).
    i,e 
        */10 * * * * ~/eye_blink.sh
