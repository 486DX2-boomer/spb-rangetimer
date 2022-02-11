function startTimer(buttonNumber) {
    let promise = fetch("http://localhost:8090/start/" + buttonNumber)
    console.log("Starting timer...")
    }

function stopTimer(buttonNumber) {
    let promise = fetch("http://localhost:8090/stop/" + buttonNumber)
    console.log("Stopping timer...")
    }