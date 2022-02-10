function startTimer() {
    let promise = fetch("http://localhost:8090/start")
    console.log("Starting timer...")
    }

function stopTimer() {
    let promise = fetch("http://localhost:8090/stop")
    console.log("Stopping timer...")
    }