# SPB Range Timer

An array of server-side stopwatches for shooting range use.

##

To Do
- [x] Figure out how to call Refresh() without typing the URL in (lol)
- [ ] All of the front end work (has to look good because customers will see it)
- [x] Front end fetch the timers from the server and display them in DOM
  - [x] Do not know how to refresh this every second- a fetch request every second? Sounds expensive
  - [ ] I am just doing 20 fetch requests a second until I figure out how to fix this
- [ ] Convert integer to time (3600 seconds == 1 hour on screen)
- [ ] ~~Start~~, ~~Stop~~, Clear Timer, Half Hour, Hour, Custom Time Interval (Hour, half hour, or arbitrary time),  Reserved for Member, Out of Order
- [ ] Having an array of 21 timers starting at 1 and ending at 20 is kind of iffy, don't know if I want to do 0-19 and then translate the value at the frontend. Either way is confusing
- [ ] I want to refactor this so I send one request and get the state of all 20 timers with one request instead of sending 20 requests every second (But I couldn't figure out how)
- [ ] Hide Timer Buttons (For range use, so savvy customers don't reset their own timers)

## References

https://code-boxx.com/simple-javascript-stopwatch/

https://jsfiddle.net/dalinhuang/op8ae79j/#:~:text=multiple%20stop%20watch%20-%20JSFiddle%20-%20Code%20Playground,2%20var%20Stopwatch%20%3D%20function%28elem%2C%20options%29%20%7B%203

https://stackoverflow.com/questions/22796620/multiple-countdown-timers-on-one-page

https://javascript.info/fetch

https://javascript.info/promise-api

https://www.educba.com/javascript-button/ 

No, I don't know how to make Javascript buttons off the top of my head. Javascript is a bad and I avoid learning it at ALL COSTS.

https://tutorialedge.net/golang/creating-simple-web-server-with-golang/

https://stackoverflow.com/questions/64032097/can-i-get-a-fetch-function-to-repeat-every-few-seconds

https://stackoverflow.com/questions/34842526/update-console-without-flickering-c

https://cssbuttons.io

https://css.gg

https://code-boxx.com/simple-javascript-stopwatch/

https://golangbyexample.com/json-response-body-http-go/