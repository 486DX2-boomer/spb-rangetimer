# SPB Range Timer

An array of server-side stopwatches for shooting range use.

CTRL + Shift + H to hide the control buttons. This is intended to be used for the screens inside the range and provides a cleaner look to customers.

## To Do

- [ ] ~~Start~~, ~~Stop~~, ~~Clear Timer~~, Half Hour, Hour, Custom Time Interval (Hour, half hour, or arbitrary time), ~~Reserved for Member~~, ~~Out of Order~~
- [ ] Add expiration behavior to timers (timers should stay at 0:00 and clear manually instead of auto clearing)
- [ ] Having an array of 21 timers starting at 1 and ending at 20 is kind of iffy, don't know if I want to do 0-19 and then translate the value at the frontend. Either way is confusing
- [ ] Refactor this so I send one request and get the state of all 20 timers with one request instead of sending 20 requests every second (But I couldn't figure out how)
- [ ] Get timer index from GET request logic (`timerIndex := strings.Split(ti, "/start/")[1]`) rewritten into a generic
- [ ] Toggling timer flags (reserved, member, out of order) are currently all separate functions, rewrite into a generic
- [x] All of the front end work (has to look good because customers will see it)
- [x] Figure out how to call Refresh() without typing the URL in (lol)
- [x] Front end fetch the timers from the server and display them in DOM
  - [x] Do not know how to refresh this every second- a fetch request every second? Sounds expensive
- [x] Convert integer to time (3600 seconds == 1 hour on screen)
- [x] Hide Timer Buttons (For range use, so savvy customers don't reset their own timers)

## Code Snippets

*Emoji button setup*
```     
<button onclick="startTimer('1')">▶</button>
<button onclick="stopTimer('1')">||</button>
<button onclick="clearTimer('1')">⯀</button>
<!-- Out of Order -->
<button>❌</button>
<!-- Member -->
<button>🤴</button>
<!-- Reservation -->
<button>🛎</button>
```

*Format seconds as an integer value into a stopwatch value*

I like this answer. It can be even shorter: 
```
new Date(1000 * seconds).toISOString().substr(11, 8)
```

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

https://www.flaticon.com/

https://stackoverflow.com/questions/6312993/javascript-seconds-to-time-string-with-format-hhmmss

https://www.sitepoint.com/css-sizing-absolute-position/

https://css-tricks.com/content-jumping-avoid/

https://itnext.io/how-to-stop-content-jumping-when-images-load-7c915e47f576

https://frontendresource.com/css-cards/

https://code-boxx.com/shortcut-keys-javascript/

https://wangchujiang.com/hotkeys/