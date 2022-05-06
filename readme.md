# SPB Range Timer

An array of server-side stopwatches for shooting range use.

CTRL + Shift + H to hide the control buttons. This is intended to be used for the screens inside the range and provides a cleaner look to customers.

## To Do

- [ ] ~~Start~~, ~~Stop~~, ~~Clear Timer~~, ~~Half Hour~~, ~~Hour~~, Custom Time Interval (Hour, half hour, or arbitrary time), ~~Reserved for Member~~, ~~Out of Order~~
- [x] Add expiration behavior to timers (timers should stay at 0:00 and clear manually instead of auto clearing)
- [x] Refactor this so I send one request and get the state of all 20 timers with one request instead of sending 20 requests every second (But I couldn't figure out how)
  - [x] OK so it turns out json.Marshal() will already jsonize an array without having to write a special method for it. (Not just single structs.) so I can just json.Marshal(t) and send that to the frontend
  - [x] Done. It was way easier than expected
- [ ] Get timer index from GET request logic (`timerIndex := strings.Split(ti, "/start/")[1]`) rewritten into a generic
- [ ] Toggling timer flags (reserved, member, out of order) are currently all separate functions, rewrite into a generic
- [x] All of the front end work (has to look good because customers will see it)
- [x] Figure out how to call Refresh() without typing the URL in (lol)
- [x] Front end fetch the timers from the server and display them in DOM
- [x] Convert integer to time (3600 seconds == 1 hour on screen)
- [x] Hide Timer Buttons (For range use, so savvy customers don't reset their own timers)
- [ ] 2 minute grace period for shooters coming onto the range (3720 seconds instead of 3600)

## Code Snippets

*Format seconds as an integer value into a stopwatch value*
```
new Date(1000 * seconds).toISOString().substr(11, 8)
```

## Issues

- UI: Laggy feel. This is because the backend ticks only once per second in addition to the 
browser fetching once per second. Because of that, activating a button can take up to one second which makes the UI feel slow. There's no way to force an "instant" update of the UI as written.
- Backend is written in Go instead of a more common language.
- The timers are not written as a reusable component.
- No user authentication.
- Not really set up for multiple deployments (server endpoint is stored in a constant instead of an environment variable or config file.)
- Some functions could be rewritten into generics.
