# ytMod
simple timer to interactive with via curl for tracking yt watch time day to day

## api end points are as follows:
### /api/timer
- /current for current watch time
- /remain for remaining watch time for the day
- /start to start a timer
- /end to end timer and add watch time to GetCurrentWatchTime
- /bonus to add time to daily allotment

this project serves as an internal tool primarily for my personal use to reduce/track yt usage
a simple 'go run main' cmd to start server and curl requests from there to start tracking
thanks
